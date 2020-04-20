package org.youdi.custom

object OP {
  def findJudge(N: Int, trust: Array[Array[Int]]): Int = {
    val ints = new Array[Int](N + 1)
    for (elem: Array[Int] <- trust) {
      ints(elem(0)) -= 1
      ints(elem(1)) += 1
    }

    for (i <- 1 to N) {
      if (ints(i).equals(N - 1)) {
        return i
      }
    }
    -1
  }
}
