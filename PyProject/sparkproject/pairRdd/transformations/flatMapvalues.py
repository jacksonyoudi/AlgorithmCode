# coding: Utf-8

from pyspark import SparkConf, SparkContext

if __name__ == '__main__':
    conf = SparkConf().setAppName("rdd").setMaster("local[*]")
    sc = SparkContext(conf=conf)
    data = [("a", ["a", "b", "c"]), ("b", ["grapes"])]
    disData = sc.parallelize(data)

    print(disData.flatMapValues(lambda x: x).collect())
