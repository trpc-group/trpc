# 前言

本文主要介绍一下tRPC服务开发中的一些名词术语，包括服务命名、服务路由命名、接口命名等，方便大家对这些命名规则有统一的认识，避免产生误解。

# tRPC命名规范

通常线上提供的业务服务，会包括多个业务子系统，每个业务子系统内部又包括多个独立服务。这些服务之间的调用一般都是使用专门的路由名字间接寻址到被调用者的ip/port来进行访问的，而不是直接写死ip/port。因此，服务会先把自身的路由名字会注册到名字服务系统中，然后将其作为这个服务的唯一身份，方便主调服务进行识别，这就要求服务自身的命名能在众多微服务中具有唯一性。同时在微服务架构下，服务很多，如何通过服务的一些信息能快速识别某个服务，方便对服务日常的运维管理，比如：服务部署、监控、告警等，这也对服务的命名有一定的规范要求。

因此，为了能方便对服务开发、服务路由、服务运维等进行管理，tRPC对服务命名、路由命名、RPC接口命名做了一些规范设计（当然这些规范名称只是推荐，不做强制要求）。

## 服务命名

tRPC在服务命名上定义了以下3个纬度信息：
1. app名（应用名），表示某个业务系统的名称，用于标识某个业务下不同服务模块的一个集合；
2. seerver名（模块名），表示具体服务模块的名称，一般也称为模块的进程名称；
3. Service名，表示具体服务提供者的名称，一般使用proto文件定义的Service名称；
其中`app.server` 的组合在全局上要具备唯一性。

服务命名这样定义的好处是：
1. 对于服务开发，可以方便脚手架工具自动生成tRPC服务代码；
2. 对于服务寻址，可以生成全局唯一的服务路由名称，方便服务注册和服务发现；
3. 对于服务运营，可以方便服务的部署、各种维度的监控数据采集、告警等；

下图是服务A使用rpc调用服务B的流程图：

![rpc_workflow](/docs/images/rpc_workflow.png)

其中服务A作为客户端，服务B作为服务端，在tRPC中我们称客户端也为 `主调(caller)`，服务端也为 `被调(callee)`。主调（caller）A通过服务路由的名字（Route ID）向名字服务拿到ip/port列表后，然后在通信协议中设置调用的RPC接口名（Interface ID），访问被调B（callee）。

接下来，我们来看看服务路由命名和RPC接口命名的规则，以及具体服务时Service与服务路由名字直接的关系是怎样的。

## 服务路由命名

为了方便对服务的路由寻址，服务的路由命名规范如下：

由`.`号分割的四段字符串 **"trpc.{app}.{server}.{service}"** 来组成。

其中，第1段固定为`trpc`，表示这个服务是一个tRPC服务，第2段到第4段参考服务命名的含义。

另外，主调（caller）与被调（callee）的命名一般也与服务的路由命名一致。

## RPC接口命名

接下来我们看看tRPC中被调如何根据主调的消息信息调用相应RPC方法来处理请求的？这就涉及到主调方与被调方在消息通信协议上对RPC方法的命名有统一的规范.

下面是一个比较简单的服务提供的接口proto文件

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

基于proto文件的用法，tRPC在trpc协议上对RPC接口的名字制定了统一的规范，RPC调用方法名由proto文件中的 **"/{package_name}.{service_name}/{method_name}"** 组成。其中 `{package_name}` 的命名我们建议使用 `trpc.{app}.{server}`，`{service_name}` 为上面proto文件中定义的service名字，`{method_name}` 为service下具体的方法名，即我们要调用的RPC接口。

## Service与路由名字的映射

首先，这里Service指的是proto文件中定义的Service，它负责定义一系列RPC接口，Service实例是Service具体的对象，而路由名字则是进行服务寻址的名字。对外提供服务时，我们需要把三者关联起来，通过路由名字来找到具体proto的Service提供者。

tRPC建议一个服务端程序只提供一个proto Service，实例化一个Service实例，其对应一个路由名字，使用一种协议和一个port对外提供接口服务。 服务组装模式如图所示：

![single_proto_single_service_instance](/docs/images/single_proto_single_service_instance.png)

tRPC也支持一个服务端程序只提供一个proto Service，实例化多个Service实例，每个Service实例其对应一个路由名字，使用一种协议和一个port对外提供接口服务，不同Service实例的port不能一样。 服务组装模式如图所示：

![single_proto_multi_service_instance](/docs/images/single_proto_multi_service_instance.png)

tRPC也支持一个服务端程序只提供多个proto Service，不同proto Service可以实例化1个或者多个Service实例，每个Service实例其对应一个路由名字，使用一种协议和一个port对外提供接口服务，不同Service实例的port不能一样。 服务组装模式如图所示：

![multi_proto_multi_service_instance](/docs/images/multi_proto_multi_service_instance.png)

注意：tRPC不支持同一个端口支持多协议。
