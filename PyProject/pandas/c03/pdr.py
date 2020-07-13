# -*- coding:utf-8 -*-

# author: youdi
# datetime: 2020/7/13 22:28
# software: IntelliJ IDEA

import pandas_datareader as pdr

if __name__ == '__main__':
    tsla = pdr.DataReader('tsla', data_source='yahoo', start='2017-1-1')
    # 只关注每天的收盘价，使用cummax得到迄今为止的收盘价最大值
    tsla_close = tsla['Close']
    tsla_cummax = tsla_close.cummax()

    #  将下行区间限制到10%，将tsla_cummax乘以0.9
    tsla_trailing_stop = tsla_cummax * .9
    tsla_trailing_stop.head(8)
