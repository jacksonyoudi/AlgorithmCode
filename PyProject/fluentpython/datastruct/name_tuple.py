# -*- coding:utf-8 -*-

# author: youdi
# datetime: 2020/5/31 16:31
# software: IntelliJ IDEA


from collections import namedtuple

City = namedtuple("City", "name country population coordinates")
if __name__ == '__main__':
    tokyo = City('Tokyo', 'JP', 36.933, (35, 139))
    for i in tokyo:
        print(i)
