if __name__ == '__main__':
    matrix = [
        [1, 2, 3],
        [4, 5, 6],
        [7, 8, 9]
    ]
    matrix = list(map(list, zip(*matrix)))[:]
    for i, num in enumerate(matrix):
        num.reverse()
        matrix[i] = num
    print(matrix)