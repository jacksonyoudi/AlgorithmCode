# coding: Utf-8

from pyspark import SparkConf, SparkContext

"""
filter
"""

if __name__ == '__main__':
    conf = SparkConf().setAppName("rdd").setMaster("local[*]")
    sc = SparkContext(conf=conf)
    data = [i for i in range(1, 11)]
    disData = sc.parallelize(data, 10)

    print(disData.filter(lambda x: x % 2 == 0).collect())
