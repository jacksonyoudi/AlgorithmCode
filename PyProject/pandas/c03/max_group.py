# -*- coding:utf-8 -*-

# author: youdi
# datetime: 2020/7/13 22:04
# software: IntelliJ IDEA


import pandas as pd

if __name__ == '__main__':
    movie = pd.read_csv("../data/movie.csv")
    movie2 = movie[['movie_title', 'title_year', 'imdb_score']]
    movie2.sort_values('title_year', ascending=False).head()

    # 用列表同时对两列进行排序
    movie3 = movie2.sort_values(['title_year', 'imdb_score'], ascending=False)

    #  用drop_duplicates去重，只保留每年的第一条数据
    movie3.drop_duplicates(subset='title_year')

    #  通过给ascending设置列表，可以同时对一列降序排列，一列升序排列
    movie4 = movie[['movie_title', 'title_year', 'content_rating', 'budget']]
    movie4_sorted = movie4.sort_values(['title_year', 'content_rating', 'budget'],
                                       ascending=[False, False, True])

    movie4_sorted.drop_duplicates(subset=['title_year', 'content_rating']).head(10)
    
