--
--
-- Tencent is pleased to support the open source community by making tRPC available.
--
-- Copyright (C) 2024 THL A29 Limited, a Tencent company.
-- All rights reserved.
--
-- If you have downloaded a copy of the tRPC source code from Tencent,
-- please note that tRPC source code is licensed under the  Apache 2.0 License,
-- A copy of the Apache 2.0 License is included in this file.
--
--
-- tRPC is licensed under the Apache 2.0 License, and includes source codes from
-- the following components:
-- 1. incubator-brpc
-- Copyright (C) 2019 The Apache Software Foundation
-- incubator-brpc is licensed under the Apache 2.0 License.
--
--

local protocol_name    = "trpc" 
local trpc_proto       = Proto(protocol_name, "tRPC Protocol Dissector")

local field_magic      = ProtoField.uint16(protocol_name .. ".magic",       "Magic",       base.HEX)
local field_type       = ProtoField.uint8(protocol_name .. ".type",         "Packet Type", base.DEC)
local field_stream     = ProtoField.uint8(protocol_name .. ".stream",       "Stream Type", base.DEC)
local field_total_size = ProtoField.uint32(protocol_name .. ".total_size",  "Total Size",  base.DEC)
local field_header_size= ProtoField.uint16(protocol_name .. ".header_size", "Header Size", base.DEC)
local field_unique_id  = ProtoField.uint32(protocol_name .. ".unique_id",   "Unique ID",   base.DEC)
local field_version    = ProtoField.uint8(protocol_name .. ".version",      "Version",     base.DEC)
local field_reserved   = ProtoField.uint8(protocol_name .. ".reserved",     "Reserved",    base.DEC)
trpc_proto.fields      = {field_magic, field_type, field_stream, field_total_size, field_header_size, field_unique_id, field_version, field_reserved}

local MAGIC_CODE_TRPC = "0930"
local PROTO_HEADER_LENGTH = 16

local tcp_src_port     = Field.new("tcp.srcport")
local tcp_dst_port     = Field.new("tcp.dstport")
local tcp_stream       = Field.new("tcp.stream")

local proto_f_protobuf_field_name  = Field.new("protobuf.field.name")
local proto_f_protobuf_field_value = Field.new("protobuf.field.value")

local data_dissector     = Dissector.get("data")
local protobuf_dissector = Dissector.get("protobuf")

----------------------------------------
-- declare functions
local check_length  = function() end
local dissect_proto = function() end

----------------------------------------
-- main dissector
function trpc_proto.dissector(tvbuf, pktinfo, root)
    local pktlen = tvbuf:len()

    local bytes_consumed = 0

    while bytes_consumed < pktlen do
        local result = dissect_proto(tvbuf, pktinfo, root, bytes_consumed)

        if result > 0 then
            bytes_consumed = bytes_consumed + result
        elseif result == 0 then
            -- hit an error
            return 0
        else
            pktinfo.desegment_offset = bytes_consumed
            -- require more bytes
            pktinfo.desegment_len = -result

            return pktlen
        end
    end

    return bytes_consumed
end

--------------------------------------------------------------------------------
-- heuristic
-- tcp_stream_id <-> {client_port, server_port, {request_id<->method_name}}
local stream_map = {}
local function heur_dissect_proto(tvbuf, pktinfo, root)
    -- dynmaic decide client or server data
    -- by first tcp syn frame
    local f_src_port = tcp_src_port()()
    local f_dst_port = tcp_dst_port()()
    local stream_n = tcp_stream().value
    if stream_map[stream_n] == nil then
        stream_map[stream_n] = {f_src_port, f_dst_port, {}}
    end

    if (tvbuf:len() < PROTO_HEADER_LENGTH) then
        return false
    end

    local magic = tvbuf:range(0, 2):bytes():tohex()
    -- for range dissectors
    if magic ~= MAGIC_CODE_TRPC then
        return false
    end

    trpc_proto.dissector(tvbuf, pktinfo, root)

    pktinfo.conversation = trpc_proto

    return true
end

trpc_proto:register_heuristic("tcp", heur_dissect_proto)

--------------------------------------------------------------------------------

