def partition(arr, low, high):
    # 标记指针
    i = (low - 1)
    # 参考数
    pivot = arr[high]

    for j in range(low, high):
        if arr[j] < pivot:
            i += 1
            arr[i], arr[j] = arr[j], arr[i]

    arr[i + 1], arr[high] = arr[high], arr[i + 1]
    return i + 1


def quick_sort(arr, low, high):
    if low < high:
        pi = partition(arr, low, high)
        quick_sort(arr, low, pi - 1)
        quick_sort(arr, pi + 1, high)


if __name__ == '__main__':
    a = [0, 1, 199, 19, 10, 3, 8, 5, 2]
    quick_sort(a, 0, len(a) - 1)
    print(a)
