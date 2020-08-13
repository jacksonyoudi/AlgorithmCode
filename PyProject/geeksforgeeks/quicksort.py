# -*- coding:utf-8 -*-
# author: youdi

from typing import List


def partition(arr: List[int], low: int, high: int) -> int:
    i = low - 1
    pivot = arr[high]
    for j in range(low, high):
        if arr[j] <= pivot:
            i += 1
            arr[j], arr[i] = arr[i], arr[j]
    arr[i + 1], arr[high] = arr[high], arr[i + 1]
    return i + 1


def quicksort(arr: List[int], low: int, high: int):
    if low < high:
        pi = partition(arr, low, high)
        quicksort(arr, low, pi - 1)
        quicksort(arr, pi + 1, high)


if __name__ == '__main__':
    A = [64, 25, 12, 22, 11]
    quicksort(A, 0, len(A) - 1)
    print(A)
