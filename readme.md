trpc统一协议的设计和定义

协议设计文档：[protocol_design.md](docs/protocol_design.md).

协议/常量/返回码定义：[trpc.proto](proto/trpc.proto).

协议中业务层在trans_info的key定义规范：  
1.框架使用的key以trpc_开头，具体常量定义在框架代码中统一管理；  
2.业务自定义key以app_开头，防止重复，不能以trpc开头；  
