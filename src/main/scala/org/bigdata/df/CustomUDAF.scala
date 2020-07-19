package org.bigdata.df

import org.apache.spark.sql.expressions._
import org.apache.spark.sql.types._
import org.apache.spark.sql.{DataFrame, Row, SparkSession}
import org.apache.spark.sql.functions._


object CustomUDAF extends UserDefinedAggregateFunction {
  override def inputSchema: StructType = {
    StructType(Array(StructField("input", IntegerType, true)))
  }

  override def bufferSchema: StructType = {
    StructType(Array(StructField("max", IntegerType, true)))
  }

  override def dataType: DataType = {
    IntegerType
  }

  override def deterministic: Boolean = {
    true
  }

  override def initialize(buffer: MutableAggregationBuffer): Unit = {
    buffer(0) = 0
  }

  override def update(buffer: MutableAggregationBuffer, input: Row): Unit = {
    val i: Int = input.getAs[Int](0)
    val i1: Int = buffer.getAs[Int](0)
    if (i > i1) {
      buffer(0) = i
    }
  }

  override def merge(buffer1: MutableAggregationBuffer, buffer2: Row): Unit = {
    if (buffer1.getAs[Int](0) > buffer2.getAs[Int](0)) {
      buffer1(0) = buffer2.getAs[Int](0)
    }
  }

  override def evaluate(buffer: Row): Any = {
    buffer.getAs[Int](0)
  }


  def main(args: Array[String]): Unit = {
    val spark: SparkSession = SparkSession
      .builder()
      .master("local[*]")
      .appName("Test")
      .getOrCreate()

    import spark.implicits._
    val df: DataFrame = spark.read.json("")

    df.select(CustomUDAF(col("grade"))).show()
  }
}


