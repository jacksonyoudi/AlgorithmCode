from typing import List


class GroupAnagrams:
    def groupAnagrams(self, strs: List[str]) -> List[List[str]]:
        res = []
        m = {}
        for s in strs:
            t = ''.join(sorted(s))
            if m.get(t, None):
                m.get(t).append(s)
            else:
                m[t] = [s]

        for v in m.values():
            res.append(v)
        return res.reverse()
