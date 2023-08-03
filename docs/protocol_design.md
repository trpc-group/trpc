[TOC]

# Capabilities and Demands

|  Capabilities |Demands   |
| ------------ | ------------ |
|  RPC	| remote interface call like local interface call |
|  Stream | the protocol should be able to solve the problem of transmitting large data packets and meet the scenario of streaming data transmission.  |
|  High Performance	 | the protocol design and implementation need to consider performance issues|
|  Compatibility & Scalability	| the protocol design and implementation should have forward compatibility and scalability |
|  Support Multiple Serialization | the protocol design should support multiple serialization methods, such as protobuf, JSON, etc. |
|  Support multiple Decompression |  the protocol design should support multiple decompression methods. | 
|  Request Timeout Control	  | The protocol should provide the ability to control request timeout  |
|  Passing Tracing Information |  The protocol should provide the ability to pass tracing information |
|  Passing Auth Information |  The protocol should provide the ability to pass auth information |
|  Passing Dyeing Information |  The protocol should provide the ability to pass dyeing information |
|  Passing Custom Business Information	 |  The protocol should provide the ability to pass custom business information |

# Design And Implementation

## Protocol Design

![ 'image.png'](/docs/images/trpc-protocol.png)

for unary RPC(one-request-one-response), the entire data packet = Fix Header（16 bytes）+ Unary Header + Unary Body.

for stream RPC, the entire data packet = Fix Header（16 bytes）+ Stream Frame(Init/Data/Feeback/CLose) 

the complete protocol definition, please see [trpc.proto](/trpc/trpc.proto)

## Protocol Details
### Fixed Header
- The first 2 bytes:
  The first two bytes of the protocol are magic numbers, indicating the start of the trpc frame, which is 0x930.

- The 3rd byte: the data frame type of the packet

```protobuf
// Two types are currently supported:
// 1. The data frame type for unary(one-response-one-response)
// 2. The data frame type for stream
enum TrpcDataFrameType {
TRPC_UNARY_FRAME = 0x00;

TRPC_STREAM_FRAME = 0x01;
}
```

- The 4th byte: streaming frame type, only valid when the data frame type is `TRPC_STREAM_FRAME`

```protobuf
// Four types are currently supported:
// `INIT` frame: FIXHEADER + TrpcStreamInitMeta
// `DATA` frame: FIXHEADER + body (business serialized data)
// `FEEDBACK` frame: FIXHEADER + TrpcStreamFeedBackMeta (triggered strategy: high/low water level and timer)
// `CLOSE` frame: FIXHEADER + TrpcStreamCloseMeta
enum TrpcStreamFrameType {
TRPC_UNARY = 0x00;

TRPC_STREAM_FRAME_INIT = 0x01;

TRPC_STREAM_FRAME_DATA = 0x02;

TRPC_STREAM_FRAME_FEEDBACK = 0x03;

TRPC_STREAM_FRAME_CLOSE = 0x04;
}
```

- The 5-8 bytes, the total size of the packet

For unary(one response one response), total data size = fixed header size + header size + body size

For stream, total data size = fixed header size + stream frame size

- The 9-10 bytes

For unary(one response one response), header size

For stream, set 0

- The 11-14 bytes

For unary(one response one response), unique-id

For stream, stream-id

stream-id needs to be unique in a single connection

- The 15 byte, the version of trpc protocol

- The 16 byte, reserved field

### The Header of Unary

After 16 bytes, for unary(one-response-one-response), it corresponds to the header of the protocol, which is divided into request header and response header.

### Request Header

