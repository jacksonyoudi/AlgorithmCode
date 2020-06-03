package com.allaboutscala.chapter2


/**
 * 如何声明任何类型的变量
 * 如何声明AnyRef类型的变量
 * 如何声明类型AnyVal的变量
 */
object ScalaType {
  def main(args: Array[String]): Unit = {
    println("Step 1: Declare a variable of type Any")
    val favoriteDonut: Any = "Glazed Donut"
    println(s"favoriteDonut of type Any = $favoriteDonut")

    println("\nStep 2: Declare a variable of type AnyRef")
    val donutName: AnyRef = "Glazed Donut"
    println(s"donutName of type AnyRef = $donutName")

    println("\nStep 3: Declare a variable of type AnyVal")
    val donutPrice: AnyVal = 2.50
    println(s"donutPrice of type AnyVal = $donutPrice")
  }
}
