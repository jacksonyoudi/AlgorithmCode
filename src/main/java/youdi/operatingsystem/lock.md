#  锁

## 原子操作
操作不可分

竞争条件
共享资源程序片段： 临界区


解决竞争条件：

1. 互斥
2. ThreadLocal，cas指令， 乐观锁

每个线程一个变量。 

优点： threadlocal
缺点： 很多不能使用

### CAS
Compare And Swap(cas)
作用是更新一个内存地址的值明确要求使用者必须确定知道内存地址中的值是多少

```shell
cas(&oldavalue, expectedValue, newValue)
cas(&i, 0, 1)



while(!cas(&i,i,i+1)) {
  // nothing
}
```

### TAS
Test and Set(tas)
设置一个内存地址的值为1

```shell
task(&lock) {
  return cas(&lock, 0, 1)
}
```


#### 锁

目标就是实现抢占(preempt)
只让给定数量的线程进入临界区，用tas或者cas来实现 



```shell
int lock = 0
enter() {
  while(!cas(&lock, 0, 1)){
    // 什么也不做
  }
}

leave(){
  lock=0;
}


```

### 语言级别的锁实现

相比cas， 锁是一种简单直观的模型
cas更底层，用cas解决问题优化空间更大
用锁解决问题，代码进入临界区之前lock，出去就unlock


自旋锁
```shell
enter(){
  while(!cas(&lock, 0, 1)) {
    //什么也不做
  }
}


```

优点： 不会主动发生context switch，线程切换
缺点： 比较消耗cpu资源


### wait操作

```shell

enter(){
  while(!cas(&lock, 0, 1)) {
    // sleep(1000ms) 主动切换线程，让出cpu
    wait();
  }
}
```
wait方法等待一个信号


生产者消费者模型
wait是一个生产者， 将当前线程挂到一个等待队列上，并休眠
notify是一个消费者，从等待队列中取出一个线程，并重新排队


```shell
enter
leave
wait
notify
```


```java
// Monitor对象
synchronized(obj){
        // enter
        //临界区
} // leave

```

Monitor实现了生产者和消费者模型
* 如果一个线程拿到锁，继续执行
* 如果一个线程竞争锁失败
Monitor调用wait方法触发生产者的逻辑，把线程加入等待集合
  
* 如果一个线程执行完成，Monitor就调用一次notify方法恢复一个等待的线程




### 信号量
控制同时多个进程进入临界区
```shell

int lock=0;

enter(){
  while(lock++>2) {}
}

leave(){
  lock--;
}

```

```shell
up(&lock){
  while(!cas(&lock, lock, lock+1)) {}
}

down(&lock) {
  while(!cas(&lock, lock, lock-1) || lock==0) {
    
  }
}
```

```shell
int lock=2;

down(&lock);
//临界区

up(&lock)
```




###信号量实现生产者和消费者模型

```shell
itn empty = N;
int mutex = 1;
int full = 0;


wait(){
  down(&empty)
  down(&mutex)
  insert()
  up(&mutex)
  up(&full)
}


notify(){
  down(&full)
  down(&mutex)
  remove()
  up(&mutex)
  up(&empty)
}

insert() {
  wait_queue.add(currentThread)
  yield()
}


remove(){
  thread = wait_queue.dequeue()
  thread.resume()
}


```


###  死锁问题

环状的依赖关系

操作顺序相同


### 分布式环境锁

redis:

setnx

原子操作


redis setnx
zk的节点操作等



### 如何控制同一时间只有2个线程运行

1. semaphore
2. 考虑生产者，消费者模型 (协程池)


如果考虑到cpu缓存的存在，会对算法有什么影响










