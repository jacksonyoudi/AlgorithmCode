# coding: Utf-8

from pyspark import SparkConf, SparkContext

if __name__ == '__main__':
    conf = SparkConf().setAppName("rdd").setMaster("spark://127.0.0.1:17077")
    sc = SparkContext(conf=conf)
    data = [1, 2, 3, 4, 5]
    disData = sc.parallelize(data, 10)
    print(disData.collect())
