## sparkcontext

sparkcontext 

sparkEnv
spark UI 
liveListenerBus
SparkStatusTracker 
TaskScheduler
dagscheduler
contextcleaner



sparkEnv:
    序列化
    rpcenv
    blockermanager
    mapoutputTacker


sparkUI:
    监控数据会以sparklistenerEvent形式投递到liveListenerBus中， sparkUI将从
sparkListener中读取数据并显示到web页面

sparkStatusTracker:
    


excutorAllocationManager

spark.dynamicAllocation.enabled


sparkcontext提供常用方法

1. broadcast
2. addsparkLister
3. runjob
4. setCheckpointDir




sparkcontext伴生对象
