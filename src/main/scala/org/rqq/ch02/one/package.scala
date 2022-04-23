package org.rqq.ch02

package object one {
  def log(msg: String): Unit = {
    println(s"${Thread.currentThread.getName}:${msg}")
  }
}
