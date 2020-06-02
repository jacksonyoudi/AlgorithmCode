package com.allaboutscala.chapter2

object WhileDe {
  def main(args: Array[String]): Unit = {
    var i: Int = 10

    while (i > 0) {
      println(i)
      i -= 1
    }

    var i1: Int = 0
    do {
      println(i1)
      i1 += 1
    } while (i1 < 10)
  }
}
