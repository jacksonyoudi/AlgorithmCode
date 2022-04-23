package org.rqq.ch02.one

object ThreadUtils {
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
