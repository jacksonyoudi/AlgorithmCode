# -*- coding:utf-8 -*-

# author: youdi
# datetime: 2020/7/14 22:46
# software: IntelliJ IDEA

import pandas as pd

if __name__ == '__main__':
    #  读取college数据集，行索引命名为INSTNM
    college = pd.read_csv('../data/college.csv', index_col='INSTNM')
    #  用索引方法get_loc，找到指定列的整数位置
    col_start = college.columns.get_loc('UGDS_WHITE')
    col_end = college.columns.get_loc('UGDS_UNKN') + 1

    #  用切片选取连续的列
    college.iloc[:5, col_start:col_end]
    #  index()方法可以获得整数行对应的标签名
    row_start = college.index[10]
    row_end = college.index[15]
