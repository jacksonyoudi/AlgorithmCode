package org.rqq.ch01

class Position(val x: Int, val y: Int) {
  def +(that: Position): Position = {
    new Position(this.x + that.x, this.y + that.y)
  }
}

object Position {
  def apply(x: Int, y: Int): Position = {
    return new Position(x, y)
  }
}
