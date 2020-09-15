class Solution:
    def scoreOfParentheses(self, S: str) -> int:
        stack = [0]

        for i in S:
            if i == '(':
                stack.append(0)
            else:
                x = stack.pop()
                stack[-1] += max(2*x, 1)

        return stack[0]

