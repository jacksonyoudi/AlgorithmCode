# -*- coding:utf-8 -*-

# author: youdi
# datetime: 2020/7/15 08:51
# software: IntelliJ IDEA


from pyspark.sql import SparkSession

if __name__ == '__main__':
    spark = SparkSession.builder.getOrCreate()

    df = spark.read.format("jdbc").options(url="jdbc:mysql://localhost:3306?user=root&password=root&useSSL=false",
                                           dbtable="ireadweek.ireadweek").load()
    df.registerTempTable("gdp")

    df2 = spark.sql("select * from gdp limit 1000")
    df2.show()
    # drop table
    spark.catalog.dropTempView("gdp")
    spark.stop()
