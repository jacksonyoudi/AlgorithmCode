package org.rqq.ch03

import org.rqq.log

import java.util.concurrent.atomic.AtomicBoolean
import scala.language.postfixOps

object AtomicLock {
  private val lock = new AtomicBoolean(false)

  def mySynchronized(body: => Unit): Unit = {
    while (!lock.compareAndSet(false, true)) {
    }
    try {
      body
    } finally {
      lock.set(false)
    }
  }

  def main(args: Array[String]): Unit = {
    var count: Int = 0

    for (i <- 0 until 10) {
      execute {
        mySynchronized {
          count += 1
        }
      }
    }
    Thread.sleep(1000)
    log(s"count is ${count}")


  }
}
