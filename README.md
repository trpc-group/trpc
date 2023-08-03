# tRPC -  A multi-language, plug-in, high-performance rpc framework

## What is tRPC

tRPC is a high-performance RPC development framework which is designed based on the concept of plug-in，the overall design follows the following principles:
- Easy to use: business developer should be simple and convenient for service development based on the framework.
- Intercommunication: Different servers may use different message communication types, encoding and decompression methods，the framework design and implementation should be able to solve the intercommunication with these servers.
- General & High Performance: The framework should be applicable to a broad class of application scenarios while sacrificing little in performance when compared to a specific application scenarios framework.
- Plugin: The entire framework should be layered and modularized according to features in terms of architecture design and specific implementation. It is best for each core module to evolve independently. For capabilities related to service governance (such as naming/config/metrics/logging/tracing/auth, etc.), Extension points should be provided to allow plugging in these features and providing default implementations;

## Features
- Works across languages：The communication between the client and the server is based on protobuf IDL, and the current open source version supports languages such as c++/go/java.
- Support multi-protocols：A multi-protocol server that is convenient and fast to build. The protocols include common protocols (such as http/grpc), and other custom protocols.
- Stream RPC: In addition to supporting unary rpc call, tRPC also supports streaming rpc call, which better solves the application scenarios of large data transmission, such as: large file upload/download,message push,  speech recognition/video understanding, etc. 
- Rich plug-in ecology: tRPC provides plug-ins that are interoperable with mainstream ecosystems (such as grpc ecology/consul/promethues/opentelemetry, etc.), which is convenient for businesses to choose their own service governance ecology.
- Scalability: Based on the plug-in design of the framework, the business can use the plug-in factory and interceptor filter for personalized secondary development to expand the framework capabilities and meet business needs, such as: parameter validation, authentication, request replay, etc.
- Flow control and overload protection: In order to prevent services from being busy or overloaded due to sudden traffic surges or server failures, tRPC implements flow control and overload protection plug-ins in different application scenarios to ensure service availability;

## Supported languages

- [C++]()
- [Go]()
- [Java]()

## To start using tRPC

Per-language quickstart guides and tutorials can be found in the
[documentation section on the trpc.io website](https://trpc.group/docs/). Code
examples are available in the examplesdirectory.

- [C++](): [quick-start]() / [user-guide]() / [examples]()
- [Go](): [quick-start]() / [user-guide]() / [examples]()
- [Java](): [quick-start]() / [user-guide]() / [examples]()

## To start developing tRPC

Contributions are welcome!

Please read [How to contribute](CONTRIBUTING.md) which will guide you through
the entire workflow of how to build the source code, how to run the tests, and
how to contribute changes to the tRPC codebase.

## Troubleshooting

Sometimes things go wrong. Please check out the
[Troubleshooting guide](TROUBLESHOOTING.md) if you are experiencing issues with
tRPC.

## License

Apache License Version 2.0