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
    blankLines = sc.accumulator(0)
    sc.parallelize(data, 4).map(extract_call_signals).collect()
    print(blankLines)
