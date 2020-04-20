from typing import List


def merge(arr: List[int]):
    if len(arr) > 1:
        mid = len(arr) // 2
        L = arr[:mid]
        R = arr[mid:]

        merge(L)
        merge(R)

        i = j = k = 0

        # 合并两个数组
        while i < len(L) and j < len(R):
            if L[i] < R[j]:
                arr[k] = L[i]
                i += 1
            else:
                arr[k] = R[j]
                j += 1
            k += 1

        while i < len(L):
            arr[k] = L[i]
            k += 1
            i += 1
        while j < len(R):
            arr[k] = R[j]
            k += 1
            j += 1


if __name__ == '__main__':
    a = [1, 4, 3, 2, 8, 3, 6]
    merge(a)
    print(a)
