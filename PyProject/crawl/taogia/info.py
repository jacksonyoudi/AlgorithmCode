# -*- coding:utf-8 -*-

# author: youdi
# datetime: 2020/6/27 18:18
# software: IntelliJ IDEA

import requests
import csv
import time


def get_list():
    pass


if __name__ == '__main__':
    headers = {
        "Accept": "application/json, text/javascript, */*; q=0.01",
        "Accept-Encoding": "gzip, deflate",
        "Accept-Language": "zh,en-US;q=0.9,en;q=0.8,zh-CN;q=0.7",
        "Connection": "keep-alive",
        "Content-Length": "406",
        "Content-Type": "application/x-www-form-urlencoded; charset=UTF-8",
        "Cookie": "JSESSIONID=85A3C31F84123D4B1454D6D6641881ED; USER=4263632320A180B1531FC69A6D2F8B8652E3DCFABD9605310641C3FFE3898398EF496960F1F0D5B5E6D08927A311858CCEB4C187323FF445AD3809EFED51DE533C7DBB68ACA78950F8127998F3F0FDC77E734111F2548295",
        "Host": "www.taogia.com",
        "Origin": "http://www.taogia.com",
        "Referer": "http://www.taogia.com/newManager/template/diamond/diamsearch.html?t=1593265126668",
        "User-Agent": "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_5) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/83.0.4103.116 Safari/537.36",
        "X-Requested-With": "XMLHttpRequest",
    }

    params = {
        "uid": "1587547404780001",
        "sign": "de8a1a5f42b253fff654785d70afaee7",
        "www": "www.taogia.com",
        "dollar_rate": 7.25,
        "order_tj": 1,
        "order_sx": 1,
        "is_accurate_and_detail": 1,
        "detail": 10402,
        "rap_ids": "",
        "d_size_min": 0.52,
        "d_size_max": 0.56,
        "shape": "%E5%9C%86%E5%BD%A2",
        "cut": "EX",
        "color": "E,F",
        "clarity": "FL,IF,VVS1,VVS2",
        "polish": "EX",
        "sym": "EX",
        "flour": "NON",
        "cert": "GIA",
        "milky": "N",
        "browness": "N",
        "sys_status": -1,
        "currentPage": 2,
        "pageSize": 200
    }

    rep = requests.post(
        "http://www.taogia.com/diamondCert/pc/diamondCert_interface_FindPageList_many_all_view_manager_specialty.xhtml",
        params, headers=headers)
    print(rep.status_code)
    data = rep.json().get("list")
    csv_txt = ['id', 'shape', 'd_size', 'color', 'clarity', 'cut', 'polish', 'sym', 'flour', 'm1', 'm2', 'm3',
               'd_depth', 'd_table', 'd_ref', 'reportNo', 'detail', 'disc', 'rate', 'location', 'certNo', 'milky',
               'eyeClean', 'browness', 'isbuy', 'los', 'isdisp', 'cert', 'types', 'disc1', 'oldRef', 'rap',
               'update_time', 'sys_status', 'green', 'up_file_name', 'is_sh', 'certificate', 'daylight', 'order_by',
               'is_size_normal', 'black', 'is_bgm', 'bt', 'bc', 'wt', 'wc', 'temp', 'order_tj', 'order_sx', 'rate_add',
               'count_rate', 'insert_type', 'currentPage', 'pageSize', 'totalPage', 'totalRecord', 'shape_en',
               'location_en', 'is_own_filter_stock', 'num_all', 'isown', 'app_check_code', 'sign', 'sessionkey',
               'status', 'girdle', 'video', 'cgno', 'remark', 'disc2', 'rate1']
    print(csv_txt)
    with open("tab.csv", 'w') as csvfile:
        writer = csv.DictWriter(csvfile, fieldnames=csv_txt)
        writer.writeheader()
        for i in data:
            writer.writerow(i)
