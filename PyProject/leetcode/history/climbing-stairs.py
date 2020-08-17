class Solution:
    """
    假设你正在爬楼梯。需要 n 阶你才能到达楼顶。
    每次你可以爬 1 或 2 个台阶。你有多少种不同的方法可以爬到楼顶呢？


    1. n = 1 (1) res = 1
    2. n = 2 (1,1  2) res = 2
    3. n = 3 (n = 1 n= 2 或 n=2 n=1) 重复的 res = 3
    4. n = 4 (n=3 + n=2)   res = 5
    """

    def climbStairs(self, n: int) -> int:
        if n <= 2:
            return n

        # 使用迭代
        f1, f2, f3 = 1, 2, 3
        for i in range(3, n + 1):
            f3 = f1 + f2
            f1 = f2
            f2 = f3
        return f3

    # 使用缓存
    def climbStairs2(self, n: int) -> int:
        if n <= 2:
            return n
        f1, f2, f3 = 1, 2, 3
        for i in range(3, n + 1):
            f3 = f1 + f2
            f1 = f2
            f2 = f3
        return f3
