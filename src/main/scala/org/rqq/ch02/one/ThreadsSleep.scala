package org.rqq.ch02.one


object ThreadsSleep {
  var uidCount = 0L

  def main(args: Array[String]): Unit = {
    val t: Thread = ThreadUtils.thread {
      Thread.sleep(1000)
      log("new thread running")
      log("completed.")
    }
    t.join()
    log("new thread joined.")
  }

  // synchronized确保了只有在没有其他线程同时执行这个代码块，
  // 或同一个this对象上没有其他同步代码块（synchronized block）被调用时，
  // 该代码块才会被执行。
  // 内置锁 监视器 monitor 其作用是确保同时只有一个线程在该对象上执行某个synchronized代码块
  def getUniqueId(): Long = this.synchronized {
    val freshUid: Long = uidCount + 1L
    uidCount = freshUid
    freshUid
  }

}
