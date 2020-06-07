# coding: Utf-8

from pyspark import SparkConf, SparkContext


def extract_call_signals(line):
    global blankLines
    if line % 2 == 0:
        blankLines += 1
    return line


if __name__ == '__main__':
    conf = SparkConf().setAppName("rdd").setMaster("local[*]")
    sc = SparkContext(conf=conf)
    data = [i for i in range(1, 11)]
    bc = sc.broadcast([1, 2, 3, 4, 5])
    print(sc.parallelize(data).flatMap(lambda x: bc.value).collect())
