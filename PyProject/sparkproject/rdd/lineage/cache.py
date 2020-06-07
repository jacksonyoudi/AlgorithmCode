# coding: Utf-8

from pyspark import SparkConf, SparkContext

"""
StorageLevel.DISK_ONLY = StorageLevel(True, False, False, False)
StorageLevel.DISK_ONLY_2 = StorageLevel(True, False, False, False, 2)
StorageLevel.MEMORY_ONLY = StorageLevel(False, True, False, False)
StorageLevel.MEMORY_ONLY_2 = StorageLevel(False, True, False, False, 2)
StorageLevel.MEMORY_AND_DISK = StorageLevel(True, True, False, False)
StorageLevel.MEMORY_AND_DISK_2 = StorageLevel(True, True, False, False, 2)
StorageLevel.OFF_HEAP = StorageLevel(True, True, True, False, 1)
"""

if __name__ == '__main__':
    conf = SparkConf().setAppName("rdd").setMaster("local[*]")
    sc = SparkContext(conf=conf)
    data = [1, 2, 3, 4, 5]
    disData = sc.parallelize(data, 10)

    # map
    disData = disData.map(lambda x: 2 * x)
    print(disData.persist().is_cached)
    # Memory Serialized 1x Replicated
    print(disData.getStorageLevel())
    disData.unpersist()
