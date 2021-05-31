### zookeeper应用场景

1. 配置管理
2. DNS服务
3. 组成员管理
4. 各种分布式锁

zk用于存储和协同相关关键数据，不适合用于大数据量存储。




### 数据模型
data tree 


### Znode分类
1. persistend 持久性的
2. enhemeral  临时性 

3. 持久顺序的znode 
4. 临时顺序znode enhemeral_sequential



### zk操作

```shell
 create /youdi
```

### 分布式锁

```shell

 create -e /lock

stat -w /lock
```

### 协同服务

集群中只能有一个master, master监听多个 worker节点


```shell
create -e /master "master1:2223"


stat -w /master


create -e /master "master2:2223"
```


```shell
create -e /workers/w1 "w1:2224"
```


### zk架构

standalone
quorum


Sesion


### Quorum

leader：可读可写
follower：只处理读请求，将写请求转发leader来处理。



### 数据一致性

全局可线性化linearizable写入。 leader写入是有顺序的
客户端FIFO顺序