```protobuf
// The request header for unary
message RequestProtocol {
  // The version of protocol
  // The specific value corresponds to `TrpcProtoVersion`
  uint32    version                     = 1;

  // Call type
  // eg: unary, one-way
  // The specific value corresponds to `TrpcCallType`
  uint32    call_type                   = 2;

  // The unique id of the request(on the conneciton)
  uint32    request_id                  = 3;

  // The timeout of the request(ms)
  uint32    timeout                     = 4;

  // Caller name
  // The specification format: trpc.application_name.server_name.proto_service_name, 4 segments
  bytes     caller                      = 5;

  // Callee name
  // The specification format: trpc.application_name.server_name.proto_service_name[.interface_name]
  bytes     callee                      = 6;

  // Interface name of callee
  // The specification format: /package.service_name/interface_name
  bytes     func                        = 7;

  // The message type of the transparent transmission information
  // such as tracing, dyeing key, gray, authentication, multi-environment, set name, etc.
  // The specific value corresponds to `TrpcMessageType`
  uint32    message_type                = 8;

  // The information key-value pair transparently transmitted by the framework
  // Currently divided into two parts:
  // 1 part is the information to be transparently transmitted by the framework layer,
  // and the name of the key must be started with `trpc-``
  // 2 part is the information to be transparently transmitted by the business layer,
  // and the business can set it by itself, it is recommended to start with `app-``, not `trpc-`
  // Note: The key-value pair in trans_info will be transparently transmitted through the whole link, please use it carefully for business.
  map<string, bytes> trans_info         = 9;

  // The serialization type of the request data
  // eg: proto/json/.., default proto
  // The specific value corresponds to `TrpcContentEncodeType`
  uint32    content_type                = 10;

  // The compression type of the requested data
  // eg: gzip/snappy/..., not used by default
  // The specific value corresponds to `TrpcCompressType`
  uint32    content_encoding            = 11;

  // The size of attachment data
  uint32    attachment_size             = 12;
}
```

### Response Header

```protobuf
// The response header for unary
message ResponseProtocol {
  // The version of protocol
  // The specific value corresponds to `TrpcProtoVersion`
  uint32    version                     = 1;

  // Call type
  // eg: unary, one-way
  // The specific value corresponds to `TrpcCallType`
  uint32    call_type                   = 2;

  // The unique id of the request(on the conneciton)
  uint32    request_id                  = 3;

  // Error code
  // The specific value corresponds to `TrpcRetCode`
  int32     ret                         = 4;

  // The error code of the interface
  // 0 means success, other means failure
  int32     func_ret                    = 5;

  // The result information when the call fails
  bytes    error_msg                    = 6;

  // The message type of the transparent transmission information
  // such as tracing, dyeing key, gray, authentication, multi-environment, set name, etc.
  // The specific value corresponds to `TrpcMessageType`
  uint32    message_type                = 7;

  // The information key-value pair transparently transmitted by the framework
  // Currently divided into two parts:
  // 1 part is the information to be transparently transmitted by the framework layer,
  // and the name of the key must be started with `trpc-``
  // 2 part is the information to be transparently transmitted by the business layer,
  // and the business can set it by itself, it is recommended to start with `app-``, not `trpc-`
  map<string, bytes> trans_info         = 8;

  // The serialization type of the request data
  // eg: proto/json/.., default proto
  // The specific value corresponds to `TrpcContentEncodeType`
  uint32    content_type                = 9;

  // The compression type of the requested data
  // eg: gzip/snappy/..., not used by default
  // The specific value corresponds to `TrpcCompressType`
  uint32    content_encoding            = 10;

  // The size of attachment data
  uint32    attachment_size             = 12;
}
```

### Stream Frame

After 16 bytes, for stream, it is followed by the stream frame

Divided into init, data, feedback, close frame.

- Init Frame is used for stream initialization

- Data Frame is used to transmit streaming data

- Feedback Frame is used to transmit flow control messages

- Close Frame is used to close the stream

#### Init Frame

```protobuf
// The message definition of stream `INIT` frame
message TrpcStreamInitMeta {
  // request meta information
  TrpcStreamInitRequestMeta request_meta = 1;

  // response meta information
  TrpcStreamInitResponseMeta response_meta = 2;

  // The window size is notified by the receiver to the sender
  uint32 init_window_size = 3;

  // The serialization type of the request data
  // eg: proto/json/.., default proto
  // The specific value corresponds to `TrpcContentEncodeType`
  uint32 content_type = 4;

  // The compression type of the requested data
  // eg: gzip/snappy/..., not used by default
  // The specific value corresponds to `TrpcCompressType`
  uint32 content_encoding = 5;
}
```

```protobuf
// The request meta information definition of stream `INIT` frame
message TrpcStreamInitRequestMeta {
  // Caller name
  // The specification format: trpc.application_name.server_name.proto_service_name, 4 segments
  bytes caller = 1;

  // Callee name
  // The specification format: trpc.application_name.server_name.proto_service_name[.interface_name]
  bytes callee = 2;

  // Interface name of callee
  // The specification format: /package.service_name/interface_name
  bytes func = 3;

  // The message type of the transparent transmission information
  // such as tracing, dyeing key, gray, authentication, multi-environment, set name, etc.
  // The specific value corresponds to `TrpcMessageType`
  uint32 message_type = 4;

  // The information key-value pair transparently transmitted by the framework
  // Currently divided into two parts:
  // 1 part is the information to be transparently transmitted by the framework layer,
  // and the name of the key must be started with `trpc-``
  // 2 part is the information to be transparently transmitted by the business layer,
  // and the business can set it by itself, it is recommended to start with `app-``, not `trpc-`
  // Note: The key-value pair in trans_info will be transparently transmitted through the whole link, please use it carefully for business.
  map<string, bytes> trans_info = 5;
};
```

```protobuf
// The response meta information definition of stream `INIT` frame
message TrpcStreamInitResponseMeta {
  // Error code
  // The specific value corresponds to `TrpcRetCode`
  int32 ret = 1;

  // The result information when the call fails
  bytes error_msg = 2;
};
```

#### Data Frame

The data frame corresponds to the message body of the actual request


#### Feedback Frame

```protobuf
// The meta information definition of stream `FEEDBACK` frame
message TrpcStreamFeedBackMeta {
  // increased window size
  uint32 window_size_increment = 1;
}
```

#### Close Frame

Close frame is divided into normal Close and abnormal Close (Reset)

```protobuf
// The closed type of trpc stream protocol
enum TrpcStreamCloseType {
  // normal closes unidirectional flow
  TRPC_STREAM_CLOSE = 0;

