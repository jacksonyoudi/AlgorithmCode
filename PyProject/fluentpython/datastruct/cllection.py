from collections import deque
from typing import List


def main(a: List[int]):
    q = deque()
    for i in a:
        q.append(i)
        j = q.pop()
        print(j)


if __name__ == '__main__':
    a = [1, 3, 4, 4]
    main(a)
