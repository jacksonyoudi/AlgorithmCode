### 初识spark


hadoop MRV1
部分： 
1. 运行时环境jobtracker 和 tasktracker
2. 编程模型mapreduce 
3. 数据处理引擎 maptask 和 reduce task 

不足：
1. 可扩展性差 
    jobtracker 负责资源和任务调度
   
2. 可用性差
3. 资源利用率低
4. 不通用




二级调度系统
Resource manager
Nodemanager  
container 


ApplicationManager


spark特点：

1. 减少磁盘IO 
2. 增加并行度
3. 避免重新计算
4. 可选的shuffle排序
5. 灵活的内存管理
6. 检查点支持
7. 易于使用
8. 支持交互式


### 2.2 spark基础知识

#### RDD 
弹性分布式数据集：

```scala
  @DeveloperApi
  def compute(split: Partition, context: TaskContext): Iterator[T]

  /**
   * Implemented by subclasses to return the set of partitions in this RDD. This method will only
   * be called once, so it is safe to implement a time-consuming computation in it.
   *
   * The partitions in this array must satisfy the following property:
   *   `rdd.partitions.zipWithIndex.forall { case (partition, index) => partition.index == index }`
   */
  protected def getPartitions: Array[Partition]

  /**
   * Implemented by subclasses to return how this RDD depends on parent RDDs. This method will only
   * be called once, so it is safe to implement a time-consuming computation in it.
   */
  protected def getDependencies: Seq[Dependency[_]] = deps

  /**
   * Optionally overridden by subclasses to specify placement preferences.
   */
  protected def getPreferredLocations(split: Partition): Seq[String] = Nil

  /** Optionally overridden by subclasses to specify how they are partitioned. */
  @transient val partitioner: Option[Partitioner] = None

  // =======================================================================
  // Methods and fields available on all RDDs
  // =======================================================================

  /** The SparkContext that created this RDD. */
  def sparkContext: SparkContext = sc



```







弹性： 失败重试， 血缘
分布式： 数据都是并行，分区， hadooppartitionor
数据集： 集合，  抽象， scala， 对比于集合 
1. 分区
2. 计算函数
3. 依赖关系
4. sc contenxt
5. 优先位置

action => dagschedular,
partitior: 数据分区
narrowDependency： 窄依赖： onetoOneDepency和rangeDepen
shuffleDependency: 宽依赖， 取决partitioner的函数
Job: 用户提交的作业。当rdd及其dag被提交给dagscheduler调度后，dagscheduler会将所有的rdd中转换及动作视为一个job。
stage： job执行阶段 shufflestage , resultStage,
task: 具体执行任务， 一个job在每个stage内都会按照rdd的partition数量，创建多个task。 task也有shuffletask和resulttask
shufle： 



#### spark模块设计

1. spark core
2. spark sql
3. spark streaming
4. Granphx
5. Mlib


### spark核心功能
1. 基础设施
   sparkconf
   rpc netty 
   ListenerBus 监听器模式异步调用的实现
   度量系统 source sinker
   
2. sparkcontext
    网路通信
    分布式部署
   消息通信
   存储体系
   计算引擎
   度量系统 
   文件服务
   web ui 
   
3. sparkEnv
 spark执行运行环境sparkenv是spark中task运行所必须的组件，sparkenv内部封装了
   rpcenv， 序列化管理器， 广播管理器， 存储体系， 度量系统metricssystem 
   map任务输出跟踪器， 输出提交协调器 outcommitcoordinator
   
4. 存储体系
   Spark的内存空间还提供了Tungsten的实现，直接操作操作系统的内存
   
5. 调度系统
  dagscheduler ，taskscheduler 内置在sparkcontext中，
   dagscheduler负责创建job，将dag中的rdd划分到不同的stage， 给stage创建对应的task
   批量提交task等功能。 
   taskscheduler负责按照fifo或者fair等调度算法对批量task进行调度，为task分配资源。将
   task发送到就按管理器的当前应用的executor上， 由executor负责执行等工作。 
   spark增加了sparksession和dataframe等api， 底层依然依赖于sparkcontext。
   
6. 计算引擎
    内存管理器 memorymanager
    Tungsten
    任务内存管理器 taskMemoryManager
    task
    外部排序器
    shuffle管理器
    

### 2. spark sql 
首先使用sql语句解析器 sqlparser 将sql 转换成语法树 tree， 并且使用规则执行器 ruleExecutor将一系列rule
应用到语法树， 最终生成物理执行计划并执行的过程。

规则包括： 语法分析器 analyer, 和优化器  optimiser 



#### spark编程模型
1. 编写 sparkcontext 应用
2. 提交 应用  rpcenv =>  rm -> am(dirver) -> rm -> excutor  sparkcontext 用blockmanager和broadcastmanager将任务的hadoop配置进行广播。 
3. 

#### RDD计算模型
RDD对各种数据计算模型的统一抽象。
rdd迭代计算。  rdd转换，就是包装，



#### spark基本架构

cluseter manager  Resource manager
worker   node manaager
excutor  
driver: 是 am , 和rm和exude 通信的媒介 


application： 
