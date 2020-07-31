# -*- coding:utf-8 -*-
# author: youdi

from typing import List


def partition(arr: List[int], low: int, high: int) -> int:
    i = low - 1  # 最小元素索引
    pivot = arr[high]  # 基准值
    for j in range(low, high):
        # 当前元素小于或等于 pivot
        if arr[j] <= pivot:
            i += 1
            arr[i], arr[j] = arr[j], arr[i]
    arr[i + 1], arr[high] = arr[high], arr[i + 1]
    return i + 1


def quick_sort(arr: List[int], low: int, high: int):
    if low < high:
        pi = partition(arr, low, high)
        quick_sort(arr, low, pi - 1)
        quick_sort(arr, pi + 1, high)


if __name__ == '__main__':
    arr = [10, 7, 8, 9, 1, 5]
    quick_sort(arr, 0, len(arr) - 1)
    print(arr, 0, len(arr) - 1)
