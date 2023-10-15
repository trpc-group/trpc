# Overview

This article introduces the background of tRPC, pluggable design, and overall architecture.

# Background

In the early days of the development of the Internet, business scenarios varied greatly and iterations were rapid. This results in more or less differences in the language technology stacks, development frameworks, communication protocols, service governance systems, operation and maintenance platforms, etc. used by their backend services.

At the same time, with the development of cloud native technology, businesses are increasingly using open source technology and cloud components. Embracing cloud native has become a mainstream trend.

The above situation also exists within Tencent, and because of its large scale and diverse business types, it is more difficult to solve and must be solved. tRPC was born in this context.

# Pluggable Design

It must not only be interconnected with the existing technology system, but also adapt to cloud native technology and develop rapidly. This requires that the development framework must be open and extensible. In order to achieve this goal, tRPC adopts pluggable design ideas in architectural design.

The tRPC pluggable design ideas are as follows:

First, tRPC layers and modularizes the entire framework, and abstracts the core functional modules into independent plugins. The framework is then responsible for the concatenation and assembly of these independent plugins to achieve the features that the framework wants to support.

In terms of specific implementation, the following key technologies are used:
1. plugin factory based on interface mechanism;
2. filter based on AOP idea;

## Plugin Factory

The overall implementation idea of the plugin factory is that the framework only defines the standard interface of the plugin and provides registration capabilities without specific implementation. When interoperating with external services or connecting with a certain service governance system, you only need to develop the corresponding specific plugins.

For example: tRPC supports multiple protocols by defined a unified Codec. Different protocols only need to be implemented according to the Codec interface.

![codec](/docs/images/codec.png)

For example: tRPC connects to different naming service systems by defined unified Registry and Selector abstract interfaces. When connecting to different naming service systems, you only need to follow the Registry and Selector interfaces to implement it.

![naming](/docs/images/naming.png)

The design of the plugin factory can bring the following benefits:
- For the framework side: the framework only defines standard interfaces, without any plugin implementation, and is completely decoupled from the specific platform;
- For the platform side: You only need to implement the plugin according to the standard interface of the framework plugin, and integrate the capabilities of the platform into the framework;
- For the user side: developers only need to be used through configuration, which is transparent to users;

## Filter

In order to make the tRPC framework more scalable and meet personalized developers needs (such as metrics, log collection, tracing, parameter verification, request replay, fault injection, etc.), tRPC use java's aspect-oriented programming(AOP) idea to support filter.

The overall idea of filter implementation is to set up buried points in the framework's process of processing RPC requests, and then insert a series of filters into the buried points.

The filter workflow is as follows:

![fitler](/docs/images/filter.png)

The ultimate goal of filter is to decouple business logic from the framework and allow them to develop cohesively. It can dynamically add or replace personalized functions to business programs without modifying the framework code.

# Pluggable Archtecture

With the above key technical supporting, let’s take a look at how the tRPC pluggable architecture is designed.

## Overall Architectural Design

The overall architecture design of tRPC is as follows:

![architecture_design](/docs/images/architecture.png)

The overall architecture consists of two parts: "**Framework Core**" and "**Plugin**". The dotted line box is tRPC, the red solid line box in the middle is the core of the framework, and the blue box is the plugin part.

The core of the framework can be divided into three layers:

- **Communication Layer**: responsible for data transmission and protocol encoding and decoding. the framework has built-in support for communication protocols such as tcp and udp and uses the tRPC protocol based on Protobuf to carry RPC messages. It also supports other transmission protocols through codec plugins;

- **Service Governance Layer**: responsible for abstracting service governance functions into plugins and connecting them with service governance systems by calling plugins to realize service discovery, load balance, monitor, tracing, etc.

- **Call Layer**: encapsulates services and service proxy entities, provides RPC call interfaces, and supports synchronous, asynchronous, one-way and streaming calls;

In addition, the framework also provides an admin management interface, so that users or the operation platform can manage services by calling the admin interface. The management interface includes functions such as updating configurations, viewing versions, modifying log levels, viewing framework runtime information, etc. At the same time, the framework also supports user-defined management interfaces to meet business customization needs.

Plugins are the bridge that connects the framework core and external service governance systems. They are roughly divided into the following plugin types according to their functions:
- Codec: provides interfaces related to protocol encoding and decoding, allowing expansion of customized protocols, serialization methods, data compression methods through plugins;
- Naming: provides service registration (registry), service discovery (selector), load balance, circuit breaker and other capability encapsulation, used to connect with various naming service systems;
- Config: provides interfaces to read local configuration files, remote configuration center configurations, etc., allows plugin extensions to support configuration files in different formats, different configuration centers, and supports reload and watch configuration updates;
- Metrics: provides interfaces to report monitor data, supports common single-dimensional reporting, such as counter, gauge, etc., and also supports multi-dimensional reporting;
- Logging: Provides a general log collection interface, allowing the log implementation to be extended through plugins and output to remote locations;
- Tracing: Provides distributed tracing capabilities, allowing reporting to the call chain system through plugins;
- Telemetry: provides the ability to collect and report telemetry data. It is a plugin that integrates tracing, metrics and logging;

When implementing a specific plugin, one needs to implement the plugin according to the standard interface of the plugin, register it in the core of the framework, and complete plugin instantiation; on the other side, the specific plugin also needs to implement features (such as service discovery, load balance) based on using the SDK or API of the external service governance system.

## Specific Architecture Design

Before talking about the specific architecture design, let's first look at the process of RPC. From a developer's perspective, it allows you to make cross-node function call like local function call. Usually a complete RPC process is as follows:

![rpc](/docs/images/rpc.png)

The above figure describes the steps that an RPC call must go through. Based on this commonality, we divide the framework into server and client horizontally and layer it vertically. Among them, the vertical server is roughly divided into Transport, Codec, Service and Filter layers, and the client is divided into Transport, Codec, ServiceProxy and Filter layers. The specific design is as follows:

![rpc](/docs/images/layer.png)

In this picture, we have added a new layer of filter layer. Its main purpose is to use the idea of ​​AOP to meet customized needs (such as parameter verification, log replay, fault injection, etc.) and service governance (such as metrics, tracing, logging, authentication, etc.) that are inserted into the request/response processing process in a cross-cutting manner. This design enhances the framework's scalability.

At the same time, tRPC modularizes each layer and adopts pluggable implementation. The framework connects the entire process of RPC calls through the idea of interface-based programming. For some modules, the framework also adopts a more fine-grained module splitting. For example: the Selector module is subdivided into sub-modules such as service discovery, service routing, load balanc and etc., and the Codec layer is also subdivided into three sub-modules: encode/decode, serialization and compression.

Through the above overall layered design, pluggable implementation of specific modules and fine-grained module splitting, the framework has strong scalability and openness. Businesses can flexibly replace plugins to achieve connecting with different systems, and can also implement personalized capabilities.
