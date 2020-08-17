class PowXN:
    def myPow(self, x: float, n: int) -> float:
        if x == 0 or x == 1.0:
            return x
        if n == 0:
            return 1
        if n < 0:
            n = -n
            x = 1/x

        if n <= 1:
            return x
        res = self.myPow(x, n // 2)
        if n % 2 == 1:
            return res * res * x
        else:
            return res * res
