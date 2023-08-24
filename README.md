# tRPC -  A multi-language, plug-in, high-performance rpc framework

## What is tRPC

tRPC is a high-performance RPC development framework which is designed based on the concept of plug-in，the overall design follows the following principles:
- Easy to use: business developer should be simple and convenient for service development based on the framework;
- Intercommunication: Different services may use different communication protocols, and the framework should be able to solve the interoperability with these services;
- General & High Performance: The framework should be applicable to a broad class of application scenarios while sacrificing little in performance when compared to a specific application scenarios framework;
- Plugin: The entire framework should be layered and modular in architecture design and implementation, and each core module should preferably be pluggable and capable of independent evolution;

You can use it:
- Build services (trpc/http(s)/grpc, etc.) with multiple ports that support multiple protocols (one port only support one protocol), and can handle client requests synchronously/asynchronously;
- Access various protocol backend services (trpc/http(s)/grpc, etc.) in a synchronous, asynchronous, and one-way;
- Streaming RPC programming, currently supports trpc streaming, grpc streaming, http streaming, etc., to implement streaming application services such as Push, File Upload/Download, and AI Serving;
- Plug-in support for various protocols and service governance systems, such as: developing customized protocols, various name-service/metrics systems/tracing systems/config-center systems/log system, etc., to facilitate service interoperability and service operation;

## Features

- Works across languages：The communication between the client and the server is based on protobuf IDL, and the current open source version supports languages such as cpp/go.
- Support multi-protocols：A multi-protocol server that is convenient and fast to build. The protocols include common protocols (such as http/grpc), and other custom protocols.
- Stream RPC: In addition to supporting unary rpc call, tRPC also supports streaming rpc call, which better solves the application scenarios of large data transmission, such as: large file upload/download,message push,  speech recognition/video understanding, etc. 
- Rich plug-in ecology: tRPC provides plug-ins that are interoperable with mainstream ecosystems (such as grpc ecology/consul/promethues/opentelemetry, etc.), which is convenient for businesses to choose their own service governance ecology.
- Scalability: Based on the plug-in design of the framework, the business can use the plug-in factory and interceptor filter for personalized secondary development to expand the framework capabilities and meet business needs, such as: parameter validation, authentication, request replay, etc.
- Flow control and overload protection: In order to prevent services from being busy or overloaded due to sudden traffic surges or server failures, tRPC implements flow control and overload protection plug-ins in different application scenarios to ensure service availability;

## Supported languages

- [Cpp](https://github.com/trpc-group/trpc-cpp)
- [Go](https://github.com/trpc-group/trpc-go)

## To start using tRPC

Per-language quickstart guides and tutorials can be found in the
[documentation section on the trpc.io website](https://trpc.group/docs/). Code
examples are available in the examplesdirectory.

- tRPC
    -  [architecture_design](https://github.com/trpc-group/trpc/blob/main/docs/en/architecture_design.md)
    -  [terminology](https://github.com/trpc-group/trpc/blob/main/docs/en/terminology.md)
    -  [plugin_ecosystem](https://github.com/trpc-group/trpc/blob/main/docs/en/plugin_ecosystem.md)
    -  [trpc_protocol](https://github.com/trpc-group/trpc/blob/main/docs/en/trpc_protocol_design.md)
- tRPC-Cpp:
    - [quick_start](https://github.com/trpc-group/trpc-cpp/blob/main/docs/en/quick_start.md)
    - [basic_tutorial](https://github.com/trpc-group/trpc-cpp/blob/main/docs/en/basic_tutorial.md)
    - [user_guide](https://github.com/trpc-group/trpc-cpp/tree/main/docs)
    - [examples](https://github.com/trpc-group/trpc-cpp/tree/main/examples)
- tRPC-Go:
    - [quick_start]()
    - [basic_tutorial]()
    - [user_guide](https://github.com/trpc-group/trpc-go/tree/main/docs/README.md)
    - [examples](https://github.com/trpc-group/trpc-go/tree/main/examples)

## To start developing tRPC

Contributions are welcome!

Please read [How to contribute](https://github.com/trpc-group/trpc/blob/main/CONTRIBUTORS.md) which will guide you through
the entire workflow of how to build the source code, how to run the tests, and
how to contribute changes to the tRPC codebase.
