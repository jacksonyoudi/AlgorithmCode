package org.rqq.ch01

import scala.collection.immutable

class Pair[P, Q](val first: P, val second: Q) {
  val twice: Int => Int = (x: Int) => 2 * x
}


object Pair {
  def runTwice(body: => Unit): Unit = {
    body
    body
  }

  def main(args: Array[String]): Unit = {
    runTwice {
      println("hello world")
    }

    for (i <- 0 until 10) {
      println(i)
    }
    (0 until 10).foreach(println)

    val ints: immutable.IndexedSeq[Int] = for (i <- 0 until 10) yield i
    (0 until 10).map(i => -1 * i)
    (0 until 10).map(-1 * _)

    (0 until 4).flatMap(x => (0 until 4).map(y => (x, y)))

    val intToInt: Map[Int, Int] = Map(1 -> 2, 2 -> 3, 3 -> 4)
    intToInt.get(5) match {
      case Some(n) => println(s"success is: ${n}")
      case None => println("not get")
      case _ => print("is empty")
    }

  }
}
