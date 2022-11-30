# Change Log

## [pb/go/trpc/v2.0.0](https://git.woa.com/trpc/trpc-protocol/tree/pb/go/trpc/v2.0.0)(2022-11-30)

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
