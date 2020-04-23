def partition(arr, low, high):
    i = low - 1

    pivot = arr[high]
    for j in range(low, high):
        if arr[j] < pivot:
            arr[i + 1], arr[j] = arr[j], arr[i + 1]
            i += 1

    arr[i + 1], arr[high] = arr[high], arr[i + 1]
    return i + 1


def quick_sort(arr, low, high):
    if low < high:
        pi = partition(arr, low, high)
        quick_sort(arr, low, pi - 1)
        quick_sort(arr, pi + 1, high)


if __name__ == '__main__':
    arr = [1, 5, 2, 8, 0, 6, 3]
    quick_sort(arr, 0, len(arr) - 1)
    print(arr)
