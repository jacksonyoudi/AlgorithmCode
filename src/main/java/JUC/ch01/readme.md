
## 进程

最小的资源分配

## 线程
程序执行的最小单位

线程状态
1. new
2. runable
3. block 
4. waiting 不见不散
5. time_wating 过时不候 
6. terminated 终结 



### wait和sleep的区别 
sleep 是thread的静态方法， wait是object方法， 任何对象都可以调用
sleep不会释放锁，它也不需要占用锁，wait会释放锁，但调用它的前提是当前线程占有锁。
都可以 interruppter 方法中 


### 并发和并行




### 管程 

monitor 监视器

锁 同步机制， 保证同一时间，只能有一个线程访问被保护的数据或者代码

jvm同步基于进入和退出，使用管程对象实现的


### 用户线程和守护线程

自定义线程 
守护线程： gc，特殊线程， 后台



主进程结束了， 用户线程还在运行，jvm存活

没有用户线程， 都是守护线程， jvm会退出。 




# Lock 


## synchronized

