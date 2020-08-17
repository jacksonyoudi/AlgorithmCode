from typing import List


def partition(arr: List[int], low: int, high: int) -> int:
    i = low - 1
    pivot = arr[high]
    for j in range(low, high):
        if arr[j] < pivot:
            i += 1
            arr[i], arr[j] = arr[j], arr[i]
    arr[i + 1], arr[high] = arr[high], arr[i + 1]
    return i + 1


def quicksort(arr: List[int], low: int, high: int):
    if low < high:
        pi = partition(arr, low, high)
        quicksort(arr, low, pi - 1)
        quicksort(arr, pi + 1, high)


if __name__ == '__main__':
    a = [1, 4, 3, 2, 8, 3, 6]
    quicksort(a, 0, len(a) - 1)
    print(a)
