# -*- coding:utf-8 -*-

# author: youdi
# datetime: 2020/7/13 22:34
# software: IntelliJ IDEA


import pandas as pd

if __name__ == '__main__':
    college = pd.read_csv('../data/college.csv', index_col='INSTNM')
    city = college['CITY']
    #  iloc可以通过整数选取
    city.iloc[3]
    city.iloc[[10, 20, 30]]

    #  loc只接收行索引标签
    city.loc['Heritage Christian University']

    # 也可以通过切片语法均匀选择多个
    city.loc['Alabama State University':'Reid State Technical College':10]

    #  也可以不使用loc，直接使用类似Python的语法
    city['Alabama State University':'Reid State Technical College':10]

    #  要想只选取一项，并保留其Series类型，则传入一个只包含一项的列表
    city.iloc[[3]]

    #  使用loc切片时要注意，如果start索引再stop索引之后，则会返回空，并且不会报警
    city.loc['Reid State Technical College':'Alabama State University':10]

    #  也可以切片逆序选取
    city.loc['Reid State Technical College':'Alabama State University':-10]
