# -*- coding:utf-8 -*-

# author: youdi
# datetime: 2020/7/6 08:28
# software: IntelliJ IDEA

import pandas as pd
import numpy as np

if __name__ == '__main__':
    df1 = pd.DataFrame(
        np.random.randint(0, 10, (4, 2)),
        index=['A', 'B', 'C', 'D'],
        columns=['a', 'b'])
    df2 = pd.DataFrame({
        "a": [1, 2, 3, 4],
        "b": [5, 6, 7, 8]
    },
        index=["A", "B", "C", "D"]
    )

    arr = np.array(
        [("item1", 1), ("item2", 2), ("item3", 3), ("item4", 4)],
        dtype=[("name", "10S"), ("count", int)]
    )
    df3 = pd.DataFrame(arr)
    print(df1)
    print(df2)
    print(df3)

    dict1 = {
        "a": [1, 2, 3],

    }
