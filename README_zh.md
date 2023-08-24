# tRPC -  多语言、插件化、高性能的rpc开发框架

## tRPC是什么

tRPC是基于插件化理念设计的一款支持多语言、高性能的rpc开发框架，整体设计遵循以下原则：
- 使用简单：业务开发基于框架进行服务开发应该简单方便；
- 互通：不同的服务可能使用不同的通信协议，框架设计和实现上应能解决与这些服务的互通；
- 通用且高性能：框架应该适用于绝大多数应用场景，相比针对特定应用场景的框架，框架只会牺牲一点性能；
- 插件化：整个框架在架构设计和具体实现上应该进行分层和模块化，各个核心模块最好可拔插，并能够独立演进；

你可以使用它：
- 搭建多个端口支持多个协议(一个端口只能对应一个协议)的服务(trpc/http(s)/grpc等)，并能同步/异步处理客户端请求；
- 以同步、异步、单向的方式访问各种协议后端服务(trpc/http(s)/grpc等)，调用各种存储系统(redis等)；
- 流式rpc编程，目前支持trpc流式、grpc流式、http流式等，实现类似push、文件上传/下载、ai类等流式应用服务；
- 插件化支持各种协议和对接服务治理系统，比如：开发自定义的协议、对接业务使用的各种名字服务/监控系统/调用链系统/配置系统/日志系统等，方便服务互通和服务运营；

## TRPC特点

- 跨语言：客户端与服务端通信基于protobuf IDL，目前开源版本支持c++/go等语言；
- 多通信协议：方便快速构建的多协议服务， 协议包括通用协议http/grpc，以及各种自定义协议；
- 流式rpc： 除了支持一应一答的rpc调用方式外，tRPC还支持流式rpc调用，更好地解决大文件上传/下载、消息push、ai类语音识别/视频理解等大数据传输的应用场景；
- 丰富插件生态：tRPC提供了与业界主流生态互通的插件（比如：grpc生态/consul/promethues/opentelemetry等），方便业务选择适合自己的服务治理生态；
- 可扩展性：基于框架插件化的设计，业务可以使用插件工厂和拦截器filter进行个性化的二次开发来扩展框架能力，满足业务需求，比如：RPC请求参数校验、鉴权、请求录制等；
- 流控和过载保护：为了防止业务因为访问量突增或服务器故障造成服务整体的繁忙或者过载不可用，tRPC实现了不同应用场景下的流量控制和过载保护插件，保证服务的可用性；

## 支持语言

- [Cpp](https://github.com/trpc-group/trpc-cpp)
- [Go](https://github.com/trpc-group/trpc-go)

## 如何使用tRPC

可以在[documentation section on the trpc.group website](https://trpc.group/docs/)找到tRPC每个语言的快速入门、基础教程等学习资料，也可以到tRPC各个语言的仓库查看详细的用户指南文档和代码示例.

- tRPC
    -  [架构设计](https://github.com/trpc-group/trpc/blob/main/docs/en/architecture_design.md)
    -  [术语介绍](https://github.com/trpc-group/trpc/blob/main/docs/en/terminology.md)
    -  [插件生态](https://github.com/trpc-group/trpc/blob/main/docs/en/plugin_ecosystem.md)
    -  [trpc协议](https://github.com/trpc-group/trpc/blob/main/docs/en/trpc_protocol_design.md)
- tRPC-Cpp:
    - [quick-start](https://github.com/trpc-group/trpc-cpp/blob/main/docs/zh/quick_start.md)
    - [basic—tutorial](https://github.com/trpc-group/trpc-cpp/blob/main/docs/zh/basic_tutorial.md)
    - [user-guide](https://github.com/trpc-group/trpc-cpp/tree/main/docs)
    - [examples](https://github.com/trpc-group/trpc-cpp/tree/main/examples)
- tRPC-Go:
    - [quick-start]()
    - [basic—tutorial]()
    - [user-guide](https://github.com/trpc-group/trpc-go/tree/main/docs/README.md)
    - [examples](https://github.com/trpc-group/trpc-go/tree/main/examples)

## 如何为tRPC做贡献

非常欢迎大家给tRPC做贡献!

建议您在为tRPC贡献之前, 先阅读一下[How to contribute](https://github.com/trpc-group/trpc/blob/main/CONTRIBUTORS_zh.md), 它会指导你了解贡献代码的整个流程, 比如: 如何提pr/如何构建代码/如何运行单元测试等;
