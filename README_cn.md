# tRPC -  多语言、插件化、高性能的rpc开发框架

## tRPC是什么

tRPC是基于插件化理念设计的一款高性能RPC开发框架，整体设计遵循以下原则:
- 使用简单：业务开发基于框架进行服务开发应该简单方便
- 互通：不同的服务可能使用不同的消息通信类型、编解码和解压缩方式，框架设计和实现上应能解决与这些服务的互通；
- 通用且高性能：框架应该适用于绝大多数应用场景，相比针对特定应用场景的框架，框架只会牺牲一点性能；
- 插件化：整个框架在架构设计和具体实现上应该按特性进行分层和模块化，各个核心模块最好能够独立演进，对于服务治理相关的能力（比如naming/config/metrics/logging/tracing/auth等），应该提供扩展点，以允许插入这些特性和提供默认实现；

## 特点
- 跨语言：客户端与服务端通信基于protobuf IDL，目前开源版本支持c++/go/java等语言；
- 多通信协议：方便快速构建的多协议服务， 协议包括通用协议http/grpc，以及各种自定义协议；
- 流式rpc： 除了支持一应一答的rpc调用方式外，tRPC还支持流式rpc调用，更好地解决大文件上传/下载、消息push、ai类语音识别/视频理解等大数据传输的应用场景；
- 丰富插件生态：tRPC提供了与业界主流生态互通的插件（比如grpc生态/consul/promethues/opentelemetry等），方便业务选择适合自己的服务治理生态；
- 可扩展性：基于框架插件化的设计，业务可以使用插件工厂和拦截器filter进行个性化的二次开发来扩展框架能力，满足业务需求，比如：RPC请求参数校验、鉴权、请求录制等；
- 流控和过载保护：为了防止业务因为访问量突增或服务器故障造成服务整体的繁忙或者过载不可用，tRPC实现了不同应用场景下的流量控制和过载保护插件，保证服务的可用性；

## 支持语言

- [C++]()
- [Go]()
- [Java]()

## 如何使用tRPC

可以在[documentation section on the trpc.io website](https://trpc.group/docs/)找到每个语言的快速入门、用户指南和代码示例等学习资料.

- [C++](): [quick-start]() / [user-guide]() / [examples]()
- [Go](): [quick-start]() / [user-guide]() / [examples]()
- [Java](): [quick-start]() / [user-guide]() / [examples]()

## 如何为tRPC做贡献

非常欢迎大家给tRPC做贡献!

建议您在为tRPC贡献之前, 先阅读一下[How to contribute](CONTRIBUTING.md), 它会指导你了解贡献代码的整个流程, 比如: 如何提pr/如何构建代码/如何运行单元测试等;

## 遇到问题怎么办

如果使用tRPC过程遇到问题, 可以先看看[Troubleshooting guide](TROUBLESHOOTING.md), 它会告诉你使用tRPC遇到各种问题的处理方式, 比如: 如何提issue/一些常见问题如何排查等;

## License

Apache License Version 2.0