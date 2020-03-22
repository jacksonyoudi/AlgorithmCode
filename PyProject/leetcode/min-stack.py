class MinStack:

    def __init__(self):
        """
        initialize your data structure here.
        """
        self._stack = []
        self._count = 0

    def push(self, x: int) -> None:
        self._stack.append(x)
        self._count += 1

    def pop(self) -> None:
        if self._count == 0:
            return
        self._stack.pop()
        self._count -= 1

    def top(self) -> int:
        if self._count >= 1:
            return self._stack[self._count - 1]
        else:
            return None

    def getMin(self) -> int:
        return min(self._stack)
