# -*- coding:utf-8 -*-

# author: youdi
# datetime: 2020/7/2 23:00
# software: IntelliJ IDEA
import csv
from bs4 import BeautifulSoup
from selenium.webdriver.chrome.options import Options
from selenium import webdriver


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
    chrome_opt = Options()
    # chrome_opt.add_argument('--headless')
    # chrome_opt.add_argument('--disable-gpu')
    driver = webdriver.Chrome(options=chrome_opt)

    with open('tab.csv', newline='') as csvfile:
        reader = csv.DictReader(csvfile)
        for row in reader:
            report_no = row.get("reportNo")
            print("Report No", report_no)
            driver.get("https://www.gia.edu/CN/report-check?reportno={}".format(report_no))
            infos = get_gia_info(driver.page_source)
            print(infos)
            # driver.close()
