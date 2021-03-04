package org.bigdata.es

import org.apache.spark.{SparkConf, SparkContext}
//import org.elasticsearch.spark._

object D01 {
  def main(args: Array[String]): Unit = {
    val conf: SparkConf = new SparkConf().setAppName("d01").setMaster("local[*]")
    conf.set("es.index.auto.create", "true")

    val sc: SparkContext = new SparkContext(conf)

    // map方式
    val numbers = Map("one" -> 1, "two" -> 2, "three" -> 3)
    val airports = Map("arrival" -> "Otopeni", "SFO" -> "San Fran")
//    sc.makeRDD(Seq(numbers, airports)).saveToEs("spark/docs")

  }
}
