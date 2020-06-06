# coding: Utf-8

from pyspark import SparkConf, SparkContext

"""
parallelize
RDD支持的两种类型的操作： 转换 transformastions actions

transformastions lazy

 
"""

if __name__ == '__main__':
    conf = SparkConf().setAppName("rdd").setMaster("local[*]")
    sc = SparkContext(conf=conf)
    data = [1, 2, 3, 4, 5]
    disData = sc.parallelize(data, 10)

    # map
    disData = disData.map(lambda x: 2 * x)
    print(disData.collect())
