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




./kafka-topics.sh --bootstrap-server 127.0.0.1:9092 --delete --topic  first
标记删除
first-0.8a1249ffafc8464b9df06d8d424b8f88-delete/



./kafka-topics.sh --bootstrap-server 127.0.0.1:9092 --create --topic first  --partitions 2 --replication-factor
 2
 
 
 ./kafka-topics.sh --bootstrap-server 127.0.0.1:9092 --describe  first


 修改分区
 ./kafka-topics.sh --bootstrap-server 127.0.0.1:9092 --alter  --topic  first  --partitions 2
 
 
 
```


### 关于topic的增删查改
1. 增 

```shell
kafka-topics --bootstrap-server xx:9092 --create --topic name 
--partitions xx --replication-factor xx 

```

2. 查

```shell
kafka-topics --bootstrap-server 9092 --describe xx 

```



# 总结

server.properties

broker.id

log.directory



## 生产者消费者 


可以消费一个不存在的数据
发送到一个不存在的主题  

默认会创建

--from-beginning 


![1F2DfA](https://raw.githubusercontent.com/jacksonyoudi/images/main/uPic/1F2DfA.png)


1. 消费者topic
50个分区，一个副本 

保存消费者的 消费message的偏移量




2. 数据保存


索引
数据
分段

```shell
      00000000000000000000.index       00000000000000000000.timeindex
      00000000000000000000.log         leader-epoch-checkpoint
```

xxx.index 索引  索引开头 偏移开头 第几条消息 


```
1. 当数据比较大了， 就会分段  
二分查找，为了 更快查询

log.segment.bytes=1073741824  1G 

   
2.  保存数据是   1 : nnn,mmm 

message位置和长度 

3. 每个段都是 从 0 开始， 和文件名进行加减
```



xx.log  数据
xx.timeindex 时间索引  七天 


```shell
Topic: first	PartitionCount: 2	ReplicationFactor: 1	Configs: segment.bytes=1073741824
	Topic: first	Partition: 0	Leader: 0	Replicas: 0	Isr: 0
	Topic: first	Partition: 1	Leader: 0	Replicas: 0	Isr: 0
```


读写 很快：

1. 顺序写
2. 索引读


分片 和 索引机制



### 日志和数据分离

1. 先停kafka
2. 先删 logs
3. 修改配置
4. 清理zk信息
    
    1. deleteall /kafka
    2. 到zk数据目录下直接删除
    








