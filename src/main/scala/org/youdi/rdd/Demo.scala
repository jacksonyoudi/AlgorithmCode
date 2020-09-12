package org.youdi.rdd

import org.apache.spark.SparkContext
import org.apache.spark.rdd.RDD


object Demo {
  def main(args: Array[String]): Unit = {
    val sc: SparkContext = SparkContext.getOrCreate()
    val tRdd: RDD[String] = sc.textFile("hello")


    // 1TB数据排序
//    val sRdd: RDD[Array[String]] = tRdd.map(_.split(" "))
//    sRdd.sortBy(
//    )


  }
}
