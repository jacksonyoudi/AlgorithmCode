class Solution:
    def isValid(self, s: str) -> bool:
        pars = {
            '(': ')',
            '{': '}',
            '[': ']'
        }

        l = list()

        for i in s:
            if i in pars.keys():
                l.insert(0, i)
            else:
                if i == pars.get(l[0]):
                    l.pop(0)
                else:
                    return False
        if l:
            return False
        else:
            return True
