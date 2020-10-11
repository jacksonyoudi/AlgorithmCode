class MyQueue:

    def __init__(self):
        """
        Initialize your data structure here.
        """
        self.in_stack = list()
        self.out_stack = list()

    def push(self, x: int) -> None:
        """
        Push element x to the back of queue.
        """

        if self.out_stack:
            self.in_stack.insert(0, x)
            return
        while self.in_stack:
            self.out_stack.insert(0, self.in_stack.pop(0))

        self.in_stack.insert(0, x)

    def pop(self) -> int:
        """
        Removes the element from in front of queue and returns that element.
        """
        if self.out_stack:
            return self.out_stack.pop(0)
        else:
            if self.in_stack:
                while self.in_stack:
                    self.out_stack.insert(0, self.in_stack.pop(0))
                return self.out_stack.pop(0)
            else:
                return None

    def peek(self) -> int:
        """
        Get the front element.
        """

        if self.out_stack:
            return self.out_stack[0]
        else:
            if self.in_stack:
                while self.in_stack:
                    self.out_stack.insert(0, self.in_stack.pop(0))
                return self.out_stack[0]
            else:
                return None

    def empty(self) -> bool:
        """
        Returns whether the queue is empty.
        """
        return not self.out_stack and not self.in_stack

# Your MyQueue object will be instantiated and called as such:
# obj = MyQueue()
# obj.push(x)
# param_2 = obj.pop()
# param_3 = obj.peek()
# param_4 = obj.empty()
