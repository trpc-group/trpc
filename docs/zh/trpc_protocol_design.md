# 前言

本文主要介绍一下为什么要设计tRPC协议，以及tRPC协议具体是怎样定义。

# 为什么要设计tRPC协议

首先，tRPC是一个多语言的RPC框架，需要有统一的传输协议进行通信，以及通信时在一些功能上进行对齐。

其次，为什么tRPC不选择http/http2协议，这里主要的原因是性能问题。

最后，除了向前兼容外，tRPC在通信协议上还有强烈扩展性的诉求，能扩展支持各种业务需求的能力。

下面的表格是tRPC对通信协议的能力和诉求。

| 能力 | 诉求 |
| ------------ | ------------ |
|  RPC	|  远程调用像调本地接口一样  |
|  兼容性  |  协议设计和实现上应具有向前兼容  |
|  高性能  |  协议设计上需要考虑高性能问题  |
|  流式传输  |  协议设计上应该能支持大数据包传输和边请求边应答的流式场景  |
|  支持请求超时控制  |  协议设计和实现上支持设置请求超时时间  |
|  支持多种序列化方式  |  协议设计和实现上可扩展支持多种序列化方式，比如：protobuf、json、text等  |
|  支持多种解压缩方式  |  协议设计和实现上可扩展支持多种解压缩方式，比如：gzip、snappy等  | 
|  支持透传调用链信息  |  协议设计和实现上支持调用链信息在服务间透传 |
|  支持消息染色  |  协议设计和实现上支持请求消息的染色key在服务间透传  |
|  支持透传用户自定义元数据	 |  协议设计和实现上支持用户设置自定义元数据在服务间透传  |

为了支持这些能力和需求，tRPC协议整体设计设计如下：

传输协议（协议头，自定义设计）+ 编码协议（协议体，默认使用protobuf，可扩展）

# tRPC协议具体设计与实现

## 协议设计

下面的图是tRPC协议的具体设计。

![ 'image.png'](/docs/images/trpc-protocol.png)

tRPC协议支持一应一答和流式两种传输方式。

对于一应一答（unary）的RPC，整个完整协议帧 = 固定帧头（16字节）+ Unary包头（变长）+ Unary包体（变长）

对于流式（stream）RPC，整个完整协议帧 = 固定帧头（16字节）+ 流帧（Init帧、Data帧、Feedback帧、Close帧）

完整的协议定义，可以查看：[trpc.proto](https://github.com/trpc-group/trpc/blob/main/trpc/trpc.proto)

## 协议具体定义

### 固定帧头
- 第1-2个字节：
协议前面两个字节为魔数字，标识trpc帧的开始，为0x930

- 第3个字节：数据类型

```protobuf
// Two types are currently supported:
// 1. The data frame type for unary(one-response-one-response)
// 2. The data frame type for stream
enum TrpcDataFrameType {
TRPC_UNARY_FRAME = 0x00;

TRPC_STREAM_FRAME = 0x01;
}
```

- 第4个字节：
流式帧类型，只有帧为流式的帧的时候生效

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

- 第5-8个字节
数据总大小
对于一应一答，总数据大小 = 帧头大小 + 包头大小 + 包体大小
对于流式，总数据大小 = 帧头大小 + 流帧大小

- 第9-10个字节：
对于一应一答，包头大小
对于流式，始终为0

- 第11-14个字节：
对于一应一答，请求id
对流式，流式id

- 第15-16个字节
保留字段，保留给以后扩容协议使用

### 一应一答包头

16个字节之后，对于一应一答，对应着协议的包头，分为请求包头和响应包头

#### 请求包头

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

#### 响应包头

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

### 流式流帧

16个字节完整帧头后面紧接着流帧

分为 Init、Data、Feedback、Close 帧
- Init帧用来做流初始化

- Data帧用来传输流式数据

- Feedback帧用来传输流控消息

- Close帧用来关闭流

#### Init帧

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

#### Data帧

DATA帧对应的是实际请求的消息体

#### Feedback帧

```protobuf
// The meta information definition of stream `FEEDBACK` frame
message TrpcStreamFeedBackMeta {
  // increased window size
  uint32 window_size_increment = 1;
}
```

#### Close帧

Close帧分为正常的 Close帧 和 异常的 Close（Reset）帧。

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

## IDL

tRPC 的IDL基于Protobuf，用户编写Proto文件，通过trpc-cmdline生成相应的桩代码

### 一应一答RPC的IDL

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

### 流式RPC的IDL

流式IDL和普通RPC一样，使用stream关键字表明是流式RPC

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
注意关键字 stream，声明其为一个流方法。这里共涉及三个方法，对应关系为
- List：服务器端流式RPC
- Record：客户端流式RPC
- Route：双向流式RPC
