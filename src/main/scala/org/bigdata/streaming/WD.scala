package org.bigdata.streaming

import org.apache.spark.streaming.StreamingContext
import org.apache.spark.streaming.Seconds
import org.apache.spark.SparkConf
import org.apache.spark.streaming.dstream.{DStream, ReceiverInputDStream}

object WD {
  def main(args: Array[String]): Unit = {
    val conf = new SparkConf()
      .setMaster("local[*]")
      .setAppName("streaming")

    val context = new StreamingContext(conf, Seconds(1))

    val lines: ReceiverInputDStream[String] = context.socketTextStream("localhost", 9999)
    val wc: DStream[(String, Int)] = lines.flatMap(_.split(" ")).map((_, 1)).reduceByKey(_ + _)
    wc.print()
    context.start()

    context.awaitTermination()


  }
}
