# -*- coding:utf-8 -*-

# author: youdi
# datetime: 2020/7/16 23:03
# software: IntelliJ IDEA

import numpy as np

if __name__ == '__main__':
    x = np.array([1, 2, 3, 4.0001, 5])
    y = np.array([1, 1.9999, 3, 4.01, 5.1])

    print(np.allclose(x, y))
    # 设置相对误差参数
    print(np.allclose(x, y, rtol=0.2))
    # 设置绝对误差
    print(np.allclose(x, y, atol=0.2))

    print(np.isclose(x, y))
    print(np.isclose(x, y, atol=0.2))
