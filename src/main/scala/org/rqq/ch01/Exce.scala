package org.rqq.ch01

object Exce {
  def compose[A, B, C](g: B => C, f: A => B): A => C = {
    {
      (a: A) => {
        g(f(a))
      }
    }
  }

  def fuse[A, B](a: Option[A], b: Option[B]): Option[(A, B)] = {
    Option((a.get, b.get))
  }


  def check[T](xs: Seq[T])(pred: T => Boolean): Boolean = {
    {
      for (elem <- xs) {
        if (!pred(elem)) {
          return false
        }
      }
      return true
    }
  }

  def premutations(x: String): Seq[String] = {
    x.toSeq.map(_.toString)
  }


  def combinations(n: Int, xs: Seq[Int]): Iterator[Seq[Int]] = {
    val result: Seq[Seq[Int]] = Seq[Seq[Int]]()
    if (xs.length <= n) {
      result.+:(xs)
      return result.iterator
    }
    return null
  }

  def matcher(regex: String): PartialFunction[String, List[String]] = {
    // 偏函数， map + filter的逻辑
    //    return PartialFunction[String, List[String]] {
    //      case String => List(regex)
    //    }

    return null
  }

  def main(args: Array[String]): Unit = {

  }
}
