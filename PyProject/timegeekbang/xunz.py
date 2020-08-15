if __name__ == '__main__':
    a = [
        [1, 2, 3, 4],
        [5, 6, 7, 8],
        [9, 10, 11, 12],
        [13, 14, 15, 16],
    ]

    result = []
    while a:
        result.extend(a[0])
        a.pop(0)
        a = [list(i) for i in list(zip(*a))]
        a.reverse()
    print(result)
