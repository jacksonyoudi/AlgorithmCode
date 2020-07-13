# -*- coding:utf-8 -*-

# author: youdi
# datetime: 2020/7/13 21:53
# software: IntelliJ IDEA

import pandas as pd

if __name__ == '__main__':
    movie = pd.read_csv("../data/movie.csv")
    movie2 = movie[['movie_title', 'imdb_score', 'budget']]
    movie2.head()

    # 用nlargest方法，选出imdb_score分数最高的100个
    movie2.nlargest(100, 'imdb_score').head()

    # 用链式操作，nsmallest方法再从中挑出预算最小的五部
    movie2.nlargest(100, 'imdb_score').nsmallest(5, 'budget')
