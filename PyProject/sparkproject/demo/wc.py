# coding: utf-8

from pyspark import SparkConf, SparkContext

if __name__ == '__main__':
    conf = SparkConf().setAppName("hello").setMaster("local")
    sc = SparkContext().getOrCreate(conf)
    logData = sc.textFile("/Users/youdi/Project/javaProject/AlgorithmCode/GoProject/main.go").cache()
    nums = logData.filter(lambda x: 'a' in x).count()
    print(nums)
