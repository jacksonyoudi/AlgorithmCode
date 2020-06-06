# coding: Utf-8

from pyspark import SparkConf, SparkContext


"""
glom() 一个分区内转成list

"""
if __name__ == '__main__':
    conf = SparkConf().setAppName("rdd").setMaster("local[*]")
    sc = SparkContext(conf=conf)
    data = [i for i in range(1, 11)]

    disData = sc.parallelize(data).map(lambda x: (x, x))

    print(disData.partitionBy(2).glom().collect())

    print(disData.partitionBy(3).glom().collect())
