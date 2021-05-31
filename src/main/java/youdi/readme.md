### 对象

* 继承 (接口，父类)
    
    1. have
    2. is 

* 封装
    
    javabean get set
    orm
  
  
* 多态
    
    (method相同)
    继承，方法重写， 父类引用指向子类的
  
对象： 

    method
    attr 

数据和操作放在一起


### jdk  JRE JVM

jdk java开发工具 
jRE 运行环境 
JVM  java虚拟机

![yNVMUK](https://raw.githubusercontent.com/jacksonyoudi/images/main/uPic/yNVMUK.png)


.java -> .class -> jvm 




### == 和 equals

== 对比的是 栈中的值， 基本数据类型是变量值， 引用类型是堆中的内存对象的地址

equals: object中默认也是采用 == 比较 通常会重写




### final作用， 为什么局部内部类和匿名内部类只能访问局部final变量
最终： 
    修饰类： 表示不可被继承
    修饰方法： 不可被子类覆盖(继承)， 可以重载 
    修饰变量： 不能修改， 一旦被赋值


成员变量： 声明中或使用前必须要赋值

内部是复制了一份，必须保证数据不变，所以使用final



### string strngbuffer stringBuiler

string  final 新对象
stringbuffer 线程安全的 synchronized 加了 
stringBuider 线程不安全的 

优先使用 stringBuider
多线程 共享变量 使用 streingbuffer

### 重载和重写 

重载： 发生在同一个类中，方法名必须相同，参数类型个数，顺序不同，返回值和访问修饰符可以不同，发生在编译时， 和返回值无关
重写： 方法名和参数必须一样。 




### 接口和抽象类
1. 抽象类中可以有普通方法， 接口只能是public abstract
2. 抽象类中成员变量可以是各种类型， 接口的变量只能是 public static final 
3. 抽象类只能继承一个， 接口可以实现多个

接口是设计目的： 类的行为的约束
抽象类： 代码复用 不允许实例化 

抽象类 is a: 包含并实现了子类的通用特性，将子类存在的差异化的特性抽象表达，交由子类实现
接口是对行为的抽象， 表达的是 like a 


### List 和 set的区别 
list 进入顺序， index访问 
set 迭代器 

### hashCode与 equals


equals: == 引用地址 

hashcode的作用获取hash码， 确定该对象在hash表中的位置， 堆中地址
hashset  再 equals

重写hashcode同时也重写equals


### arrayList 和LinkedList 区别 
arraylist： 动态数组， 连续存储， 扩容， 随机访问， 插入需要调整顺序， copy数组到新的，使用尾插法
LinkedList:  使用迭代器， index访问， 消耗大




### hashMap和HashTable的区别，底层实现？
1. hashtable 加了 synchronized
2. hbastable 不允许是null
线程安全
   
jdk8 64: 红黑树,到 6了， 又退化成 链表了

数组扩容：
    

### conCurrentHashMap原理，jdk7 jdk8 区别
线程安全

分段锁
jdk7: reentrantLock + segment + HashEntry 一个segment 中包含一个hashEntry数组，每个hashEntry又是一个链表结构
get方法无需加锁  volite 


jdk8： synchronized + CAS + Node + 红黑树 + volite
查找 替换 赋值 使用CAS
锁： 



### IOC容器

    

### 什么是字节码？使用字节码的好处？


### java类加载器

启动类加载器 getBootstrapClassloader
应用加载器 AppClassLoader
扩展类加载器 ExtClassLoader













