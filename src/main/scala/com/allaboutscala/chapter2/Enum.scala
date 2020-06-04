package com.allaboutscala.chapter2


/**
 * How to create an Enumeration
 * How to print the String value of an Enumeration
 * How to print the id of an Enumeration
 * How to print all the values listed in an Enumeration
 * How to use pattern matching to find specific Enumeration
 * How to change the ordering of Enumeratio
 */
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

    println("\nStep 4: How to print all the values listed in Enumeration")
    println(s"Donut types = ${Donut.values}")

    println("\nStep 5: How to pattern match on enumeration values")
    Donut.values.foreach {
      case d if (d == Donut.Strawberry || d == Donut.Glazed) => println(s"Found favourite donut = $d")
      case _ => None
    }

    println("\nStep 6: How to change the default ordering of enumeration values")
    object DonutTaste extends Enumeration{
      type DonutTaste = Value

      val Tasty       = Value(0, "Tasty")
      val VeryTasty   = Value(1, "Very Tasty")
      val Ok          = Value(-1, "Ok")
    }

    println(s"Donut taste values = ${DonutTaste.values}")
    println(s"Donut taste of OK id = ${DonutTaste.Ok.id}")


  }
}
