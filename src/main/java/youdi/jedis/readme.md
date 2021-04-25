### redis事务

组队  mutli  队列

执行  exec

discard: 放弃组队


mutli失败了， 所有的命令都失败

exec执行，其中一个失败， 不影响其他的


### watch
监控这一个key， 如果key被改动， 事务就被打断了。
乐观锁 (版本好)



### Redis事务三特性
1. 单独的隔离性
    事务中的所有命令都会序列化，按顺序执行， 事务在执行过程中， 不会被其他客户端发送来的命令所打断
   
2. 没有隔离级别的概念
    队列中的命令没有提交之前都不会实际执行的，因为事务提交前任务指令都不会实际执行
   
3. 不保证原子性
    事务中如果有一个命令执行失败，其后的命令仍然会被执行，没有回滚
   



### 秒杀系统

1. 超时
   
    使用连接池
   
2. 超卖
   
    乐观锁 + 事务
   
3. 乐观锁造成库存遗留问题
   
    lua脚本语言



### 持久化

#### RDB



在指定的时间间隔内将内存中数据集快照写入磁盘中

snapshot 

dump.rdb 文件



持久化的规则
```shell
save 3600 1
save 300  100
save 60 10000
```


##### save VS  bgsave

save只管备份，其他的不管， 进行阻塞，其他的都不能操作
bgsave: redis会在后台异步进行快照操作，快照同时还是可以响应客户端请求
可以通过lastsave


flushall

stop-writes-on-bgsave-error yes

checksum yes




缺点： 最后一次持久化后的数据可能丢失(还未触发)

wof写时复制


#### redis备份


### AOF
append only file

日志追加的方式， 只记录写的方式

以AOF为准

redis-check-aof  --fix xx.aof

同步频率设置
appendfsync 

always
everysec
no 操作系统负责




bgrewrite
    64M

 ```shell
>info replication

在slave上执行
slaveof master:port
```



slaveof no one




### 哨兵模式

```shell
slaveof host port
```

sentinel.conf

```shell
sentinel monitor mymaster host port cnt
cnt：多少个哨兵同意数据迁移


26379


```








