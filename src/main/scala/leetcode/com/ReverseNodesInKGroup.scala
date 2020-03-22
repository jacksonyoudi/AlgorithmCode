package leetcode.com

object ReverseNodesInKGroup {
  def reverseKGroup(head: ListNode, k: Int): ListNode = {
    if (head == null) {
      return null
    }
    var cur: ListNode = head
    var start: ListNode = head
    var tail: ListNode = head
    //    terminat and current logic
    for (_ <- 0 until k) {
      if (cur == null) {
        return head
      }

      val nextNode: ListNode = cur.next
      cur.next = start
      start = cur
      cur = nextNode
    }

    tail.next = this.reverseKGroup(cur, k)
    start
  }
}
