package leetcode.com

//object ContainerWithMostWater {
//  def maxArea(height: Array[Int]): Int = {
//    var max: Int = 0
//    var maxTmp: Int = 0
//    var start: Int = 0
//    var end: Int = height.length - 1
//
//    while (start < end) {
//      if (height(start) < height(end)) {
//        maxTmp = height(start) * (end - start)
//        start += 1
//      } else {
//        maxTmp = height(end) * (end - start)
//        end -= 1
//      }
//
//      if (maxTmp > max) {
//        max = maxTmp
//      }
//
//    }
//    max
//  }
//
//}
