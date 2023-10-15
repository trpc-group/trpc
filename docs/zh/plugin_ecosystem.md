# 前言

本文主要介绍一下tRPC目前的插件生态建设情况。

# tRPC插件的成熟度分类

tRPC插件在成熟度上，目前分为以下4级：

| 成熟度   | 描述                             |
| -------- | -------------------------------- |
| stable   | 已经大规模使用，放心使用         |
| on trial | 已测试，有一定业务用过，基本靠谱 |
| tested   | 已测试，还没有业务用过，可能有坑 |
| archived  | 已归档，不再维护，不推荐使用 |

# tRPC各语言的插件生态建设情况

目前tRPC在插件生态的建设上，插件类型主要包含协议、拦截器、名字服务系统、配置中心系统、监控系统、调用链系统、远程日志系统、可观测系统等。

## tRPC-Cpp插件列表

| 插件类型  |   插件   |  成熟度  |  文档 |
| -------- | -------- | :------: | :---------------------------------------------------------------------------------: |
| 协议  | HTTP |  stable  | [链接](https://github.com/trpc-group/trpc-cpp/blob/main/docs/zh/http_protocol_service.md) |
| 协议  | gRPC |  stable  | [链接](https://github.com/trpc-group/trpc-cpp/blob/main/docs/zh/grpc_protocol_service.md) |
| 序列化  | Protobuf |  stable  | [链接](https://github.com/trpc-group/trpc-cpp/blob/main/docs/zh/serialization.md) |
| 序列化  | JSON |  stable  | [链接](https://github.com/trpc-group/trpc-cpp/blob/main/docs/zh/serialization.md) |
| 序列化  | text |  stable  | [链接](https://github.com/trpc-group/trpc-cpp/blob/main/docs/zh/serialization.md) |
| 序列化  | binary |  stable  | [链接](https://github.com/trpc-group/trpc-cpp/blob/main/docs/zh/serialization.md) |
| 解压缩  | Gzip |  stable  | [链接](https://github.com/trpc-group/trpc-cpp/blob/main/docs/zh/compression.md) |
| 解压缩  | lz4 |  stable  | [链接](https://github.com/trpc-group/trpc-cpp/blob/main/docs/zh/compression.md) |
| 解压缩  | Snappy |  stable  | [链接](https://github.com/trpc-group/trpc-cpp/blob/main/docs/zh/compression.md) |
| 解压缩  | zlib |  stable  | [链接](https://github.com/trpc-group/trpc-cpp/blob/main/docs/zh/compression.md) |
| 名字服务系统  | MeshPolaris |  tested  | [链接](https://github.com/trpc-group/cpp-naming-polarismesh/blob/main/README_zh.md) |
| 配置中心系统  | etcd |  tested  | [链接](https://github.com/trpc-group/cpp-config-etcd/blob/main/README_zh.md) |
| 监控系统  | Prometheus |  stable  | [链接](https://github.com/trpc-group/trpc-cpp/blob/main/docs/zh/prometheus_metrics.md) |
| 调用链系统  | Jaeger |  stable  | [链接](https://github.com/trpc-group/cpp-tracing-jaeger/blob/main/README_zh.md) |
| 远程日志系统  | CLS |  stable  | [链接](https://github.com/trpc-group/cpp-logging-cls/blob/main/README_zh.md) |
| 可观测系统  | OpenTelemetry |  stable  | [链接](https://github.com/trpc-group/cpp-telemetry-opentelemetry/blob/main/README_zh.md) |

## tRPC-Go插件列表

| 插件类型  |   插件   |  成熟度  |  文档 |
| -------- | -------- | :------: | :---------------------------------------------------------------------------------: |
| 协议  | HTTP |  stable  | [链接](https://github.com/trpc-group/trpc-go/tree/main/http) |
| 协议  | gRPC |  stable  | [链接](https://github.com/trpc-ecosystem/go-codec/tree/main/grpc) |
| 名字服务系统  | MeshPolaris |  tested  | [链接](https://github.com/trpc-ecosystem/go-naming-polarismesh) |
| 配置中心系统  | etcd |  tested  | [链接](https://github.com/trpc-ecosystem/go-config-etcd) |
| 监控系统  | Prometheus |  stable  | [链接](https://github.com/trpc-ecosystem/go-metrics-prometheus) |
| 调用链系统  | Jaeger |  stable  | [链接](https://github.com/trpc-ecosystem/go-opentracing-jaeger) |
| 远程日志系统  | CLS |  stable  | [链接](https://github.com/trpc-ecosystem/go-log-cls) |
| 可观测系统  | OpenTelemetry |  stable  | [链接](https://github.com/trpc-ecosystem/go-opentelementry) |
| 拦截器 | debuglog | stable | [链接](https://github.com/trpc-ecosystem/go-filter/tree/main/debuglog) |
| 拦截器 | degrade（服务端熔断限流）| stable | [链接](https://github.com/trpc-ecosystem/go-filter/tree/main/degrade) |
| 拦截器 | filterextensions（方法级别的拦截器）| stable | [链接](https://github.com/trpc-ecosystem/go-filter/tree/main/filterextensions) |
| 拦截器 | hystrix（服务端熔断限流）| stable | [链接](https://github.com/trpc-ecosystem/go-filter/tree/main/hystrix) |
| 拦截器 | jwt (认证鉴权）| stable | [链接](https://github.com/trpc-ecosystem/go-filter/tree/main/jwt) |
| 拦截器 | masking（敏感数据脱敏）| stable | [链接](https://github.com/trpc-ecosystem/go-filter/tree/main/masking) |
| 拦截器 | mock（client mock）| stable | [链接](https://github.com/trpc-ecosystem/go-filter/tree/main/mock) |
| 拦截器 | recover（捕获 panic 并恢复）| stable | [链接](https://github.com/trpc-ecosystem/go-filter/tree/main/recovery) |
| 拦截器 | referer（HTTP Referer 安全校验）| stable | [链接](https://github.com/trpc-ecosystem/go-filter/tree/main/referer) |
| 拦截器 | slime（重试/对冲）| stable | [链接](https://github.com/trpc-ecosystem/go-filter/tree/main/slime) |
| 拦截器 | transinfo-blocker（元数据透传屏障）| stable | [链接](https://github.com/trpc-ecosystem/go-filter/tree/main/transinfo-blocker) |
| 拦截器 | tvar（统计监控）| stable | [链接](https://github.com/trpc-ecosystem/go-filter/tree/main/tvar) |
| 拦截器 | validation（参数自动校验）| stable | [链接](https://github.com/trpc-ecosystem/go-filter/tree/main/validation) |
