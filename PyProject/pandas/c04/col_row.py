# -*- coding:utf-8 -*-

# author: youdi
# datetime: 2020/7/13 22:56
# software: IntelliJ IDEA

import pandas as pd

if __name__ == '__main__':
    #  读取college数据集，给行索引命名为INSTNM；选取前3行和前4列
    college = pd.read_csv('../data/college.csv', index_col='INSTNM')
    # index
    college.iloc[:3, :4]

    #  用loc实现同上功能
    college.loc[:'Amridge University', :'MENONLY']

    #  选取两列的所有的行
    college.iloc[:, [4, 6]].head()

    #  loc实现同上功能
    college.loc[:, ['WOMENONLY', 'SATVRMID']]

    #  选取不连续的行和列
    college.iloc[[100, 200], [7, 15]]

    #  用loc和列表，选取不连续的行和列
    rows = ['GateWay Community College', 'American Baptist Seminary of the West']
    columns = ['SATMTMID', 'UGDS_NHPI']

    college.loc[rows, columns]

    #  iloc选取一个标量值
    college.iloc[5, -4]

    #  loc选取一个标量值
    college.loc['The University of Alabama', 'PCTFLOAN']

    #  iloc对行切片，并只选取一列
    college.iloc[90:80:-2, 5]

    #  loc对行切片，并只选取一列
    start = 'Empire Beauty School-Flagstaff'
    stop = 'Arizona State University-Tempe'
    college.loc[start:stop:-2, 'RELAFFIL']
