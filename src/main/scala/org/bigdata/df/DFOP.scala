package org.bigdata.df

import org.apache.spark.SparkConf
import org.apache.spark.sql.{DataFrame, SparkSession}

object DFOP {
  def main(args: Array[String]): Unit = {
    val conf: SparkConf = new SparkConf().setMaster("local[*]")
    val spark: SparkSession = SparkSession.builder().config(conf).getOrCreate()
    import spark.implicits._

//    val df: DataFrame = spark.read.parquet("")
//    df.select("age").where("name is not null and age > 100").foreach(println)
//    df.select("name", "age").groupBy("age").count().foreach(println)
//    df.select("age", "name").filter("age > 100").foreach(println)
//
//
//    df.rollup("name", "subject").avg("grade").show()
//    df.groupBy("name").pivot("subject").avg("grade").show()
//    df.cube("name", "subject").avg("grade").show()

  }
}
