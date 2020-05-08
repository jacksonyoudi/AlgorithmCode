from typing import List


def spiral_order(matrix: List[List[int]]) -> List[int]:
    res = []
    while matrix:
        res += matrix.pop(0)
        # matrix = list(map(list, list(zip(*matrix))))[::-1]
        matrix = [list(i) for i in zip(*matrix)][::-1]

    return res


if __name__ == '__main__':
    a = [
        [1, 2, 3],
        [4, 5, 6],
        [7, 8, 9],
    ]

    b = [
        [1, 2, 3, 4],
        [5, 6, 7, 8],
        [9, 10, 11, 12],
        [13, 14, 15, 16],
    ]

    c = spiral_order(b)
    print(c)
