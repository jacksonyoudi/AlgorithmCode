# producer 

副本有leader， 读写都是leader 操作， follower


```shell
Topic: first	PartitionCount: 2	ReplicationFactor: 1	Configs: segment.bytes=1073741824
	Topic: first	Partition: 0	Leader: 0	Replicas: 0	Isr: 0
	Topic: first	Partition: 1	Leader: 0	Replicas: 0	Isr: 0


Leader: broker.id 
```

### 分区原因
1. 方便集群扩展 灵活
2. 可以提高并发 


### 分区原则
 
消息封装成一个 `ProducerRecord` 对象

1. 指定分区
2. 指定key
3. 粘性分区器，会随机选择一个分区，并尽可能一直使用该分区



### 数据可靠性

1. producer发送数据到topic partion的可靠性保证
ack机制
   
策略： 
    1. leader未落盘 异步
    2. leader落盘   半同步
    3. follower落盘 同步



ISR： in-sync replica set(ISR) 动态维护， 和leader保持同步的follower集合。
默认10s， 10s没有同步完成，就踢出 (早期版本 消息条数 过半， 频繁进进出出)

OSR： out sync replica


#### ack机制
acks:
0
1:  leader挂了 
-1： 未ack， producer会重发， 会重复  ISR只有一个leader，可能会丢数据


### exectly Once

acks -1 可以去重 

at least Once + 幂等性 = exectly Once

要启用幂等性 只需要将 producer的参数中 enable.idempotence设置为true即可。
pid partition seqnumer做缓存  seqnumer 消息的序列化id
不能跨分区和会话 需要事务



at most once
ad least once




# zk 

元数据信息 
0.9 版本 之前offset 存储zk


# consumer 

消费者组
一个消费者消费一个 partition

消费者数量和partion 数量相同的时候， 效率最大 

消费者偏移量是 记录消费者组的 offset


## 消费方式 
    pull 拉取方式
    
    问题： kafka中没有数据， 消费者可能会陷入循环中，一直返回空数据。
    timeout

## 分区分配策略
    
    1. roundrobin 
    2. range 



roundrobin: 以消费者组为单位，将消费者里订阅的所有主题全部取出来，将主题的每个分区进行排序，再对消费者者进行轮询
，有可能导致消费者消费不是自己订阅的主题里的数据，但是完全可以避免的。


range 范围分配，以主题为单位，如果消费者没有订阅此主题，就不会收到数据，不会导致数据错放问题。但是有可能导致消费者
组里 分区分配不均。 




### offset

1. consumer.properties

exclude.internal.topics=false


##  指定消费者组

```shell
vim consumer.properties
group.id = mygroup
```

--group xx 

消费者组有消费者进出，都会进行 再分区



# 总结

roundrobin 
range


offset: 
    消费者组 




# HW LEO 

high watermark： 所有副本中的最小的LEO
log end offset: 每个副本的最后一个offset

为了保护 消费者


多去少补
抢 




# kafka 高效读写数据

分区，+  索引，文件 + 追加写


1. 顺序写磁盘 预留空间 几千倍
2. 应用pageche 页缓存 系统不崩溃内存
    
    如果生产和消费速度相当，可以不用刷盘了 交换数据
    
    
3.  零拷贝
    
    pagecache
    
    sendfile
    mmap



预期  
单分区： 


## 分区不能减少， 保护消费者



### zookeper在kafka中作用
kafka集群中有一个broker会被选举为controller, 负责管理集群的broker的上下线。
所有topic的分区副本分配和leader选举等工作。 

controller的管理工作都是依赖于zookeeper的




## kafka事务

生产者事务

Transaction id
 





