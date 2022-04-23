package org

package object rqq {
  def log(msg: String): Unit = {
    println(s"${Thread.currentThread.getName}:${msg}")
  }
}
