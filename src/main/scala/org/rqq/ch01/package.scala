package org.rqq

package object ch01 {
  def log(msg: String): Unit = {
    println(s"${Thread.currentThread.getName}:${msg}")
  }
}
