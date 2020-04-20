from typing import List


class Solution:
    def canFinish(self, numCourses: int, prerequisites: List[List[int]]) -> bool:
        # 标记是否有走过
        visit = [0 for _ in range(numCourses)]

        def is_visit(i):
            # 没有走过 mark 0
            # 正在走， mark -1
            # 已经走过标记为 1

            # have ring
            if visit[i] == -1:
                return False

            # 上次走过
            if visit[i] == 1:
                return True

            visit[i] = -1
            for j in prerequisites[i]:
                if not is_visit(j):
                    return False
            visit[i] = 1
            return True

        for i in range(numCourses):
            if not is_visit(i):
                return False
        return True
