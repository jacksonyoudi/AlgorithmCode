## 安装

broker.id 修改



```
 ./zkServer.sh start
 
 bin/kafka-server-start.sh -daemon  config/server.properties

 
```



## topic



```shell
--bootstrap-server 必须要

./kafka-topics.sh --bootstrap-server 127.0.0.1:9092 --list



./kafka-topics.sh --bootstrap-server 127.0.0.1:9092 --create --topic first
```