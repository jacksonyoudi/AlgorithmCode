### 如何更好利用资源

计算： 利用CPU处理算数运算
计算密集型


IO密集型
IO本质对设备的读写
1. 读取键盘， 文件， 磁盘



### 读取磁盘数据到内存中的这个过程，cpu不需要一个个字节处理
DMA模块 直接内存访问


![lK1M6q](https://raw.githubusercontent.com/jacksonyoudi/images/main/uPic/lK1M6q.png)

不能停止工作， 执行空闲指令

### top：
us: 用户
sy: 系统
nice:  优先级
idle: 
wait:
hi: 中断
si： 软中毒

load average: 



1. 如果平均负载很高，cpu的IO wait也很高，说明cpu因为需要大量等待IO无法处理完成工作

原因： 线上服务器打印日志太频繁了，读写数据库，网络太频繁了。
考虑使用批量 写入更快，  缓存 缓冲


### 通信量 traffic

复用发起写操作程序的连接和缓冲区


iotop
df

/proc/diskstats



### 决定进程、线程数量

是否  进程数、线程数据 = cpu核心

node.js是事件驱动

java是一个多线程的模型，线程和内核线程对应比例 1:1
Go有轻量级级线程，多个轻量级线程复用一个内核级线程


IO-wait cpu等待io完成

1 - p^n


进程： IO 延迟读写，优化减少读写，批量
线程： cpu







cpu密集： 等于core
IO密集型： 大于core

压测 + 监控






