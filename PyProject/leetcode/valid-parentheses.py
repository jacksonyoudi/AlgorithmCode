class ValidParentheses:
    def isValid(self, s: str) -> bool:
        dic = {'{': '}', '[': ']', '(': ')', '?': '?'}
        stack = []
        for c in s:
            if c in dic.keys():
                stack.append(c)
            elif dic.get(stack.pop(), None) != c:
                return False
        return len(stack) == 0
