# coding: Utf-8

from pyspark import SparkConf, SparkContext
from pyspark.streaming import StreamingContext

if __name__ == '__main__':
    conf = SparkConf().setMaster("local[*]").setAppName("streaming")
    sc = SparkContext(conf)
    streaming = StreamingContext(sc, 1)
