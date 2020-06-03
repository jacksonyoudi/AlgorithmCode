package com.allaboutscala.chapter2

/**
 *
 * 如何使用Tuple2类创建具有两个数据点的Tuple
 * 如何访问元组中的每个单独元素
 * 如何创建和使用具有两个以上数据点的元组
 * 如何在元组列表上使用模式匹配
 * 如何使用（）括号轻松创建元组
 */
object TupleD {
  def main(args: Array[String]): Unit = {
    println("Step 1: Using Tuple2 to store 2 data points")
    val tuple: (String, String) = Tuple2("Glazed Donut", "Very Tasty")
    println(s"Glazed Donut with 2 data points = $tuple")

    println("\nStep 2: Access each element in tuple")
    val donutType = tuple._1
    val donutTasteLevel = tuple._2
    println(s"$donutType taste level is $donutTasteLevel")


    println("\nStep 3: Using TupleN classes to store more than 2 data points")
    val tuple1: (String, String, Double) = Tuple3("Glazed Donut", "Very Tasty", 2.50)
    val glazedDonut = tuple1
    println(s"${glazedDonut._1} taste level is ${glazedDonut._2} and it's price is ${glazedDonut._3}")

    println("\nStep 4: How to use pattern matching on Tuples")
    val strawberryDonut = Tuple3("Strawberry Donut", "Very Tasty", 2.50)
    val plainDonut = Tuple3("Plain Donut", "Tasty", 2)

    val donutList = List(glazedDonut, strawberryDonut, plainDonut)

    val priceOfPlainDonut = donutList.foreach { tuple => {
      tuple match {
        case ("Plain Donut", taste, price) => println(s"Donut type = Plain Donut, price = $price")
        case d if d._1 == "Glazed Donut" => println(s"Donut type = ${d._1}, price = ${d._3}")
        case _ => None
      }
    }
    }


    println("\nStep 5: Shortcut for creating Tuples")
    val chocolateDonut = ("Chocolate Donut", "Very Tasty", 3.0)
    println(s"Chocolate donut taste level = ${chocolateDonut._2}, price = ${chocolateDonut._3}")

  }
}
