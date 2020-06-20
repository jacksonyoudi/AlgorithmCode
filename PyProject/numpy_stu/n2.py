import numpy as np


def m():
    # 使用[]创建ndarray
    a = np.array([1, 2, 3, 4, 5])
    b = np.array([[1, 2, 3], [4, 5, 6]], dtype=np.float)

    # dtype ndmin
    print(a)


if __name__ == '__main__':
    m()
