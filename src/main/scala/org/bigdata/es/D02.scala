package org.bigdata.es

import org.apache.spark.rdd.RDD
import org.apache.spark.{SparkConf, SparkContext}
//import org.elasticsearch.spark.rdd.EsSpark

object D02 {
  def main(args: Array[String]): Unit = {
    val conf: SparkConf = new SparkConf().setAppName("d01").setMaster("local[*]")
    conf.set("es.index.auto.create", "true")

    val sc: SparkContext = new SparkContext(conf)
    val upcomingTrip: Trip = Trip("OTP", "SFO")
    val lastWeekTrip: Trip = Trip("MUC", "OTP")

    val rdd: RDD[Trip] = sc.makeRDD(Seq(upcomingTrip, lastWeekTrip))
//    EsSpark.saveToEs(rdd, "spark/docs", Map("es.mapping.id" -> "id"))
  }
}


// define a case class
case class Trip(departure: String, arrival: String)