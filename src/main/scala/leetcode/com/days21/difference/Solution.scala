package leetcode.com.days21.difference

class Solution {
  def corpFlightBookings(bookings: Array[Array[Int]], n: Int): Array[Int] = {
    val nums: Array[Int] = new Array[Int](n)
    bookings.map(
      booking => {
        nums(booking(0) - 1) += booking(2)
        if (booking(1) < n) {
          nums(booking(1)) -= booking(2)
        }
      }
    )
    nums.drop(0).scan(0)(_ + _)

  }
}
