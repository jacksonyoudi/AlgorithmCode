# coding: Utf-8

from pyspark import SparkConf, SparkContext

"""
mapPartitions
"""


def foo(index, iterator):
    yield index, list(iterator)


if __name__ == '__main__':
    conf = SparkConf().setAppName("rdd").setMaster("local[*]")
    sc = SparkContext(conf=conf)
    data = [i for i in range(1, 11)]

    disData = sc.parallelize(data, 2)

    print(disData.mapPartitionsWithIndex(foo).collect())
