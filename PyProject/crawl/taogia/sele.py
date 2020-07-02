# -*- coding:utf-8 -*-

# author: youdi
# datetime: 2020/6/29 22:25
# software: IntelliJ IDEA

from bs4 import BeautifulSoup
from selenium import webdriver
from selenium.webdriver.chrome.options import Options
import time

if __name__ == '__main__':
    chrome_opt = Options()
    chrome_opt.add_argument('--headless')
    chrome_opt.add_argument('--disable-gpu')
    driver = webdriver.Chrome(options=chrome_opt)
    driver.get("https://www.gia.edu/CN/report-check?reportno=2201951132")
    time.sleep(10)
    print("hello")
    html = driver.page_source
    print(html)
    soup = BeautifulSoup(html, 'html.parser')
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
            print("key:", key, "value:", strong.text)
# key: 尺寸 value: 4.91 - 4.96 x 3.01 mm
# key: 克拉重量 value: 0.45 carat
# key: unknown value: E
# key: unknown value: VVS2
# key: unknown value: Excellent
# key: unknown value: 61.0
# key: unknown value: 60
# key: unknown value: 33.0°
# key: 冠高 value: 13.0%
# key: 亭部角度 value: 41.6°
# key: 亭深 value: 44.5%
# key: 星形刻面长度 value: 50%
# key: 腰下刻面 value: 75%
# key: 腰围 value: Medium to Slightly Thick, Faceted, 4.0%
# key: 尖底 value: None
# key: unknown value: Excellent
# key: 对称性 value: Excellent
# key: 荧光 value: None
# key: 净度特征 value: Pinpoint, Feather