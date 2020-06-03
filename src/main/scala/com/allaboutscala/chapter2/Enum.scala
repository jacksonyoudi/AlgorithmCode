package com.allaboutscala.chapter2

object Enum {
  def main(args: Array[String]): Unit = {
    println("Step 1: How to create an enumeration")
    object Donut extends Enumeration {
      type Donut = Value

      val Glazed = Value("Glazed")
      val Strawberry = Value("Strawberry")
      val Plain = Value("Plain")
      val Vanilla = Value("Vanilla")
    }

    println("\nStep 2: How to print the String value of the enumeration")
    println(s"Vanilla Donut string value = ${Donut.Vanilla}")
    println(s"Vanilla Donut string value = ${Donut.Vanilla.id}")
  }
}
