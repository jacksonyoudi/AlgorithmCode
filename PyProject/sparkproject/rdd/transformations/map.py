# coding: Utf-8

from pyspark import SparkConf, SparkContext

"""
map 
"""


def add(x):
    return x + 10


if __name__ == '__main__':
    conf = SparkConf().setAppName("map").setMaster("local[*]")
    sc = SparkContext(conf=conf)
    data = [1, 2, 3, 4, 5]
    disData = sc.parallelize(data, 10)

    # map
    disData = disData.map(add)
    print(disData.collect())
