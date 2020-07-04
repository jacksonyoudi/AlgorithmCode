# -*- coding:utf-8 -*-

# author: youdi
# datetime: 2020/7/2 23:00
# software: IntelliJ IDEA
import csv
from bs4 import BeautifulSoup
from selenium.webdriver.chrome.options import Options
from selenium import webdriver
import redis
import json


def get_gia_info(html):
    soup = BeautifulSoup(html, 'html.parser')
    infos = []
    trs = soup.find_all('tr')
    for tb in trs:
        sl = tb.select('span div')
        if not sl or len(sl) == 0:
            continue
        else:
            a = sl[0].string
            if a:
                key = a.string.strip("\n").strip()
            else:
                key = "unknown"
        strong = tb.find('strong')
        if strong:
            infos.append(strong.text)
            print("key:", key, "value:", strong.text)
    return infos


if __name__ == '__main__':
    rds = redis.Redis(host='localhost', port=6379, db=9)

    chrome_opt = Options()
    # chrome_opt.add_argument('--headless')
    # chrome_opt.add_argument('--disable-gpu')
    driver = webdriver.Chrome(options=chrome_opt)

    with open('tab_2.csv', newline='') as csvfile:
        reader = csv.DictReader(csvfile)
        for row in reader:
            report_no = row.get("reportNo")
            print("Report No", report_no)
            data = rds.get(report_no)
            if data:
                print("exists", data)
                continue
            driver.get("https://www.gia.edu/CN/report-check?reportno={}".format(report_no))
            infos = get_gia_info(driver.page_source)
            print(infos)
            if len(infos) > 0:
                print("write redis")
                rds.set(report_no, json.dumps(infos))
            # driver.close()
