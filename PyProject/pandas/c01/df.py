# -*- coding:utf-8 -*-

# author: youdi
# datetime: 2020/7/5 22:10
# software: IntelliJ IDEA

import pandas as pd

if __name__ == '__main__':
    df = pd.read_csv("../gdp.csv")
    print(df)
    print(df.columns)
    print(df.dtypes)
    print(df.shape)
    print(df.index)
