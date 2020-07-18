package org.bigdata.df

import org.apache.spark.sql.expressions.UserDefinedFunction
import org.apache.spark.sql.{DataFrame, SparkSession}
import org.apache.spark.sql.functions._

// 7358437003
object CustomUDF {
  def main(args: Array[String]): Unit = {
    val spark: SparkSession = SparkSession.builder
      .master("local[*]")
      .appName("udf")
      .getOrCreate()

    import spark.implicits._

    val df: DataFrame = spark.read.json("")

    def level(grade: Int): String = {
      if (grade > 85) "A"
      else if (grade <= 85 & grade > 75)
        "B"
      else if (grade <= 75 & grade >= 60)
        "C"
      else if (grade < 60)
        "D"
      else
        "Error"
    }

//    spark.udf.register("level", level _)
    val myudf: UserDefinedFunction = udf(level _)
    df.select(col("name"), myudf(col("grade"))).show()

  }

}
