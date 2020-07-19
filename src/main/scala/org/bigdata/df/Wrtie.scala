package org.bigdata.df

import org.apache.spark.sql.{DataFrame, SaveMode, SparkSession}

object Wrtie {
  def main(args: Array[String]): Unit = {

    val spark: SparkSession = SparkSession
      .builder()
      .appName("wd")
      .master("local[*]")
      .getOrCreate()

    val df: DataFrame = spark.read.json("")
    df.select("name").write.json("")

    df.select("name").filter("").write.parquet("")

    df.select("name").filter("").write.csv("")
    df.select("name").filter("").write
      .format("csv")
      .mode(SaveMode.Overwrite)
      .option("", "")
      .save()
  }
}