-- check packet length, return length of packet if valid
check_length = function(tvbuf, offset)
    local msglen = tvbuf:len() - offset

    if msglen ~= tvbuf:reported_length_remaining(offset) then
        -- captured packets are being sliced/cut-off, so don't try to desegment/reassemble
        LM_WARN("Captured packet was shorter than original, can't reassemble")
        return 0
    end

    if msglen < PROTO_HEADER_LENGTH then
        -- we need more bytes, so tell the main dissector function that we
        -- didn't dissect anything, and we need an unknown number of more
        -- bytes (which is what "DESEGMENT_ONE_MORE_SEGMENT" is used for)
        return -DESEGMENT_ONE_MORE_SEGMENT
    end

    -- if we got here, then we know we have enough bytes in the Tvb buffer
    -- to at least figure out whether this is valid trpc packet

    local magic = tvbuf:range(offset, 2):bytes():tohex()
    if magic ~= MAGIC_CODE_TRPC then
        return 0
    end

    local packet_size = tvbuf:range(offset+4, 4):uint()
    if msglen < packet_size then
        -- Need more bytes to desegment full trpc packet
        return -(packet_size - msglen)
    end

    return packet_size
end

--------------------------------------------------------------------------------

dissect_proto = function(tvbuf, pktinfo, root, offset)
    local len = check_length(tvbuf, offset)
    if len <= 0 then
        return len
    end

    -- update 'Protocol' field
    if offset == 0 then
        pktinfo.cols.protocol:set("tRPC")
    end

    local f_src_port = tcp_src_port()()
    local f_dst_port = tcp_dst_port()()

    local direction
    local stream_n = tcp_stream().value
    if f_src_port == stream_map[stream_n][1] then
        pktinfo.private["pb_msg_type"] = "message,trpc.RequestProtocol"
        direction = "request"
    end
    if f_src_port == stream_map[stream_n][2] then
        pktinfo.private["pb_msg_type"] = "message,trpc.ResponseProtocol"
        direction = "response"
    end

    -- check packet length, 
    local magic_value       = tvbuf(offset, 2)
    local type_value        = tvbuf(offset+2, 1)
    local stream_value      = tvbuf(offset+3, 1)
    local total_size_value  = tvbuf(offset+4, 4)
    local header_size_value = tvbuf(offset+8, 2)
    local unique_id_value   = tvbuf(offset+10, 4)
    local version_value     = tvbuf(offset+14, 1)
    local reserved_value    = tvbuf(offset+15, 1)

    local header_length   = header_size_value:uint()
    local total_length    = total_size_value:uint()
    local tree            = root:add(trpc_proto, tvbuf:range(offset, len), "tRPC Protocol Data")

    data_dissector:call(tvbuf, pktinfo, tree)

    local t = tree:add(trpc_proto, tvbuf)
    t:add(field_magic, magic_value)
    t:add(field_type, type_value)
    t:add(field_stream, stream_value)
    t:add(field_total_size, total_size_value)
    t:add(field_header_size, header_size_value)
    t:add(field_unique_id, unique_id_value)
    t:add(field_version, version_value)
    t:add(field_reserved, reserved_value)

    -- solve the problem of parsing errors when multiple RPCs are included in a packet
    local protobuf_field_names = { proto_f_protobuf_field_name() }
    local pre_field_nums = #protobuf_field_names
    pcall(Dissector.call, protobuf_dissector, tvbuf(offset+16, header_length):tvb(), pktinfo, tree)

    -- Add bussiness rpc pb
    -- Get invoke rpc method name from trpc.RequestProtocol
    protobuf_field_names = { proto_f_protobuf_field_name() }
    local cur_field_nums = #protobuf_field_names
    local protobuf_field_values = { proto_f_protobuf_field_value() }
    local method
    -- default request id
    local request_id = 0
    for k = pre_field_nums + 1, cur_field_nums do
        local v = protobuf_field_names[k]
        if v.value == "func" then
            method = protobuf_field_values[k].range:string(ENC_UTF8)
        elseif v.value == "request_id" then
            request_id = protobuf_field_values[k].range:uint()
        end
    end

    local tvb_body = tvbuf:range(offset + 16 + header_length, total_length - header_length - 16):tvb()
    if method ~= nil then
        -- only req contains method, correlate it with request id so that response protocol can use.
        stream_map[stream_n][3][request_id] = method
        pktinfo.private["pb_msg_type"] = "application/trpc," .. method .. "," .. direction
    else
        -- get method for the same request id
        method = stream_map[stream_n][3][request_id]
        pktinfo.private["pb_msg_type"] = "application/trpc," .. method .. "," .. direction
    end
    pcall(Dissector.call, protobuf_dissector, tvb_body, pktinfo, tree)

    return total_length
end