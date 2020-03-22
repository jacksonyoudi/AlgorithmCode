package leetcode.com


object MergeTwoSortedLists {
  def mergeTwoLists(l1: ListNode, l2: ListNode): ListNode = {
    var one: ListNode = l1
    var two: ListNode = l2
    if (one != null && two != null) {
      if (one.x > two.x) {
        var l: ListNode = one
        one = two
        two = l
      }
      one.next = this.mergeTwoLists(one.next, two)
    }
    if (one != null) {
      one
    } else {
      two
    }
  }
}
