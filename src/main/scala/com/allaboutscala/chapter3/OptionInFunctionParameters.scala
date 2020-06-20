package com.allaboutscala.chapter3


/**
 * 如何将Option用作功能参数的一部分
 * 如何提供选项参数None的默认值
 * 调用函数时如何将Some（）传递给Option参数
 */
object OptionInFunctionParameters {
  def main(args: Array[String]): Unit = {
    println(s"""Total cost = ${calculateDonutCost("Glazed Donut", 5, Some("hello"))}""")
  }

  def calculateDonutCost(name: String, num: Int, code: Option[String] = None): Double = {
    println(s"${name},${num}")

    code match {
      case Some(c) => {
        println(c)
        val d: Double = 0.1
        return 2.50 * num * (1 - d)
      }
      case None => 2.50 * num
    }
  }

}
