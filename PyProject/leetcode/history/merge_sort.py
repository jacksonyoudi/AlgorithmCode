from typing import List


def merge_sort(arr: List[int]):
    if len(arr) > 1:
        mid = len(arr) // 2
        L = arr[:mid]
        R = arr[mid:]
        merge_sort(L)
        merge_sort(R)

        arr.clear()

        while len(L) > 0 and len(R) > 0:
            if L[0] < R[0]:
                arr.append(L[0])
                L.pop(0)
            else:
                arr.append(R[0])
                R.pop(0)
        for i in L:
            arr.append(i)
        for j in R:
            arr.append(j)


if __name__ == '__main__':
    a = [1, 4, 3, 2, 8, 3, 6]
    merge_sort(a)
    print(a)
