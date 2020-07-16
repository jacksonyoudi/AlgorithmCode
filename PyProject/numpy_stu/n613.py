# -*- coding:utf-8 -*-

# author: youdi
# datetime: 2020/7/16 23:09
# software: IntelliJ IDEA

import numpy as np

if __name__ == '__main__':
    x = np.arange(8)
    np.append(x, 8)
    np.append(x, [9, 10])
    np.insert(x, 1, 11)
    x[3] = 12
    x = np.array([[1, 2, 3], [4, 5, 6]])
    x[0, 2] = 4
    x[1:, 1:] = 1
    x[1:, 1:] = [1, 2]
    x[1:, 1:] = [[1, 2], [3, 4]]
