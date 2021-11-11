package com.qqread.ch03

import org.apache.spark.{SparkConf, SparkContext}

object D01 {
  def main(args: Array[String]): Unit = {
    val conf: SparkConf = new SparkConf().setAppName("test").setMaster("local")
//    conf.set()

    val sc: SparkContext = new SparkContext(conf)


  }
}


