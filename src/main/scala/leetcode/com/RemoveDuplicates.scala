package leetcode.com


object RemoveDuplicates {
  def removeDuplicates(nums: Array[Int]): Int = {
    if (nums.isEmpty) {
      return 0
    } else if (nums.length == 1) {
      return 1
    }
    var i: Int = 0
    var j: Int = 1

    while (j < nums.length) {
      if (nums(j) == nums(i)) {
        j += 1
      } else {
        i += 1
        nums(i) = nums(j)
        j += 1
      }
    }
    i + 1

  }

}
