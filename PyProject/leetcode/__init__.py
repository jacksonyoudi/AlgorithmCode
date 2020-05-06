if __name__ == '__main__':
    a = [
        [1, 2, 3],
        [4, 5, 6],
        [7, 8, 9]
    ]

    res = []
    while a:
        res += a.pop(0)
        a = list(map(list, zip(*a)))[::-1]

    print(res)
