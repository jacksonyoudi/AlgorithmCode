package leetcode.com

object MoveZeroes {
  def moveZeroes(nums: Array[Int]): Unit = {
    var j: Int = -1
    var tmp: Int = 0
    for (i <- 0 until nums.length) {

      if (nums(i) != 0) {
        j += 1;
        tmp = nums(i)
        nums(i) = nums(j)
        nums(j) = tmp
      }
    }
  }
}
