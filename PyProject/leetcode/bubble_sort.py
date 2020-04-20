from typing import List


class BubbleSort:
    def bubble_sourt(l: List[int]) -> List[int]:
        for i in range(len(l) - 1):
            for j in range(len(l) - i - 1):
                if l[j] > l[j + 1]:
                    l[j], l[j + 1] = l[j + 1], l[j]
        return l


if __name__ == '__main__':
    a = [1, 4, 3, 6, 9, 2, 0, 7, 8, 5]
    a = BubbleSort.bubble_sourt(a)
    print(a)
