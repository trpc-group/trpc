该仓库维护了一些trpc框架脚手架工具使用的pb扩展：
- trpc.ext.proto，定义了rpc service methodoption，允许定义别名；
- swagger.proto，定义了生成swagger apidocs时的一些swagger描述；
- validate.proto，定义了protoc-gen-secv依赖的一些校验规则；
- api/annotations.proto，定义了RESTful api需要的扩展option

## Examples

下面是几个使用示例。

### 自定义名称：rpc方法、go tag
```protobuf

import "trpc.proto";

// 自定义rpc方法名 by `trpc.alias` method option
service HelloSvr {
    // Here, trpc.alias will build the bridge between `cmd` 0x1000/0x1 and handlefunc `Hello`.
    // Actually, trpc framework will pass the option `trpc.alias = "ilive/0x1000/0x1"` to `codec`,
    // then codec will split "/ilive/0x1000/0x1" as `bigcmd=0x1000` and `subcmd=0x1".
    //
    // By this way, we can generally support nearly all kinds of existed protocols in our company.
    rpc Hello(HelloReq) returns(HelloRsp) {
        option (trpc.alias) = "/ilive/0x1000/0x1";
    }
}

// 自定义go tag by `trpc.tag` field option
message SayHello {
    string name = 1 [(trpc.go_tag)='xml:"name"'];
}
```

### 定义 RESTful API
```protobuf

syntax = "proto3";
package helloworld;

import "trpc/api/annotations.proto";

option go_package="git.code.oa.com/examples/helloworld";

// HelloReq request
message HelloReq{}

// HelloRsp response
message HelloRsp{}

service helloworld_svr {
    // Hello say hello
    rpc Hello(HelloReq) returns(HelloRsp) {
        option (trpc.api.http) = {
              get: "/say/{name}"
              additional_bindings: {
                    get: "/say/strval/{strVal}",
              }
        };
    };
}

```

## 注意事项

protoc、protoc-gen-go生成go代码时会检测同一个pb文件名重复注册的问题，未来版本可能会直接panic，
为避免此问题，特意将trpc.proto重命名为trpc.ext.proto，以避免与trpc默认协议文件trpc.proto冲突。

在使用工具trpc-go/trpc-go-cmdline进行代码生成时，import pb文件时，依然写`import "trpc.proto"`，
重命名trpc.ext.proto，这么做只是为了避免生成的桩代码panic，pb引用时写trpc.proto不受影响。
