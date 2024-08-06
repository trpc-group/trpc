# Change Log

## ## [pb/go/trpc/relfection/v0.1.0](https://git.woa.com/trpc/trpc-protocol/tree/pb/go/trpc/reflection/v0.1.0)(2024-08-02)

- add reflection go stub code to support server reflection 

## [pb/go/trpc/v0.2.1](https://git.woa.com/trpc/trpc-protocol/tree/pb/go/trpc/v0.2.1) (2024-08-02)

### Features

- remove reflection package to fix https://git.woa.com/trpc-go/trpc-go/issues/987

## [pb/go/trpc/v0.2.0](https://git.woa.com/trpc/trpc-protocol/tree/pb/go/trpc/v0.2.0)(2024-07-02) 

### Features
- add openapiv2.proto support json_schema
- regenerate stub code to include LZ4_BLOCK_COMPRESS field
- add reflection.proto and corresponding go stub code to support server reflection !36

## [0.4.1](https://git.woa.com/trpc/trpc-protocol/tree/v0.4.0)(2024-04-02)
### Features And Improvements
- trpc流式协议新增流空闲关闭的错误码

### Thanks to our Contributors
irislao(劳昭丹)

## [0.4.0](https://git.woa.com/trpc/trpc-protocol/tree/v0.4.0)(2023-05-17)
### Features And Improvements
- 拓展trpc协议支持lz4解压缩类型

### Thanks to our Contributors
seadonli(李晓东)

## [pb/go/trpc/v0.1.1](https://git.woa.com/trpc/trpc-protocol/tree/pb/go/trpc/v0.1.1)(2023-05-09)
### Features
- 新增客户端读取 Frame 错误码 TRPC_CLIENT_READ_FRAME_ERR(171) 

## [pb/go/trpc/v0.1.0](https://git.woa.com/trpc/trpc-protocol/tree/pb/go/trpc/v0.1.0)(2023-04-19)
### Features
- 新增 trpc 协议

## [0.3.0](https://git.woa.com/trpc/trpc-protocol/tree/v0.3.0)(2023-02-23)
### Features And Improvements
- 修整trpc目录下各个proto文件的路径名以规范化
- 提供trpc/v2子目录以供 trpc-go trpc.tech v2 使用
- 增加java package选项
- 拓展trpc协议支持attachment

### Thanks to our Contributors
bochenchen(陈维翔);wineguo(郭琪周)

## [pb/go/trpc/v2.0.0](https://git.woa.com/trpc/trpc-protocol/tree/pb/go/trpc/v2.0.0-alpha)(2022-11-30)

- 修整 trpc 目录下各个 proto 文件的路径名以规范化
- 提供 trpc/v2 子目录以供 trpc-go trpc.tech v2 使用

## [0.2.1](https://git.woa.com/trpc/trpc-protocol/tree/v0.2.1)(2022-08-12)
### Features And Improvements
- snappy新增两种的压缩类型(block/stream)，以应对不同语言snappy的sdk实现差异
- 增加http的扩展定义HttpRule

## [0.2.0](https://git.woa.com/trpc/trpc-protocol/tree/v0.2.0)(2022-05-09)
### Features And Improvements
- 定义扩展option（trpc/proto/trpc_options.proto）
- 扩展option支持alias
- 为trpc协议的`TrpcContentEncodeType`新增枚举值`TRPC_THRIFT_ENCODE`
- 调整更合理的目录结构：从`proto`目录到`trpc/proto`目录

### Thanks to our Contributors
minchangwei(韦明昌)

## [0.1.0](https://git.woa.com/trpc/trpc-protocol/tree/v0.1.0)(2021-12-10)
### Features And Improvements
- 定义trpc一应一答协议
- 定义trpc流式协议
