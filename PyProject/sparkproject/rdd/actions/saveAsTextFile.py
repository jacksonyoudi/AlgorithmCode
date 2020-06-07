# coding: Utf-8

from pyspark import SparkConf, SparkContext

if __name__ == '__main__':
    conf = SparkConf().setAppName("rdd").setMaster("local[*]")
    sc = SparkContext(conf=conf)
    data = [i for i in range(1, 11)]
    disData = sc.parallelize(data, 10)
    print(disData.saveAsTextFile("/Users/changyouliang/project/javaproject/AlgorithmCode/PyProject/sparkproject/out/asfile"))
