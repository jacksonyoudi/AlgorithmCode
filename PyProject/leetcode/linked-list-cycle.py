# Definition for singly-linked list.
class ListNode:
    def __init__(self, x):
        self.val = x
        self.next = None


class LinkedListCycle:
    def hasCycle(self, head: ListNode) -> bool:
        """
        使用双指针方法
        :param head:
        :return:
        """
        if head is None or head.next is None:
            return False

        slow = head
        fast = head.next

        while slow != fast:
            if slow is None or fast is None:
                return False
            slow = slow.next
            if fast.next is None:
                return False
            fast = fast.next.next
        return True
