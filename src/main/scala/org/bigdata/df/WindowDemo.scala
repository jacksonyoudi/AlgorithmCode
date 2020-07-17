package org.bigdata.df

import org.apache.spark.SparkConf
import org.apache.spark.sql.expressions.{Window, WindowSpec}
import org.apache.spark.sql.{DataFrame, SparkSession}
import org.apache.spark.sql.functions._

object WindowDemo {
  def main(args: Array[String]): Unit = {
    val conf: SparkConf = new SparkConf().setMaster("local[*]")
    val spark: SparkSession = SparkSession.builder().config(conf).getOrCreate()
    val windowSpec: WindowSpec = Window
      .partitionBy("name", "subject")
      .orderBy(desc("grade"))

    val df: DataFrame = spark.read.parquet("")

    df.select(
      col("name"),
      col("subject"),
      col("year"),
      col("grade"),
      row_number().over(windowSpec).as("xixi")
    ).show()

    val slide: WindowSpec = Window
      .partitionBy("key")
      .orderBy("num")
      .rangeBetween(Window.currentRow + 2, Window.currentRow + 20)

    df.select(
      col("key"),
      sum("num").over(slide)
    )
      .sort("key")
      .show()

  }
}
