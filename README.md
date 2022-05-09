trpc 统一协议的设计和定义

协议设计文档：[protocol_design.md](docs/protocol_design.md).

协议/常量/返回码定义：[trpc.proto](trpc/proto/trpc.proto).
拓展选项定义：[trpc_options.proto](trpc/proto/trpc_options.proto)

协议中业务层在 `trans_info` 的 key 定义规范：

- 框架使用的 key 以 `trpc_` 开头，具体常量定义在框架代码中统一管理；
- 业务自定义 key 以 `app_` 开头，防止重复，不能以 `trpc_` 开头；
