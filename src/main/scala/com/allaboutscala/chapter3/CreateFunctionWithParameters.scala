package com.allaboutscala.chapter3

/**
 * 如何用参数定义功能
 * 如何使用参数调用函数
 * 如何向功能参数添加默认值
 * 如何使用具有默认值的参数调用函数
 */
object CreateFunctionWithParameters {
  def main(args: Array[String]): Unit = {
    println(this.calculate("a", 10))
    println(this.calculate2(10, "a"))
    println(this.calculate2(30))
  }


  def calculate(name: String, num: Int): Double = {
    println("Calculate:", name)
    return 2.50 * num
  }

  def calculate2(num: Int, name: String = "hello"): Double = {
    println("name", name)
    2.50 * num
  }
}

