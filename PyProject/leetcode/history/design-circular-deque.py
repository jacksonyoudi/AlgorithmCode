class MyCircularDeque:

    def __init__(self, k: int):
        """
        Initialize your data structure here. Set the size of the deque to be k.
        """
        self._deque = []
        self._size = k

    def insertFront(self, value: int) -> bool:
        """
        Adds an item at the front of Deque. Return true if the operation is successful.
        """
        if len(self._deque) == self._size:
            return False
        self._deque.insert(0, value)
        return True

    def insertLast(self, value: int) -> bool:
        """
        Adds an item at the rear of Deque. Return true if the operation is successful.
        """
        if len(self._deque) == self._size:
            return False
        self._deque.append(value)
        return True

    def deleteFront(self) -> bool:
        """
        Deletes an item from the front of Deque. Return true if the operation is successful.
        """
        if len(self._deque) == 0:
            return False
        self._deque.pop(0)
        return True

    def deleteLast(self) -> bool:
        """
        Deletes an item from the rear of Deque. Return true if the operation is successful.
        """
        if len(self._deque) == 0:
            return False
        self._deque.pop()
        return True

    def getFront(self) -> int:
        """
        Get the front item from the deque.
        """
        if len(self._deque) == 0:
            return None
        return self._deque[0]

    def getRear(self) -> int:
        """
        Get the last item from the deque.
        """
        if len(self._deque) == 0:
            return None
        return self._deque[-1]

    def isEmpty(self) -> bool:
        """
        Checks whether the circular deque is empty or not.
        """
        if len(self._deque) == 0:
            return True
        return False

    def isFull(self) -> bool:
        """
        Checks whether the circular deque is full or not.
        """
        if len(self._deque) == self._size:
            return True
        return False
