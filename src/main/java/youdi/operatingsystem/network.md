### 协议分层


应用层

传输层
    data => 段  segment

网络层
    Internet protocal
     rip，
    ip协议 如何划分子网， 子网掩码
    寻址
    路由(链路) 


数据链路层 data link layer 

物理层



### 多路复用

multiplex:

多个信号，复用一个信道




### 传输层 多路复用

建立一个连接也是需要开销的，所以可以多个请求复用一个TCP连接复用连接一方面可以节省流量，另一方面降低延迟。 
非阻塞

切片


### 网络层多路复用

应用  --- 应用
建立无穷多个传输层  (资源有限)
两个应用之间的线路，设备，需要跨越的网络往往是固定的


网络层： 不断从传输层接收数据



### 多路复用

让多个信号共用一个信道，那么在这个信道上，信息密度就会增加，在密度增加的同时，通过并行发送信号的方式，可以减少阻塞。 





多路复用

提升吞吐量， 数据建间隔





### UDP协议

### 可靠性
    
TCP 支持可靠性的


校验和 checksum
一种非常常见的校验和算法


TCP 实现了 请求， 响应 连接的模型



TCP是一个双工协议，需要数据可以双向传输



TCP三次握手

connect   listen()
syn_send     sync_receved
establisted  estableished

syn 同步
ack 响应
sync_ack  减少握手次数
ack

TCP协议四次挥手

close()

FIN_WAIT1    fin     close_wait

FIN_WAIT_2   Ack      last_Ack

Time_wait    fin      

closed       ack       closed 


ack 和 fin不能合并， 可能还有数据



### 连接

### 封包排序
    
TCP segment 
    滑动窗口， 快速重排


UDP Datagram 
标注了序号

Round Rip 



HTTP协议  TCP 

HTTP3.0  UDP  
quick协议负责可靠性




### I/O模型

从网卡到操作系统
DMA


![XFuuub](https://raw.githubusercontent.com/jacksonyoudi/images/main/uPic/XFuuub.png)

DDOS 保护机制

### socket编程模型

application 
socket

protocol 


socket是一种特殊文件， 每一个都是一个双向的管道， 一端是应用， 一端是缓冲区





### I/O多路复用

在进程内部就需要关注一个数据结构来描述自己会关注哪些socket文件的哪些事件(读，写， 异常)





1. 线性结构
2. 索引结构

select


1024个连接

每次select操作会阻塞当前线程，
在阻塞期间所有的操作系统产生的每个消息，都会通过遍历手段查看是否在3个集合中



Poll()

Epoll()

红黑树



### 同步

程序执行顺序
知道执行顺序




![03M3fm](https://raw.githubusercontent.com/jacksonyoudi/images/main/uPic/03M3fm.png)


