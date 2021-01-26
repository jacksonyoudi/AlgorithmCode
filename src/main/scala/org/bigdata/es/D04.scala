package org.bigdata.es

import org.apache.spark.{SparkConf, SparkContext}
import org.elasticsearch.spark._

object D04 {
  def main(args: Array[String]): Unit = {
    val conf: SparkConf = new SparkConf().setAppName("d01").setMaster("local[*]")
    conf.set("es.index.auto.create", "true")

    val sc: SparkContext = new SparkContext(conf)

    val game = Map(
      "media_type" -> "game",
      "title" -> "FF VI",
      "year" -> "1994")

    val book = Map("media_type" -> "book", "title" -> "Harry Potter", "year" -> "2010")
    val cd = Map("media_type" -> "music", "title" -> "Surfing With The Alien")

    sc.makeRDD(Seq(game, book, cd)).saveToEs("my-collection-{media_type}/doc")

  }
}


