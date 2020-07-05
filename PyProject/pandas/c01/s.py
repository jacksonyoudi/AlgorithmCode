# -*- coding:utf-8 -*-

# author: youdi
# datetime: 2020/7/5 21:59
# software: IntelliJ IDEA


import pandas as pd

if __name__ == '__main__':
    s = pd.Series(range(1, 6), index=['a', 'b', 'c', 'd', 'e'])
    print("index", s.index)
    print("values", s.values)
    print(s[1:3])
    print(s['b':'c'])
    print(s.get(1))
    print(dict(s.items()))

