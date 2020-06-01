package com.allaboutscala.chapter2

object Escape {
  def main(args: Array[String]): Unit = {
    println("\nStep 2: Using backslash to escpae quotes")
    val donutJson2: String = "{\"donut_name\":\"Glazed Donut\",\"taste_level\":\"Very Tasty\",\"price\":2.50}"
    println(s"donutJson2 = $donutJson2")

    println("\nStep 3: Using triple quotes \"\"\" to escape characters")
    val donutJson3: String = """{"donut_name":"Glazed Donut","taste_level":"Very Tasty","price":2.50}"""
    println(s"donutJson3 = $donutJson3")


    val donutJson4: String =
      """
        |{
        |"donut_name":"Glazed Donut",
        |"taste_level":"Very Tasty",
        |"price":2.50
        |}
      """
        .stripMargin('|')

    println(donutJson4)

  }
}
