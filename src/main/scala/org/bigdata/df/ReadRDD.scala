package org.bigdata.df

import org.apache.spark.SparkConf
import org.apache.spark.rdd.RDD
import org.apache.spark.sql.types.{StringType, StructField, StructType}
import org.apache.spark.sql.{DataFrame, Row, SparkSession}

object ReadRDD {
  def main(args: Array[String]): Unit = {
    val conf: SparkConf = new SparkConf().setMaster("local[*]")
    val spark: SparkSession = SparkSession.builder().config(conf).getOrCreate()

    val s: String = "id f1 f2 f3 f4"
    import spark.implicits._
    val fields: Array[StructField] = s.split(" ").map(StructField(_, StringType, true))
    val schema: StructType = StructType(fields)

    val rowRDD: RDD[Row] = spark.sparkContext.textFile("xxx").map(_.split(",")).map(
      data => Row(data(0), data(1), data(2), data(3), data(4))
    )

    val frame: DataFrame = spark.createDataFrame(rowRDD, schema)
    frame.show()
  }
}
