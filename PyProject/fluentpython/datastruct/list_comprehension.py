# -*- coding:utf-8 -*-

# author: youdi
# datetime: 2020/5/31 15:59
# software: IntelliJ IDEA


if __name__ == '__main__':
    symbols = 'abcdef'
    codes = [ord(i) for i in symbols if ord(i) > 100]
    print(codes)

    c = list(filter(lambda c: c > 100, map(ord, symbols)))
    print(c)

    colors = ['black', 'white']
    sizes = ['s', 'm', 'l']
    selects = [(color, size) for color in colors for size in sizes]
    print(selects)
