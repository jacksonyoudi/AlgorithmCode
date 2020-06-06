# coding: Utf-8

from pyspark import SparkConf, SparkContext

"""
flatMap 
"""

if __name__ == '__main__':
    conf = SparkConf().setAppName("rdd").setMaster("local[*]")
    sc = SparkContext(conf=conf)
    data = [8, 9, 10]
    disData = sc.parallelize(data)

    print(disData.map(lambda x: range(1, x)).collect())
    print(disData.flatMap(lambda x: range(1, x)).collect())
