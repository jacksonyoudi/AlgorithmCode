˜### Hbase安装

管理页面 http://localhost:16010/


### shell命令

```shell
help

list
  查看所有表



```

* 创建表

```shell
create table_name column_family



hbase:013:0> desc 'student'
Table student is ENABLED
student
COLUMN FAMILIES DESCRIPTION
{NAME => 'info', BLOOMFILTER => 'ROW', 
IN_MEMORY => 'false', VERSIONS => '1', KEEP_DELETED_CELLS => 'FALSE'
, DATA_BLOCK_ENCODING => 'NONE', COMPRESSION => 'NONE', TTL => 'FOREVER', MIN_VERSIONS => '0', BLOCKCACHE =
> 'true', BLOCKSIZE => '65536',
 REPLICATION_SCOPE => '0'}

```

* 写数据

```shell
  hbase> put 'ns1:t1', 'r1', 'c1', 'value'
  hbase> put 't1', 'r1', 'c1', 'value'
  hbase> put 't1', 'r1', 'c1', 'value', ts1
  hbase> put 't1', 'r1', 'c1', 'value', {ATTRIBUTES=>{'mykey'=>'myvalue'}}
  hbase> put 't1', 'r1', 'c1', 'value', ts1, {ATTRIBUTES=>{'mykey'=>'myvalue'}}
  hbase> put 't1', 'r1', 'c1', 'value', ts1, {VISIBILITY=>'PRIVATE|SECRET'}




put 'student','1001','info:name','changyou'

```
稀疏表


* 读取数据

```shell
get 'student','1001'

scan 'student'

scan 'student',{STARTROW=>'1001', STOPROW=> '1002'}

get  'student','1001','info:sex'
```
按字典顺序的


rowkey  一定要有顺序的

* 统计

```shell
count 'table'
```

* 删除数据
```shell
deleteall 'table','rowkey'

delete 'table','rowkey','family:key'

```

* 清空表数据

```shell
trucncate 'table'
```
hbase操作是 原子性的

* 删除表

```shell
disable 'table'
drop 'table'
```


* 变更表信息

```shell
alter 'table',{NAME=>'info', VERSION=>3}
```

* 获取所有信息

```shell
get  'student','1001',{COLUMN=>'info:name',VERSIONS=>3}



scan 'student',{RAW=>true,VERSIONS=>3}
```


标记
```shell
hbase:045:0> scan 'student',{RAW=>true,VERSIONS=>3}
ROW                         COLUMN+CELL
 1001                       column=info:name, timestamp=2021-04-29T23:08:38.390, type=Delete
 1001                       column=info:name, timestamp=2021-04-29T23:08:38.390, value=jack
 1001                       column=info:name, timestamp=2021-04-29T22:57:04.493, value=youdi
 1001                       column=info:sex, timestamp=2021-04-29T22:53:42.216, value=man
1 row(s)
Took 0.0124 seconds
```




#### 数据结构

hbase数据访问有3种方式

1. 通过rowkey
2. 使用正则 like range
3. scan

rowkey是保存在字节数组 64kb长度, 存储时按rowkey字典序排序的 
位置相关性 (预读)



#### column family
建表的时候，必须要指定


#### CELL
{rowkey, column family:column , version}


#### TimeStamp
版本， ttl

#### namespace

```shell
create_namespace 
list_namespace


create table 'namespace:table','cf:column'
```



### Hbase读取过程
 1. client向zk请求 /hbase/meta-region-server机器，获取数据
 2. 到指定meta数据的region-server请求，查询表 student的region-server
 3. 到student的region-server请求 1001数据


### Hbase写操作

1. 请求zk,获取meta
2. 

```shell
regionServer
store(cf)

WAl (Hlog)  write ahead log 
mem store
HFile

```




### 类加载器

1. 启动类加载器
    c++实现的,加载java的核心类库
2. 扩展类加载器 sun.misc.Launcher$ExtClassLoader
   扩展类加载器加载java扩展类库
3. 应用类加载器
    加载classPath下面的类的
   
双亲委派机制



### hbase索引

索引表

user info:name youdi


动态列
user_index 
    
    rowkey info:youdi
    rowkey info:jackson




### 协处理器




### 高可用




### 预分区

regionserver维护 region的 StartRowKey  EndRowKey

middleKey

split 

数据倾斜


提前将分区设置好，减少合并以及split的操作


1. 手动设定预分区
```shell
create 'tablename','info','partition1',SPLITS => ['1001','2000','3000','4000']

```
2. 生成16进制序列预分区

```shell
create 'table1','info','partition2',{
  NUMREGIONS =>15,SPLITALGO => 'HEXSTRINGSplit'
}

```

3. 按照文件中设置的规则预分区
创建splits.txt文件内容如下
   
```shell
aaaa
bbbb
cccc
dddd
```
然后执行
```shell
create 'table','partion2',SPLITS_FILE => 'splits.txt'

```

4. 使用JAVAAPI创建预分区



#### rowkey设计原则

1. 唯一性原则(类似关系型数据库的主键)
2. 长度原则 (满足需求的，越短越好  60-80字节)
3. 散列原则(让rowkey均匀分配)


```shell
1. hash 随机数 散列值
2. 字符串翻转
3. 字符串拼接



```


### 内存优化
stop the work


```shell

Hstorewe文件大小
hbase.hregion.max.filesize    10G
```


### flush compact split机制



