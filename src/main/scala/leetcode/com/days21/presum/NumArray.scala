package leetcode.com.days21.presum

class NumArray(_nums: Array[Int]) {
  val ints: Array[Int] = _nums.scan(0)(_ + _)

  def sumRange(left: Int, right: Int): Int = {

    ints(right + 1) - ints(left)
  }
}
