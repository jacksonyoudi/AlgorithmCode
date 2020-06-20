package com.allaboutscala.chapter3


/**
 *
 * 如何定义和使用没有参数且没有返回类型的函数
 * 如何定义和使用没有副作用的函数
 * 如何定义和使用没有返回类型的函数
 * 定义无返回类型的函数时如何使用类型推断
 */
object CreateFunction {
  def main(args: Array[String]): Unit = {
    val myFavoriteDonut = favoriteDonut()
    println(s"My favorite donut is $myFavoriteDonut")

    println("\nStep 2: How to define and use a function with no parenthesis")
    println(s"My least favorite donut is $leastFavoriteDonut")


    println("\nStep 3: How to define and use a function with no return type")

    printDonutSalesReport()
  }

  def favoriteDonut(): String = {
    "Glazed Donut"
  }

  def leastFavoriteDonut = "Plain Donut"

  def printDonutSalesReport(): Unit = {
    // lookup sales data in database and create some report
    println("Printing donut sales report... done!")
  }
}