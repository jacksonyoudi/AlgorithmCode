package com.allaboutscala.chapter2


/**
 * How to use Option and Some
 * How to use Option and None
 * How to use Pattern Matching with Option
 */
object OptionSomeNone {
  def main(args: Array[String]): Unit = {
    println("Step 1: How to use Option and Some - a basic example")
    val glazedDonutTaste: Option[String] = Some("Very Tasty")
    println(s"Glazed Donut taste = ${glazedDonutTaste.get}")

    println("\nStep 2: How to use Option and None - a basic example")
    val glazedDonutName: Option[String] = None
    println(s"Glazed Donut name = ${glazedDonutName.getOrElse("Glazed Donut")}")

    println("\nStep 3: How to use Pattern Matching with Option")
    glazedDonutName match {
      case Some(name) => println(s"Received donut name = $name")
      case None => println(s"No donut name was found!")
    }

    println("\nTip 1: Filter None using map function")
    glazedDonutTaste.map(taste => println(s"glazedDonutTaste = $taste"))
    glazedDonutName.map(name => println(s"glazedDonutName = $name"))
  }
}
