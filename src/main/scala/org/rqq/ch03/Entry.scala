package org.rqq.ch03

import java.util.concurrent.atomic.AtomicReference
import scala.annotation.tailrec
import org.rqq.log

sealed trait State

class Idle extends State

class Creating extends State

class Copy(val n: Int) extends State

class Deleting extends State


class Entry(val isDir: Boolean) {
  private val state: AtomicReference[State] = new AtomicReference[State](new Idle)


  @tailrec private def prepareForDeleting(entry: Entry): Boolean = {
    val s: State = entry.state.get
    s match {
      case i: Idle =>
        if (entry.state.compareAndSet(s, new Deleting)) {
          true
        } else {
          prepareForDeleting(entry)
        }

      case c: Creating =>
        log("file currently created, cannot delete.")
        false

      case c: Copy =>
        log("copy currently created, cannot delete")
        false
      case d: Deleting =>
        false
    }
  }


  def releaseCopy(e: Entry): Copy = {
    e.state.get match {
      case c: Copy => {
        val s: State = if (c.n == 1) new Idle else new Copy(c.n - 1)
        if (e.state.compareAndSet(c, s)) {
          c
        } else {
          releaseCopy(e)
        }
      }
    }
  }
}


