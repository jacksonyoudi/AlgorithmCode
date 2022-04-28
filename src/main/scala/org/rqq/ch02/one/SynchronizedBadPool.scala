package org.rqq.ch02.one

import scala.collection.mutable

object SynchronizedBadPool {
  private val tasks: mutable.Queue[() => Unit] = mutable.Queue[() => Unit]()

  def main(args: Array[String]): Unit = {

  }
}




