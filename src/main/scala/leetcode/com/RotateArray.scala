package leetcode.com

/**
 * 移动
 */
object RotateArray {
  def rotate(nums: Array[Int], k: Int): Unit = {
    if (nums.length == 1) {
      return
    }

    val i: Int = k % nums.length
    if (i == 0) {
      return
    }
    // 0- i-1
    // i - nums.length -1

    val values: Array[Int] = new Array(i)
    for (j <- (nums.length-1) to (i + 1)) {
      values(j - i - 1) = nums(j)
    }

    for (j <- 0 until i) {
      nums(i + j) = nums(j)
      nums(j) = values(j)
    }
  }
}
