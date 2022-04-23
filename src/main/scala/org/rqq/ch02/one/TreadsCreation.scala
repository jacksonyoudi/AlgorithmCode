package org.rqq.ch02.one

object TreadsCreation {
  def main(args: Array[String]): Unit = {
    val thread: Thread = new Thread() {
      override def run() = {
        println("new thread running.")
      }
    }
    thread.start()
    thread.join()
    println("new Thread joined.")

  }

  def thread(body: => Unit): Thread = {
    val thread1: Thread = new Thread {
      override def run(): Unit = {
        body
      }
    }
    thread1.start()
    thread1
  }
}
