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

