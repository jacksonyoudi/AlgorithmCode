# coding: Utf-8

from pyspark import SparkConf, SparkContext

if __name__ == '__main__':
    conf = SparkConf().setAppName("rdd").setMaster("local[*]")
    sc = SparkContext(conf=conf)
    sc.setCheckpointDir("/tmp/cp")
    data = [1, 2, 3, 4, 5]
    disData = sc.parallelize(data, 10)
    # map
    disData = disData.map(lambda x: 2 * x)
    disData.cache()
    disData.checkpoint()

