package com.allaboutscala.chapter3

object CreateFunctionWithOptionReturnType {
  def main(args: Array[String]): Unit = {

  }

  println(s"Step 1: Define a function which returns an Option of type String")

  def dailyCouponCode(): Option[String] = {
    // look up in database if we will provide our customers with a coupon today
    val couponFromDb = "COUPON_1234"
    Option(couponFromDb).filter(_.nonEmpty)
  }
}