  // Exception closes bidirectional stream
  TRPC_STREAM_RESET = 1;
}
```

```protobuf
// The meta information definition of trpc stream protocol for closing stream
message TrpcStreamCloseMeta {
  // The type of stream closure, close one end, or close all
  int32 close_type = 1;

  // Error code
  // The specific value corresponds to `TrpcRetCode`
  int32 ret = 2;

  // The result information when the call fails
  bytes msg = 3;

  // The message type of the transparent transmission information
  // such as tracing, dyeing key, gray, authentication, multi-environment, set name, etc.
  // The specific value corresponds to `TrpcMessageType`
  uint32    message_type  = 4;

  // The information key-value pair transparently transmitted by the framework
  // Currently divided into two parts:
  // 1 part is the information to be transparently transmitted by the framework layer,
  // and the name of the key must be started with `trpc-``
  // 2 part is the information to be transparently transmitted by the business layer,
  // and the business can set it by itself, it is recommended to start with `app-``, not `trpc-`
  // Note: The key-value pair in trans_info will be transparently transmitted through the whole link, please use it carefully for business.
  map<string, bytes> trans_info = 5;

  // The error code of the interface
  // 0 means success, other means failure
  int32 func_ret = 6;
}
```


# 3 IDL

tRPC's IDL is based on Protobuf

Unary IDL

```protobuf
syntax = "proto3";

package proto;

service StreamService {
    rpc List(Request) returns ( Response) {};
    rpc Record( Request) returns (Response) {};
    rpc Route( StreamRequest) returns ( Response) {};
}


message Point {
  string name = 1;
  int32 value = 2;
}

message Request {
  Point pt = 1;
}

message Response {
  Point pt = 1;
}
```


Stream IDL

Support:

- Server-side streaming RPC
- Client Streaming RPC
- Bidirectional streaming RPC

```protobuf
syntax = "proto3";

package proto;

service StreamService {
    rpc List(StreamRequest) returns (stream StreamResponse) {};
    rpc Record(stream StreamRequest) returns (StreamResponse) {};
    rpc Route(stream StreamRequest) returns (stream StreamResponse) {};
}


message StreamPoint {
  string name = 1;
  int32 value = 2;
}

message StreamRequest {
  StreamPoint pt = 1;
}

message StreamResponse {
  StreamPoint pt = 1;
}
```
