# -*- coding:utf-8 -*-

# author: youdi
# datetime: 2020/7/16 22:52
# software: IntelliJ IDEA

import numpy as np

if __name__ == '__main__':
    a = np.array([1, 2, 3, 4])
    b = np.array((1, 2, 3, 4, 5))
    np.array(range(10))
    np.array([[1, 2, 3], [4, 5, 6]])
    np.arange(8)
    np.linspace(0, 10, 11)
    np.linspace(0, 10, 11, endpoint=False)
    np.logspace(0, 100, 10)
    np.zeros(3)
    np.ones(3)
    np.zeros((3, 4))
    # 单位矩阵
    np.identity(2)
    np.empty((3, 4))

    np.random.randint(0, 50, 5)

    # 对角矩阵
    np.diag([1, 2, 3, 4])
