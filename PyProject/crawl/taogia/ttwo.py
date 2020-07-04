# -*- coding:utf-8 -*-

# author: youdi
# datetime: 2020/7/1 20:54
# software: IntelliJ IDEA

import csv
import redis
import json

if __name__ == '__main__':
    m = ['col_1',
         'col_2',
         'col_3',
         'col_4',
         'col_5',
         'col_6',
         'col_7',
         'col_8',
         'col_9',
         'col_10',
         'col_11',
         'col_12',
         'col_13',
         'col_14',
         'col_15',
         'col_16',
         'col_17',
         'col_18',
         'col_19']

    m_null = {'col_1': '',
              'col_2': '',
              'col_3': '',
              'col_4': '',
              'col_5': '',
              'col_6': '',
              'col_7': '',
              'col_8': '',
              'col_9': '',
              'col_10': '',
              'col_11': '',
              'col_12': '',
              'col_13': '',
              'col_14': '',
              'col_15': '',
              'col_16': '',
              'col_17': '',
              'col_18': '',
              'col_19': ''}

    headers = ['id',
               'shape',
               'd_size',
               'color',
               'clarity',
               'cut',
               'polish',
               'sym',
               'flour',
               'm1',
               'm2',
               'm3',
               'd_depth',
               'd_table',
               'd_ref',
               'reportNo',
               'detail',
               'disc',
               'rate',
               'location',
               'certNo',
               'milky',
               'eyeClean',
               'browness',
               'isbuy',
               'los',
               'isdisp',
               'cert',
               'types',
               'disc1',
               'oldRef',
               'rap',
               'update_time',
               'sys_status',
               'green',
               'up_file_name',
               'is_sh',
               'certificate',
               'daylight',
               'order_by',
               'is_size_normal',
               'black',
               'is_bgm',
               'bt',
               'bc',
               'wt',
               'wc',
               'temp',
               'order_tj',
               'order_sx',
               'rate_add',
               'count_rate',
               'insert_type',
               'currentPage',
               'pageSize',
               'totalPage',
               'totalRecord',
               'shape_en',
               'location_en',
               'is_own_filter_stock',
               'num_all',
               'isown',
               'app_check_code',
               'sign',
               'sessionkey',
               'status',
               'girdle',
               'video',
               'cgno',
               'remark',
               'disc2',
               'rate1',
               'col_1',
               'col_2',
               'col_3',
               'col_4',
               'col_5',
               'col_6',
               'col_7',
               'col_8',
               'col_9',
               'col_10',
               'col_11',
               'col_12',
               'col_13',
               'col_14',
               'col_15',
               'col_16',
               'col_17',
               'col_18',
               'col_19']
    rds = redis.Redis(host='localhost', port=6379, db=9)
    infos = []
    with open('tab_1.csv', newline='') as csvfile:
        reader = csv.DictReader(csvfile)
        for row in reader:
            report_no = row.get("reportNo")
            data = rds.get(report_no)
            d = dict(row)
            if data is not None:
                ext_data = json.loads(data)

                d.update(dict(zip(m, ext_data)))
            else:
                d.update(m_null)
            infos.append(d)

    with open("all_tab.csv", 'w') as csvwrfile:
        writer = csv.DictWriter(csvwrfile, fieldnames=headers)
        writer.writeheader()
        for i in infos:
            writer.writerow(i)
