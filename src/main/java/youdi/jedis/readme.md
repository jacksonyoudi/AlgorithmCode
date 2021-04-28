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



### redis使用



### 行式存储数据库
### 列式存储数据库



### redis应用



### 缓存穿透
1. 应用服务器压力变大
2. redis命中率降低
3. 一直查询数据库


原因： 

1. redis查询不到， 数据库中查询
2. 出现很多非常多url的请求




解决：

1. 对空值做缓存
2. 设置白名单 使用bitmap
3. 采用布隆过滤器
4. 实时监控  和运维一起处理


### 缓存击穿

1. 数据库访问压力瞬时增加
2. redis里面没有出现大量key过期
3. redis正常运行



原因： 
redis中某个key过期了， 大量访问使用这个key。



解决方案：
1. 预先设置热门数据
2. 实时调整
3. 使用锁




### 缓存雪崩

1. 数据库压力变大
2. 服务器崩溃


在极少时间段内， 查询大量key的集中过期情况


解决方案：

1. 构建多级缓存架构
nginx缓存 + redis缓存  + 其他缓存(ehcache)
   
2. 使用锁或队列

3. 设置过期标志更新缓存

4. 缓存失效时间分开
    加上随机时间
   





### 分布式锁



#### 实现方案

1. 基于数据实现分布式锁
2. 基于zookeeper
3. 基于redis实现





#### redis分布式锁
```shell
setnx key value
del key
```

给锁设置过期时间

```shell
expire key ttl
```

1.使用setnx上锁，通过del释放锁
2. 设置过期时间，过期时间就会释放
3. 原子操作 上锁的同时设置过期时间



```shell
set key value nx ex 12
```



问题：
释放他人持有的锁

使用uuid
类似乐观锁

```shell
set lock uuid
```

第一步 uuid表示不同的操作
set lock uuid nx ex 10 

第二步 释放锁的时候，首先判断当前uuid和要释放的uuid是否一样






### 临界删除锁


A：

```shell
1. 上锁
2. 具体操作
  1) 比较uuid
  2) 删除操作的时候，正要删除，还没删除，锁到了过期，自动释放了
3. 删除操作
```

B:

```shell
1. b上锁
2. 具体操作

3. a释放b的锁
```

没有进行原子性操作

lua脚本




### 新功能
ACL











