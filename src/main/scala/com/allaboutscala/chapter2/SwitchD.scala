package com.allaboutscala.chapter2


/**
 * 如何使用基本模式匹配
 * 如何使用模式匹配并将结果返回给调用方
 * 如何通过通配符使用模式匹配
 * 如何在相同条件下对两个或多个项目使用模式匹配
 * 如何在case子句中将模式匹配与if表达式一起使用
 * 如何在类型上使用模式匹配
 */
object SwitchD {
  def main(args: Array[String]): Unit = {
    val a: String = "a"
    a match {
      case "b" => println("b")
      case "c" => println("c")
      case "a" => println("a")
      case _ => println("\nStep")
    }

    var donutType = "Plain Donut"
    println("\nStep 2: Pattern matching and return the result")
    val tasteLevel = donutType match {
      case "Glazed Donut" => "Very tasty"
      case "Plain Donut" => "Tasty"
    }
    println(s"Taste level of $donutType = $tasteLevel")


    println("\nStep 4: Pattern matching with two or more items on the same condition")
    val tasteLevel3 = donutType match {
      case "Glazed Donut" | "Strawberry Donut" => "Very tasty"
      case "Plain Donut" => "Tasty"
      case _ => "Tasty"
    }
    println(s"Taste level of $donutType = $tasteLevel3")


    println("\nStep 5; Pattern matching and using if expressions in the case clause")
    val tasteLevel4 = donutType match {
      case donut if (donut.contains("Glazed") || donut.contains("Strawberry")) => "Very tasty"
      case "Plain Donut" => "Tasty"
      case _ => "Tasty"
    }
    println(s"Taste level of $donutType = $tasteLevel4")


    println("\nStep 6: Pattern matching by types")
    val priceOfDonut: Any = 2.50
    val priceType = priceOfDonut match {
      case price: Int => "Int"
      case price: Double => "Double"
      case price: Float => "Float"
      case price: String => "String"
      case price: Boolean => "Boolean"
      case price: Char => "Char"
      case price: Long => "Long"
    }
    println(s"Donut price type = $priceType")

  }
}
