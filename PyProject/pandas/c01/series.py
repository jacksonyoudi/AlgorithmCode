# -*- coding:utf-8 -*-

# author: youdi
# datetime: 2020/6/20 15:20
# software: IntelliJ IDEA


import pandas as pd

if __name__ == '__main__':
    s = pd.Series([1, 2, 3, 4, 5], index=["a", "b", "c", "d", "e"])
    print(s.index)
    print(s.values)
    s1 = pd.Series([1, 2, 3, 4, 5])
    print(s1)
    df = pd.read_csv("../gdp.csv")
    print(df)
