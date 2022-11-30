trpc 统一协议的设计和定义

协议设计文档：[protocol_design.md](docs/protocol_design.md).

协议/常量/返回码定义：[trpc.proto](trpc/proto/trpc.proto).
拓展选项定义：[trpc_options.proto](trpc/proto/trpc_options.proto)

协议中业务层在 `trans_info` 的 key 定义规范：

- 框架使用的 key 以 `trpc_` 开头，具体常量定义在框架代码中统一管理；
- 业务自定义 key 以 `app_` 开头，防止重复，不能以 `trpc_` 开头；

目录结构:

```shell
trpc
├── api                        # 包含 restful 功能所需的 proto 文件
│   ├── annotations.proto
│   └── http.proto
├── proto                      # 包含 alias 功能所需的 proto 文件
│   └── trpc_options.proto
├── swagger                    # 包含 swagger 功能所需的 proto 文件
│   └── swagger.proto
├── trpc.proto                 # trpc 协议 proto 文件
├── validate                   # 包含 validate 功能所需的 proto 文件
|   └── validate.proto
└── v2                         # v2 子目录, 内部结构和 trpc 目录一致
    ├── api
    │   ├── annotations.proto
    │   └── http.proto
    ├── BUILD
    ├── proto
    │   └── trpc_options.proto
    ├── swagger
    │   └── swagger.proto
    ├── trpc.proto
    └── validate
        └── validate.proto
```

`v2` 目录的存在是为了应对 trpc-go 在 v0.x.x 和 v2.x.x 之间的共存问题, 由于 module name 的变更, v0.x.x 和 v2.x.x 本质上是两个仓库

这意味着假如他们仍然使用同一套 proto 文件桩代码时, 两个仓库在一个项目中被引用的话, proto 相关的信息就会重复注册导致 panic

为此, v2 版本的 proto 文件及其桩代码需要满足以下三个要求:

1. 生成桩代码时传给 protoc 的文件名（包含相对路径）共存时不能相同
2. proto 文件的 package name 不能相同，否则下面的字段在同一个 package 下就会有重复定义
3. 如果有自定义的 extension，那么 extension number 需要不同（trpc.proto 协议文件不涉及这个问题，annotations.proto 这些会涉及）

因此添加一个 `v2` 子目录并修改 proto package name 为 `trpc.v2` 可以解决上述问题

此外本仓库还作为与 rick 平台 import 路径的统一媒介, 详见文档: [tRPC-Go 代码生成插件 proto 依赖切换](https://doc.weixin.qq.com/doc/w3_AGkAxgZOAFMmQcAOcIkSOWgvNFkY3?scode=AJEAIQdfAAoLH9SMkiAGkAxgZOAFM)
