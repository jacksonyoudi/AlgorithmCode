# coding: utf-8


from pyspark.sql import SparkSession

if __name__ == '__main__':
    # conf = SparkConf().setAppName("sql").setMaster("local[*]")
    spark = SparkSession.builder.getOrCreate()
    df = spark.read.csv("./gdp.csv")

    df.registerTempTable("gdp")

    df2 = spark.sql("select * from gdp limit 1000")
    # filter
    # df3 = df2.filter("_c2='1969'")
    # df3 = df2.select("_c0").distinct()

    df3 = df2.select("_c0").groupBy("_c3").count()
    df3.show()
    # drop table
    spark.catalog.dropTempView("gdp")
    spark.stop()
