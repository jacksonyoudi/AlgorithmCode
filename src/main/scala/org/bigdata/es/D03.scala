package org.bigdata.es

import org.apache.spark.{SparkConf, SparkContext}
import org.elasticsearch.spark._


object D03 {
  def main(args: Array[String]): Unit = {
    val conf: SparkConf = new SparkConf().setAppName("d01").setMaster("local[*]")
    conf.set("es.index.auto.create", "true")

    val sc: SparkContext = new SparkContext(conf)

    val json1 = """{"reason" : "business", "airport" : "SFO"}"""
    val json2 = """{"participants" : 5, "airport" : "OTP"}"""
    sc.makeRDD(Seq(json1, json2)).saveToEs("spark/json-trips")

  }
}
