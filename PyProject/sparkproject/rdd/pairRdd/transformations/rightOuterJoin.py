# coding: Utf-8

from pyspark import SparkConf, SparkContext

if __name__ == '__main__':
    conf = SparkConf().setAppName("rdd").setMaster("local[*]")
    sc = SparkContext(conf=conf)
    x = sc.parallelize([("a", 1), ("c", 3)])
    y = sc.parallelize([("a", 2), ("b", 2)])
    # x.leftOuterJoin(y).foreach(print)
    print(x.rightOuterJoin(y).collect())
    print(y.fullOuterJoin(x).collect())
