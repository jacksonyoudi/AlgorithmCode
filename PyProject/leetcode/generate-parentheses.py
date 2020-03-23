from typing import List


class GenerateParentheses:
    def generateParenthesis(self, n: int) -> List[str]:
        res = []
        self.helper("", res, 0, 0, n)
        res.reverse()
        return res

    def is_valid(self, s: str) -> bool:
        q = []
        for i in s:
            if i == "(":
                q.append(i)
            else:
                if len(q) == 0:
                    return False
                q.pop()
        if len(q) == 1:
            return True
        else:
            return False

    def helper(self, s: str, res: List[str], left: int, right: int, max: int):
        # recursion terminator
        if left >= max and right >= max:
            if self.is_valid(s):
                res.append(s)
            return

            # process logic in the current level

        # drill down
        if left >= max:
            self.helper(s + ")", res, left, right + 1, max)
        elif right >= max:
            self.helper(s + "(", res, left + 1, right, max)
        else:
            self.helper(s + ")", res, left, right + 1, max)
            self.helper(s + "(", res, left + 1, right, max)

        # reverse the status the current if needed
