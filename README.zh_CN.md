[English](README.md) | 中文

# tRPC -  多语言、插件化、高性能的RPC开发框架

## tRPC是什么

tRPC是基于插件化理念设计的一款支持多语言、高性能的RPC框架，整体设计遵循以下原则：
- 简单：用户基于框架进行服务开发简单方便；
- 高性能：框架具备高性能，能适用互联网海量访问场景；
- 插件化：框架在架构设计和具体实现上进行分层和模块化，各个核心模块可拔插，并能够独立演进；

你可以使用它：
- 搭建多个端口支持多个协议(一个端口只能对应一个协议)的服务(tRPC/HTTP(s)/gRPC等)，并能同步/异步处理客户端请求；
- 以同步、异步、单向的方式访问各种协议后端服务(tRPC/HTTP(s)/gRPC等)，调用各种存储系统(redis等)；
- 流式RPC编程，目前支持tRPC流式、gRPC流式、HTTP流式等，实现类似Push、文件上传/下载、AI类等流式应用服务；
- 插件化支持各种协议和对接服务治理系统，比如：开发自定义的协议、对接业务使用的各种名字服务/监控系统/调用链系统/配置系统/日志系统等，方便服务互通和服务运营；

## TRPC特点

- 跨语言：基于Protocol Buffers来实现跨语言之间的服务通信;
- 多通信协议：支持多种通信协议，能够与不同框架进行互通（比如gRPC）；
- 流式RPC：支持流式RPC，更好地适用于大文件上传/下载、消息Push、AI类语音识别/视频理解等多种应用场景；
- 丰富插件生态：提供大量对接业界微服务组件的插件（比如Consul/Promethues/Opentelemetry等），方便用户构建适合自己的服务治理体系；
- 可扩展性：基于框架插件化的设计，用户可以进行二次开发来扩展框架能力，比如：RPC请求参数校验、鉴权、请求录制等；
- 流控和过载保护：提供多种应用场景下的流量控制和过载保护插件，防止服务因为访问突增造成过载而不可用；

## 支持语言

- [Cpp](https://github.com/trpc-group/trpc-cpp)
- [Go](https://github.com/trpc-group/trpc-go)

## 如何使用tRPC

可以在 [tRPC官网](https://trpc.group/docs/) 找到tRPC每个语言的快速入门、基础教程等学习资料，也可以到tRPC各个语言的仓库查看详细的用户指南文档和代码示例。

- tRPC
    - [架构设计](https://github.com/trpc-group/trpc/blob/main/docs/zh/architecture_design.md)
    - [术语介绍](https://github.com/trpc-group/trpc/blob/main/docs/zh/terminology.md)
    - [插件生态](https://github.com/trpc-group/trpc/blob/main/docs/zh/plugin_ecosystem.md)
    - [tRPC协议](https://github.com/trpc-group/trpc/blob/main/docs/zh/trpc_protocol_design.md)
- tRPC-Cpp:
    - [快速开始](https://github.com/trpc-group/trpc-cpp/blob/main/docs/zh/quick_start.md)
    - [基础教程](https://github.com/trpc-group/trpc-cpp/blob/main/docs/zh/basic_tutorial.md)
    - [用户指南](https://github.com/trpc-group/trpc-cpp/blob/main/docs/README.zh_CN.md)
    - [代码示例](https://github.com/trpc-group/trpc-cpp/tree/main/examples)
- tRPC-Go:
    - [快速开始](https://github.com/trpc-group/trpc-go/blob/main/docs/quick_start.zh_CN.md)
    - [基础教程](https://github.com/trpc-group/trpc-go/blob/main/docs/basics_tutorial.zh_CN.md)
    - [用户指南](https://github.com/trpc-group/trpc-go/tree/main/docs/README.zh_CN.md)
    - [代码示例](https://github.com/trpc-group/trpc-go/tree/main/examples)

## 如何用wireshark分析tRPC协议

参考 [docs/zh/wireshark_trpc.md](docs/zh/wireshark_trpc.md)。

## 如何参与贡献

非常欢迎大家给tRPC做贡献!

建议您在为tRPC贡献之前, 先阅读一下 [如何参与贡献](https://github.com/trpc-group/trpc/blob/main/CONTRIBUTORS.zh_CN.md) , 它会指导你了解贡献代码的整个流程, 比如: 如何提PR/如何构建代码/如何运行单元测试等。

## 建议反馈

* 欢迎在[Issues](https://github.com/trpc-group/trpc/issues)中提出bug、疑惑、修改建议。