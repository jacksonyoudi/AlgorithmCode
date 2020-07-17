package org.bigdata.df

import org.apache.spark.sql.{DataFrame, SparkSession}
import org.apache.spark.SparkConf

object DF {
  def main(args: Array[String]): Unit = {
    val conf: SparkConf = new SparkConf().setAppName("sql").setMaster("local[*]")
    val spark: SparkSession = SparkSession.builder().appName("sql").config(conf).getOrCreate()

    // read files
    val df: DataFrame = spark.read.csv("***")
    val df1: DataFrame = spark.read.parquet("***")
    val df2: DataFrame = spark.read.orc("***")
    val df3: DataFrame = spark.read.json("***")
    val df4: DataFrame = spark.read.text("***")


    val df5: DataFrame = spark.read
      .format("jdbc")
      .option("url", "")
      .option("dbtable", "")
      .option("user", "")
      .option("password", "")
      .load()


  }
}
