### 分布式文件系统

* 领域语言

### 模型选择

列名：

family    Qualifier

family： 
    权限控制， 关联数据




### map的集合应该用什么数据结构呢？

使用B+树，需要一种顺序，比较好的做法，url按字典序排序


行内数据：
    先将大表按照列拆分



### 查询和写入

tablet  分布的最小单位

chunk的抽象



table -> tablet -> chunk -> Block



### 分布式文件的管理




### socket文件存在哪里
/proc/net/tcp
lsof -i -a -p 2377



### 日志文件的冗余如何处理？

碎片进行整理


### hash表

B+数  范围
hash表： 只能点查 


### master节点如果宕机了， 影响有多大， 如何恢复？


