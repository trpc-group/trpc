English | [中文](README.zh_CN.md)

# tRPC -  A multi-language, pluggable, high-performance RPC framework

## What is tRPC

tRPC is a high-performance RPC framework that designed based on the concept of pluggable，the overall design follows the following principles:
- Simplicity: Users develop service very simply based on the framework.
- High Performance: The performant of the framework can be applicable to massive scenario on the internet.
- Pluggable: The framework is layered and modularized in architectural design and implementation, and each core module be pluggable and evolves independently.

By using tRPC, you can:
- Build services (trpc/http(s)/grpc, etc.) with multiple ports that support multiple protocols (one port only support one protocol), and can handle client requests synchronously/asynchronously.
- Access various protocol backend services (trpc/http(s)/grpc, etc.) in a synchronous, asynchronous, and one-way.
- Streaming RPC programming, currently supports trpc streaming, grpc streaming, http streaming, etc., to implement streaming application services such as Push, File Upload/Download, and AI Serving.
- Support various protocols and service governance systems, such as: customized protocols, various name-service/metrics systems/tracing systems/config-center systems/log system, etc., to facilitate service interoperability and operation.

## Features

- Works across languages：Implements cross-language service communication based on Protocol Buffers IDL.
- Support multi-protocols：Supports multiple communication protocols and interoperates with different frameworks (such as grpc).
- Streaming RPC: Supports streaming RPC, which is better suited for various application scenarios such as large file upload/download, message push, AI speech recognition/video understanding, etc.
- Rich plugin ecosystem: Provides a large number of plugins that docking to microservice components (such as Consul/Promethues/OpenTelemetry, etc.) to facilitate users to build their own service governance system.
- Scalability: Based on the pluggable design of the framework,  users can develop secondary to expand the framework capabilities, such as: parameter validation, authentication, log replay, etc.
- Flow control and overload protection: provides flow control and overload protection plugins in a variety of application scenarios to prevent services from being overloaded and unavailable due to burst traffic.

## Supported languages

- [Cpp](https://github.com/trpc-group/trpc-cpp)
- [Go](https://github.com/trpc-group/trpc-go)

## To start using tRPC

Per-language quickstart guides and tutorials can be found in the [tRPC website](https://trpc.group/docs/) . Code examples are available in the examples directory.

- tRPC
    - [architecture design](https://github.com/trpc-group/trpc/blob/main/docs/en/architecture_design.md)
    - [terminology](https://github.com/trpc-group/trpc/blob/main/docs/en/terminology.md)
    - [plugin ecosystem](https://github.com/trpc-group/trpc/blob/main/docs/en/plugin_ecosystem.md)
    - [trpc protocol](https://github.com/trpc-group/trpc/blob/main/docs/en/trpc_protocol_design.md)
- tRPC-Cpp:
    - [quick start](https://github.com/trpc-group/trpc-cpp/blob/main/docs/en/quick_start.md)
    - [basic tutorial](https://github.com/trpc-group/trpc-cpp/blob/main/docs/en/basic_tutorial.md)
    - [user guide](https://github.com/trpc-group/trpc-cpp/tree/main/docs)
    - [examples](https://github.com/trpc-group/trpc-cpp/tree/main/examples)
- tRPC-Go:
    - [quick start](https://github.com/trpc-group/trpc-go/blob/main/docs/quick_start.md)
    - [basic tutorial](https://github.com/trpc-group/trpc-go/blob/main/docs/basics_tutorial.md)
    - [user guide](https://github.com/trpc-group/trpc-go/tree/main/docs/README.md)
    - [examples](https://github.com/trpc-group/trpc-go/tree/main/examples)

## To start developing tRPC

Contributions are welcome!

Please read [How to contribute](https://github.com/trpc-group/trpc/blob/main/CONTRIBUTORS.md) which will guide you through the entire workflow of how to build the source code, how to run the tests, and how to contribute changes to the tRPC codebase.
