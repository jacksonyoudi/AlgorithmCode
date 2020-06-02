package com.allaboutscala.chapter2

import scala.collection.immutable.NumericRange

object RangleD {
  def main(args: Array[String]): Unit = {
    println("Step 1: Create a simple numeric range from 1 to 5 inclusive")
    val inclusive: Range.Inclusive = 1 to 5
    println(s"Range from 1 to 5 inclusive = $inclusive")


    val range: Range = 1 to 10 by 2
    val chars: NumericRange[Char] = 'a' to 'z' by 2

    val list: List[Int] = (1 to 10).toList
    (1 to 10).foreach(println)
  }
}
