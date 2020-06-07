# coding: Utf-8

from pyspark import SparkConf, SparkContext
from operator import add

if __name__ == '__main__':
    conf = SparkConf().setAppName("rdd").setMaster("local[*]")
    sc = SparkContext(conf=conf)
    data = [i for i in range(1, 11)]
    disData = sc.parallelize(data, 10)

    print(disData.reduce(add))

    print(sc.parallelize((2 for _ in range(10))).map(lambda x: 1).cache().reduce(add))
