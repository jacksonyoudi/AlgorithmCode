
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





多线程编程步骤 

上部
第一步： 创建资源类， 在资源类创建属性和操作方法
第二步： 创建多个线程， 调用资源类的操作方法

中部

创建资源类 
1. 判断
2. 干活
3. 通知
   
创建多个线程， 调用资源类的操作方法





高内聚 低耦合








## 创建线程的多种方式
1. 继承 thread 类
2. 实现runable接口
3. 使用 callable接口
4. 使用线程池




### Lock接口

可重入锁


synchronized 发生异常，会释放， 
lock会出现死锁的情况



### 线程间通信




### synchronized 

1. 锁的是谁
2. 范围




method的 锁定的 是 this 对象

static method 锁定的是 class 对象 

java中每一个对象都可以作为锁

对于普通方法，锁是当前实例对象
对于静态同步方法，锁是当前类的class对象
对于同步代码块，锁时 synchronized 括号里配置的对象




### 公平锁和非公平锁

```shell
    public ReentrantLock() {
        sync = new NonfairSync();
    }


    public ReentrantLock(boolean fair) {
        sync = fair ? new FairSync() : new NonfairSync();
    }
```


非公平锁：
1. 线程可能会饿死
2. 效率高 


公平锁
1. 阳光普照
2. 效率相对低




### 可重入锁 递归锁
synchronized隐式的，Lock显式的都是可重入的

lock成对出现



### 死锁
1. 什么是死锁
    持有
    互相等待
    互斥
    不能剥夺
    循环等待
   
   
2. 验证死锁

jps 
jstack  jvm自带的 堆栈追踪工具



### 创建多线程方式

1. Tread 
2. Runable 没有返回值 
3. callable 有返回值


1.是否有返回值
2. 是否会抛出异常
3. 实现方法名称不同 call run


### CountDownLatch



### CyclicBarrier


### Segement

### 乐观锁 和 悲观锁

 表锁
 行锁 



读锁： 共享锁，发生死锁
写锁： 独占锁， 发生死锁




### 线程池 

1. 减少开销
2. 提高响应



```shell
    public ThreadPoolExecutor(int corePoolSize,
                              int maximumPoolSize,
                              long keepAliveTime,
                              TimeUnit unit,
                              BlockingQueue<Runnable> workQueue,
                              ThreadFactory threadFactory,
                              RejectedExecutionHandler handler) {
        if (corePoolSize < 0 ||
            maximumPoolSize <= 0 ||
            maximumPoolSize < corePoolSize ||
            keepAliveTime < 0)
            throw new IllegalArgumentException();
        if (workQueue == null || threadFactory == null || handler == null)
            throw new NullPointerException();
        this.acc = System.getSecurityManager() == null ?
                null :
                AccessController.getContext();
        this.corePoolSize = corePoolSize;
        this.maximumPoolSize = maximumPoolSize;
        this.workQueue = workQueue;
        this.keepAliveTime = unit.toNanos(keepAliveTime);
        this.threadFactory = threadFactory;
        this.handler = handler;
    }


corePoolSize: 常驻线程数量
maximumPoolSize： 最大线程数量
keepAliveTime: 存活时间
unit： 单位
workQueue：  阻塞队列
threadFactory： 线程工厂
handler： 拒绝策略

```


#### 拒绝策略
1. 默认策略
直接抛异常
   
2. callerRunsPolicy 调用者运行

3. discardoldestPolicy

4. discardPolicy


















