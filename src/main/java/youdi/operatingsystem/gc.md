### 为什么产生内存泄露

因为应用的内存管理一直处于一个和应用程序执行并发的状态

系统性能下降， 一种突然的崩溃
因为一旦内存被占用，系统性能就开始雪崩式下降


### 什么gc

应用程序
GC 应用的实际内存管理
OS 虚拟内存

Garbage collection 


### GC功能
1. 操作系统进程交互，负责申请内存
2. 应用会向GC申请内存
3. 承担垃圾回收能力，标记不用的对象并回收
4. 针对应用特性进行动态的优化


GC：
* 吞吐量
* 足迹
* 暂停时间 stop work


### GC目标
提高吞吐量，如利用多核优势就是一种最直观的方法


阿姆达定律
s= 1/(1-p)

GC能够支持并行处理
不能用用太长的暂停时间
    
    如果GC导致应用太久，是灾难性的
    暂停时间很短， GC和应用交替频繁


三者之间的关系在类似相同的成本代价下不可兼得的关系
编程语言会提供参数让你选择根据自己的应用特性决定GC行为




### 引用计算算法 reference Counter

循环引用
在多线程环境下引用计数算法一旦算错1次，处理竞争关系比较耗费性能
内存碎片 (整理)


引用计数法出错概率大
容错能力差


### Root Tracing 算法

从引用路径上，如果一个对象的引用链中包括一个根对象(Root Object)，那么这个对象就是活动的

所有引用关系的源头

![vc1nuq](https://raw.githubusercontent.com/jacksonyoudi/images/main/uPic/vc1nuq.png)


root tracking 的容错性很好
GC通过不断的执行Root Tracking算法找到需要回收的元素



### 标记清除算法(mark sweep) 算法
白色： 一种不确定的状态： 可能被回收
黑色： 一种确定的状态： 不会被回收


算法实现过程中， 假设有两个全局变量是已知的：

* heapset中拥有所有的对象
* rootSet中拥有所有的RootObject

```shell
for obj in heapSet {
  obj.color = white
}



func mark(obj) {
  if obj.color == white {
    obj.color = black
    for v in references(obj) {
      mark(v)
}
}}


for root in rootSet {
  mark(root)
}

// 遍历整个heapSet找到白色的对象进行回收

for obj in heapSet {
  if obj.color == white {
    free(obj)
  }
}
```



算法有两个阶段
1. 标记阶段
2. 清除阶段

浮动垃圾 float Garbage

希望GC能够增量回收



### 如何解决内存的循环引用问题？
root tracking算法
    



### 标记清除算法(mark sweep) 算法
标记程序  找到不用的内存
清除程序 sweep  回收不用的资源
变更程序    用户程序对内存进行了修改 mutation 


不允许并行执行

read  on write 一遍读

合并版本
stw



### 三色标记-清除算法(Tri-Color Mark Sweep)
白色： 需要GC的对象
黑色： 确定不需要GC的对象
灰色： 可能不需要GC的对象(增量任务)




初始化完成后， 会启动标记任务，在标记程序中， 可以暂停标记程序执行mutation





![wqDRci](https://raw.githubusercontent.com/jacksonyoudi/images/main/uPic/wqDRci.png)

