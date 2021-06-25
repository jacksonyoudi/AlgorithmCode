# kafka




## 优点

* 解耦
* 缓冲 削峰
* 可恢复
* 灵活性 & 峰值处理能力
* 异步通信



## 消息队列的两种模式

1. 点对点

消费者主动拉取数据，消息收到后消息清除

2. 发布订阅模式
    消费者消费数据后，不会清除
   
pushish
subscribe



## kafka架构

1. producer 
2. consumer
3. broker
4. zookeper


topic: 相同消息发送一起， 
replicate，副本，不能保存在一个
partition: 分区 提高写的并发