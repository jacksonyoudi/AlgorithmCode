package leetcode.com

import scala.util.control.Breaks

class ListNode(var _x: Int = 0) {
  var next: ListNode = null
  var x: Int = _x
}


/**
 * 使用双指针方法
 */

object LinkedListCycleII {
  def detectCycle(head: ListNode): ListNode = {
    if (head == null || head.next == null) {
      return null
    }
    var slow: ListNode = head
    var fast: ListNode = head.next
    var hasCycle: Boolean = false

    Breaks.breakable {
      while (slow != null && fast != null) {
        if (slow.eq(fast)) {
          hasCycle = true
          Breaks.break
        }

        slow = slow.next

        if (fast.next == null) {
          return null
        }
        fast = fast.next.next
      }
    }

    if (hasCycle) {
      slow = head
      while (!slow.eq(fast)) {
        slow = slow.next
        fast = fast.next
      }
      return slow
    }
    null
  }
}
