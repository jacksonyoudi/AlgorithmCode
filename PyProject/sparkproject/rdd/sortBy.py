# coding: Utf-8

from pyspark import SparkConf, SparkContext

"""
sortBy
"""

if __name__ == '__main__':
    conf = SparkConf().setAppName("rdd").setMaster("local[*]")
    sc = SparkContext(conf=conf)
    data = [(i, j) for i in range(1, 11) for j in "abcdefghig".split()]

    disData = sc.parallelize(data, 10)

    print(disData.sortBy(lambda x: x[1], True).collect())
