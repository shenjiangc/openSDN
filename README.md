# openSDN
software define network

为超大规模SDN提供高性能、可扩展sdn控制器。
1、数据集选用ScyllaDB、Cassandra列式数据库
2、通过gocql与数据库交互
3、产品概念参考阿里云
4、北向提供REST API，gRPC
5、南向提供gRPC
6、南向协议相关层如OPENFLOW，由各自的适配层实现。
7、分层架构查看ONOS
8、VPC控制器、NAT控制器、LB控制器等分别实现

project/
├── cmd/
│   ├── main.go
│   └── ...
├── internal/
│   ├── pkg1/
│   │   ├── file1.go
│   │   └── ...
│   └── pkg2/
│       ├── file2.go
│       └── ...
├── pkg/
│   ├── file3.go
│   └── ...
├── vendor/
│   └── ...
├── api/
│   ├── proto/
│   │   ├── service1.proto
│   │   └── ...
│   └── service1/
│       ├── server.go
│       ├── client.go
│       ├── handler.go
│       └── ...
├── web/
│   ├── static/
│   │   ├── css/
│   │   ├── js/
│   │   └── ...
│   ├── templates/
│   ├── server.go
│   └── ...
├── configs/
│   ├── config1.yaml
│   ├── config2.json
│   └── ...
├── scripts/
│   ├── script1.sh
│   ├── script2.py
│   └── ...
├── tests/
│   ├── pkg1/
│   │   ├── file1_test.go
│   │   └── ...
│   ├── pkg2/
│   │   ├── file2_test.go
│   │   └── ...
│   └── ...
├── docs/
│   ├── README.md
│   ├── design.md
│   └── ...
├── Makefile
└── go.mod

cmd/目录：包含所有可执行文件的源代码，每个子目录代表一个不同的可执行文件。
internal/目录：包含项目的内部库和工具，只能在项目内部使用，不能被外部代码导入。内部库可以根据需要进一步划分子目录，每个子目录代表一个独立的内部库。
pkg/目录：包含可导出的库代码，可以被其他项目导入。pkg目录也可以根据需要进一步划分子目录，每个子目录代表一个独立的库。
vendor/目录：包含项目依赖的所有第三方包，用于解决依赖管理问题。
api/目录：包含项目的API代码，包括协议定义和服务实现。
web/目录：包含Web应用程序的代码，包括静态文件和模板文件。
configs/目录：包含配置文件，包括YAML、JSON等格式。
scripts/目录：包含与项目相关的脚本文件，包括部署脚本、自动化测试脚本等。
tests/目录：包含所有单元测试和集成测试代码。
docs/目录：包含与项目相关的文档，包括设计文档、API文档等。
Makefile文件：包含项目构建和部署的Makefile规则。
go.mod文件：包含项目的模块定义，指定了项目的依赖关系和版本信息。


https://github.com/scylladb/scylla-cdc-go




# golang技术选型

- [v]日志：https://github.com/uber-go/zap
- [v]CLI命令： https://github.com/spf13/cobra
- [v]配置文件：github.com/spf13/viper （支持etcd）
- [v]字段参数验证： https://github.com/go-playground/validator
~~ - [X]路由库：https://github.com/gorilla/mux (不选，1：项目停止 2：go-restful自带mux) ~~
~~ - [X]gorm: gorm.io/gorm (不需要) ~~
- [v]cql:  
    - gocql : https://github.com/gocql/gocql 
    - gocqlx : https://github.com/scylladb/gocqlx
- []rpc: https://github.com/smallnest/rpcx （支持etcd,性能高于grpc）
- []rpcx-plugin-cql : 给rpcx增加cql插件，以支持cassandra & scylladb
- [X]register： etcd
- []OPENFLOW： cgo调用ovs自带openflow
- [v]RESTful/openapi/Swagger: https://github.com/emicklei/go-restful
- [v]Swagger-ui: swagger-ui-dist


- [] openapi
- [] scylladb的使用         需要有存储/数据库经验
- [] gocql的使用            
- [] rpcx的使用             需要有网络编程基础   给rpcx增加cql插件，以支持cassandra & scylladb
- [] flow设计               需要精通openflow协议



数据库： aerospike          
        scylladb
        Cassandra
        mongodb

