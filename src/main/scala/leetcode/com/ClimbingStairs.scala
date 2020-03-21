package leetcode.com

object ClimbingStairs {
  val map: Map[Int, Int] = Map[Int, Int]((1, 1), (2, 2))

  def climbStairs(n: Int): Int = {
    if (n <= 2) {
      return n
    }
    if (map.getOrElse(n, 0) == 0) {

    }
    return climbStairs(n - 1) + climbStairs(n - 2)


  }

}
