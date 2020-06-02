package com.allaboutscala.chapter2

object TypeInterface {
  def main(args: Array[String]): Unit = {
    println("Step 1: Immutable variable")
    val donutsToBuy: Int = 5

    println("\nStep 2: Scala Types")
    val donutsBoughtToday = 5
    val bigNumberOfDonuts = 100000000L
    val smallNumberOfDonuts = 1
    val priceOfDonut = 2.50
    val donutPrice = 2.50f
    val donutStoreName = "Allaboutscala Donut Store"
    val donutByte = 0xa
    val donutFirstLetter = 'D'
    val nothing = ()

    val a: String = donutPrice.toString
  }
}
