package com.allaboutscala.chapter2

/**
 * 定义类型
 */
object Types {
  def main(args: Array[String]): Unit = {
    // 不可以变
    val i: Int = 5
    val unit: Unit = ()

    lazy val aaa: String = "aaa"

    // 声明一个没有初始化的变量
    val a: String = ""
  }
}
