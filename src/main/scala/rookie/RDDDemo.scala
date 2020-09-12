package rookie

import org.apache.spark.rdd.RDD
import org.apache.spark.{SparkConf, SparkContext}
import org.apache.spark.sql.{DataFrame, SparkSession}

/**
 * * A Resilient Distributed Dataset (RDD), the basic abstraction in Spark. Represents an immutable,
 * * partitioned collection of elements that can be operated on in parallel.
 *
 * RDD是以下三个字母的首字母缩写(Resilient Distributed Dataset) 它表示弹性分布式数据集，
 * 是spark最基本的数据抽象，代表一个不可变，可分区，里面元素可以被进行并行操作的集合。
 *
 * DataSet
 *    数据集，可以理解是一个集合，集合中存储了很多数据
 * Distributed
 *    数据是进行了分布式存储，为了方便后期进行分布式计算
 * Resilient
 *    弹性，rdd的数据可以保存在内存中或磁盘中
 *
 *
 *
 * * Internally, each RDD is characterized by five main properties:
 * *
 * *  - A list of partitions
 * 一个RDD有很多分区，一组分区列表
 *
 * *  - A function for computing each split
 * *  - A list of dependencies on other RDDs
 * *  - Optionally, a Partitioner for key-value RDDs (e.g. to say that the RDD is hash-partitioned)
 * *  - Optionally, a list of preferred locations to compute each split on (e.g. block locations for
 * *    an HDFS file)
 *
 *
 *
 * spark优点
 *  分布式
 *  基于内存，也可以使用磁盘
 *  迭代式计算  shuffle
 *
 *
 *
 *  spark-client
 *    RDD grapha
 *    Scheduler
 *    BlockManager
 *    Shuffle tracker
 *
 *  Spark worker
 *    Task threads
 *    Block manager
 */



object RDDDemo {
  def main(args: Array[String]): Unit = {
    val conf: SparkConf = new SparkConf()
    //    val sc: SparkSession = SparkSession.builder().config(conf).getOrCreate()
    val sc: SparkContext = SparkContext.getOrCreate(conf)
    val data: RDD[Int] = sc.makeRDD(List[Int](1, 2, 3, 4))
  }
}
