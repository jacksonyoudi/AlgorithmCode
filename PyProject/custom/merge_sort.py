def merge_sort(arr):
    if len(arr) > 1:
        mid = len(arr) // 2
        L = arr[:mid]
        R = arr[mid:]
        merge_sort(L)
        merge_sort(R)
        arr.clear()

        while len(L) > 0 and len(R) > 0:
            if L[0] > R[0]:
                arr.append(R[0])
                R.pop(0)
            else:
                arr.append(L[0])
                L.pop(0)
        for j in L:
            arr.append(j)
        for j in R:
            arr.append(j)


if __name__ == '__main__':
    a = [1, 9, 3, 0, 4, 10, 5, 4, 6]
    merge_sort(a)
    print(a)
