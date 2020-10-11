from typing import List


class Solution:
    def generateMatrix(self, n: int) -> List[List[int]]:
        datas = [list(range(n * i + 1, n * i + n + 1)) for i in range(0, n)]
        result = []
        while datas:
            result.append(datas.pop(0))
            datas = list(map(list, zip(*datas)))[::-1]
        return result


if __name__ == '__main__':
    n = 3
    a = [list(range(n * i + 1, n * i + n + 1)) for i in range(0, n)]
    print(a)
