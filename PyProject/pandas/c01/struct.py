# -*- coding:utf-8 -*-

# author: youdi
# datetime: 2020/6/18 08:35
# software: IntelliJ IDEA


import numpy as np
import pandas as pd

if __name__ == '__main__':
    s = pd.Series([1, 3, 5, np.nan, 6, 8])
    print(s.index)