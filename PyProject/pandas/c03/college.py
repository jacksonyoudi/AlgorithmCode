# -*- coding:utf-8 -*-

# author: youdi
# datetime: 2020/7/7 23:28
# software: IntelliJ IDEA


import pandas as pd

if __name__ == '__main__':
    collage = pd.read_csv("../data/college.csv")
    print(collage.head())
