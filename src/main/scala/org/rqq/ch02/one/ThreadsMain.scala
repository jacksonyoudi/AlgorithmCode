package org.rqq.ch02.one

object ThreadsMain {
  def main(args: Array[String]): Unit = {
    val thread: Thread = Thread.currentThread
    val name: String = thread.getName
    println(s"I am the thread ${name}")
  }
}
