package org.youdi.custom

object RangeIn {
  def main(args: Array[String]): Unit = {
    val i: Int = 11

    val inclusive: Range.Inclusive = 1 to 12
    if (inclusive.toArray.contains(i)) {
      println("hello")
    }
  }

}
