package jackson.jianshu

object Var {
  def main(args: Array[String]): Unit = {
    val i: Int = 6
    // i = 7 编译不通过

    var j: Int = 7
    j = 7

    println(i)
    println(j)



  }
}
