package org.rqq

package object ch03 {
  def execute(body: => Unit): Unit = {
    val t: Thread = new Thread(new Runnable {
      override def run() = {
        body
      }
    })

    t.start()

  }
}
