### hbase在HDFS中文件布局

1. .hbase-snapshot 
快照的元数据
   
2. .tmp 

```shell
临时文件目录，主要用于HBase表的创建和删除操作。
表创建的时候首先会在tmp目录下执行，执行成功后再将tmp目录下的表信息移动到实际表目录下。
表删除操作会将表目录移动到tmp目录下，一定时间过后再将tmp目录下的文件真正删除。
```

3. MasterProcWALs

```shell
存储Master Procedure过程中的WAL文件。
```

4. WALs：存储集群中所有RegionServer的HLog日志文件。

5. archive：文件归档目录。这个目录主要会在以下几个场景下使用。
   
   ○所有对HFile文件的删除操作都会将待删除文件临时放在该目录。
   
    ○进行Snapshot或者升级时使用到的归档目录。
   
    ○Compaction删除HFile的时候，也会把旧的HFile移动到这里。
    
    删除，snapshot, compact

6. coorupt

    存储损坏的hlog文件或者hfile 

7. data 存储集群中所有region的hfile数据， hfile文件在data目录下的完整路径


```shell
 /hbase/data/default/usertable/fal3333t/family/9908758969576
```
default 是命名空间
usertable是表名
fal3333t region名称
family 列簇名
9908758969576 hfile名字



data目录下还有

.tabledesc 表描述文件  schema
.tmp  记录flush 和 compact 过程的中间结果
.regioninfo 
recovered.edits   wal 回放数据
hbase.id 
hbase.version 
oldwals wal归档目录




### Hbase客户端实现

hbase客户端配置文件
hbase-site.xml
core-site.xml
hdfs-site.xml


Connection是HBase客户端进行一切操作的基础，它维持了客户端到整个HBase集群的连接，例如一个HBase集群中有2个Master、5个RegionServer，那么一般来说，这个Connection会维持一个到Active Master的TCP连接和5个到RegionServer的TCP连接。


Connection还缓存了访问的Meta信息，这样后续的大部分请求都可以通过缓存的Meta信息定位到对应的RegionServer。



Table是一个非常轻量级的对象，它实现了用户访问表的所有API操作，例如Put、Get、Delete、Scan等。本质上，它所使用的连接资源、配置信息、线程池、Meta缓存等，都来自步骤2创建的Connection对象。
因此，由同一个Connection创建的多个Table，都会共享连接、配置信息、线程池、Meta缓存这些资源。


hbase:meta表
专门用来存放整个集群所有region信息。 
namespace

整个表只有一个名为info的 family
只有一个region, 为了确保meta表多次操作的原子性。 因为hbase本质上只支持region级别的事务。
注意表结构中用到了MultiRowMutationEndpoint这个coprocessor

![bXn9XU](https://raw.githubusercontent.com/jacksonyoudi/images/main/uPic/bXn9XU.png)
总体来说，hbase:meta的一个rowkey就对应一个Region，rowkey主要由TableName（业务表名）、StartRow（业务表Region区间的起始rowkey）、Timestamp（Region创建的时间戳）、EncodedName（上面3个字段的MD5 Hex值）4个字段拼接而成。每一行数据又分为4列，分别是info:regioninfo、info:seqnumDuringOpen、info:server、info:serverstartcode。



metacache 缓存到客户端

/hbase/meta-region-server

1. 没有缓存， 请求
2. 有缓存， 数据已经过期 请求
3. 有缓存， 正常使用



scan： 
loadcache


lsm


