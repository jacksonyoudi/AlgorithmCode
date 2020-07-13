# -*- coding:utf-8 -*-

# author: youdi
# datetime: 2020/7/13 22:18
# software: IntelliJ IDEA


import pandas as pd

if __name__ == '__main__':
    movie = pd.read_csv('../data/movie.csv')
    movie2 = movie[['movie_title', 'imdb_score', 'budget']]
    movie_smallest_largest = movie2.nlargest(100, 'imdb_score').nsmallest(5, 'budget')

    #  用sort_values方法，选取imdb_score最高的100个
    movie.sort_values('imdb_score', ascending=False).head(100) \
        .sort_values("budget").head(5)

    movie2.nlargest(100, 'imdb_score').tail()

    movie2.sort_values('imdb_score', ascending=False).head(100).tail()
