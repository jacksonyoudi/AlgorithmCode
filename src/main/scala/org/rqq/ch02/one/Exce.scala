package org.rqq.ch02.one

object Exce {
  def main(args: Array[String]): Unit = {

  }

  def parallel[A, B](a: => A, b: => B): (A, B) = {
    val t1: Thread = new Thread() {
      override def run() = {
        a
      }
    }
    t1.start()
    t1.join()

  }

}
