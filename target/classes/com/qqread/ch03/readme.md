# spark基础设施

## 3.1 spark配置

配置： 

sparkconf
```concurrentHashMap[string, string]```

getSystemProperties
spark.

conf api:

set() 

```scala
  def setMaster(master: String): SparkConf = {
    set("spark.master", master)
  }

  /** Set a name for your application. Shown in the Spark web UI. */
  def setAppName(name: String): SparkConf = {
    set("spark.app.name", name)
  }
```

sparkconf 支持clone 


```scala
  override def clone: SparkConf = {
    val cloned = new SparkConf(false)
    settings.entrySet().asScala.foreach { e =>
      cloned.set(e.getKey(), e.getValue(), true)
    }
    cloned
  }
```


## 3.2 spark内置RPC框架

transportContext 传输上下文
transportconf 
rpchandler
messageEncoder
messageDecoder
transportFrameDecoder
RpcResponserCallBack



