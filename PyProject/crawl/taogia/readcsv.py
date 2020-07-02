# -*- coding:utf-8 -*-

# author: youdi
# datetime: 2020/7/1 20:54
# software: IntelliJ IDEA

import csv

if __name__ == '__main__':
    with open('tab.csv', newline='') as csvfile:
        reader = csv.DictReader(csvfile)
        for row in reader:
            print(row.get("reportNo"))
