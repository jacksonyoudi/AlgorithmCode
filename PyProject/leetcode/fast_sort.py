def partition(arr, low, high):
    i = (low - 1)
    # 参考数 取最大的
    pivot = arr[high]

    for j in range(low, high):
        # 如果 < pivot, 指针后移,与小区间数交换
        if arr[j] < pivot:
            # increment index of smaller element
            i = i + 1
            arr[i], arr[j] = arr[j], arr[i]

    # 将pivot放在中间，就可以了
    arr[i + 1], arr[high] = arr[high], arr[i + 1]
    return (i + 1)


def quickSort(arr, low, high):
    if low < high:
        pi = partition(arr, low, high)
        quickSort(arr, low, pi - 1)
        quickSort(arr, pi + 1, high)


if __name__ == '__main__':
    a = [1, 5, 2, 3, 0]
    quickSort(a, 0, len(a) - 1)
    print(a)
