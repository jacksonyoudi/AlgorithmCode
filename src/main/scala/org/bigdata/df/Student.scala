package org.bigdata.df


import org.apache.spark.sql.{DataFrame, Dataset, SparkSession}

object Student {
  def main(args: Array[String]): Unit = {
    val spark: SparkSession = SparkSession
      .builder()
      .appName("Student")
      .master("local[*]")
      .getOrCreate()

    val df: DataFrame = spark.read.json("/Users/youdi/Project/javaProject/AlgorithmCode/src/main/scala/org/bigdata/df/stu.json")
    import spark.implicits._


    val ds: Dataset[Stu] = df.map(
      a => Stu(
        a.getAs[String](1),
        a.getAs[Long](0),
        a.getAs[String](2)))

    //    val ds1: Dataset[Stu] = ds.as[Stu]
    ds.show()
  }
}
