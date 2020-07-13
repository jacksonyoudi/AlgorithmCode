# -*- coding:utf-8 -*-

# author: youdi
# datetime: 2020/7/13 22:49
# software: IntelliJ IDEA


import pandas as pd

if __name__ == '__main__':
    college = pd.read_csv('../data/college.csv', index_col='INSTNM')

    #  选取第61行
    college.iloc[60]

    #  也可以通过行标签选取
    college.loc['University of Alaska Anchorage']

    #  选取多个不连续的行
    college.iloc[[60, 99, 3]]
    labels = ['University of Alaska Anchorage',
              'International Academy of Hair Design',
              'University of Alabama in Huntsville']
    #  也可以用标签来选取
    college.loc[labels]

    #  iloc可以用切片连续选取
    college.iloc[99:102]

    #  loc可以用标签连续选取
    start = 'International Academy of Hair Design'
    stop = 'Mesa Community College'
    college.loc[start:stop]

    #  .index.tolist()可以直接提取索引标签，生成一个列表
    college.iloc[[60, 99, 3]].index.tolist()