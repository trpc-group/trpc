# Overview

This article mainly introduces some terms in tRPC service development, including service naming, service route naming, interface naming, etc., so that everyone can have a unified understanding of these naming rules and avoid misunderstandings.

# Naming Specification

Usually, business online service will include multiple business subsystems, and each business subsystem includes multiple independent services. Calls between these services generally use special route names to indirectly get the callee's IP/port for accessing, rather than directly use the hard-coding IP/port. Therefore, the service will first register its route name to the naming service system, and then use it as the unique identity of the service. This requires that the name of the service must be unique among many micro-services. At the same time, under the micro-service architecture, with so many services, how can a certain service be quickly identified through some information of the service, so as to facilitate the operation and maintenance management of the service, such as service deployment, monitoring, alarming, etc. This also has certain standard requirements for the naming of services.

Therefore, in order to facilitate the management of service development, service route, service operation and maintenance, etc., tRPC has made some standardized designs for service naming, route naming, and RPC interface naming (of course, these standardized names are only recommendations and are not mandatory).

## Service Naming

tRPC defines the following three dimensions of service naming:
1. app name, which represents the name of a certain business system and is used to identify a collection of different service modules;
2. server name (module name), which represents the name of the specific service module, generally also called the process name;
3. service name, which represents the name of the specific service provider. Generally, the service name defined in the proto file;

The benefits of service naming like this are:
1. For service development, scaffolding tools can be used to automatically generate tRPC service code;
2. For service addressing, a globally unique service route name can be generated to facilitate service registration and service discovery;
3. For service operation, it can facilitate service deployment, monitor data collection, alarms, etc.;

The following figure is a flow chart of service A using rpc to call service B:

![rpc_workflow](/docs/images/rpc_workflow.png)

Among them, service A serves as the client and service B serves as the server. In tRPC, we call the client also the caller and the server the callee. After the caller A obtains the ip/port list from the naming service through the name of the service route (Route ID), it then sets the callee RPC interface name (Interface ID) in the communication protocol and accesses the callee B.

Next, let's take a look at the rules for service route naming and RPC interface naming, as well as the direct relationship between Service and service route name.

## service route naming

In order to facilitate addressing of the service, the service route naming convention is as follows:

It is composed of four segments of string **"trpc.{app}.{server}.{service}"** separated by `.`.

Among them, the first paragraph is fixed to `trpc`, indicating that this service is a tRPC service. The second to fourth paragraphs refer to the meaning of service naming.

The combination of `trpc.{app}.{server}` must be globally unique. In addition, the name of caller and callee are generally consistent with the service route name.

## RPC Interface Naming

Next, let's take a look at how the callee calls the corresponding RPC method to process the request based on the message information of the caller in tRPC? This involves the caller and the callee have unified specifications for the naming of RPC methods in the message communication protocol.

The following is a simple interface proto file provided by the service

```protobuf
syntax = "proto3";

package trpc.test.helloworld;

service Greeter {
  rpc SayHello (HelloRequest) returns (HelloReply) {}
}

message HelloRequest {
   string msg = 1;
}

message HelloReply {
   string msg = 1;
}

```

Based on the usage of the protobuf, tRPC has formulated a unified specification for the name of the RPC interface on the trpc protocol. The RPC method name is composed of **"/{package_name}.{service_name}/{method_name}"** in the proto file.

Among them, we recommend using `trpc.{app}.{server}` for the name of `{package_name}`, `{service_name}` is the service name defined in the proto file above, and `{method_name}` is the specific method name under the service , that is the RPC interface we want to call.


## Mapping of service and service route name

First of all, Service here refers to the Service defined in the proto file, which is responsible for defining a series of RPC interfaces. The Service instance is the specific object of the Service, and the service route name is the name for service addressing. When providing services, we need to associate the three and find the service provider through the service route name.

tRPC recommends that a server program only provide one proto service and Service instance, which corresponds to a servie route name, and use one protocol and one port to provide interface service. The service assembly model is shown in the figure:

![single_proto_single_service_instance](/docs/images/single_proto_single_service_instance.png)

tRPC also supports a server program to provide only one proto service and instantiate multiple service instances. Each Service instance corresponds to a service route name and uses one protocol and one port to provide interface service. The ports of different service instances cannot be the same. The service assembly model is shown in the figure:

![single_proto_multi_service_instance](/docs/images/single_proto_multi_service_instance.png)

tRPC also supports a server program to provide multiple proto services. Different proto services can instantiate one or more service instances. Each service instance corresponds to service route name and uses one protocol and one port to provide interface service. The ports of different service instances cannot be the same. The service assembly model is shown in the figure:

![multi_proto_multi_service_instance](/docs/images/multi_proto_multi_service_instance.png)

Note: tRPC does not support multiple protocols on the same port.

