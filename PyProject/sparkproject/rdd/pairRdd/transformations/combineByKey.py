# coding: Utf-8

from pyspark import SparkConf, SparkContext

if __name__ == '__main__':
    conf = SparkConf().setAppName("rdd").setMaster("local[*]")
    sc = SparkContext(conf=conf)

    a = ["a", "a", "c", "d", "d", "c", "e"]
    b = [1, 2, 3, 4, 1, 3, 7]
    data = list(zip(a, b))
    disData = sc.parallelize(data)
    d = disData.combineByKey(lambda x: [x], lambda x, y: x.append(y), lambda x, y: x + y)

    print(d.collect())
