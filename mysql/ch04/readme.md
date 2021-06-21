#### 4.1 子查询



= ANY
IN

ALL



### 子查询可以解决的经典问题

1. 行号 

```sql
select empid,
    (
        select count(*) from sales T2
        where t2.empid < t1.empid
    ) as rownum
from  
    sales as T1



select emp_no,dept_no,@a:= @a+1 as row_num
from dept_empt, (select @a:=0) t 

```




事务：
ACID 

事务分类：
扁平事务
带有保存点的扁平事务： 切断血缘关系
链事务： 优化 savepoint， 传递给下一个事务
嵌套事务： 树
分布式事务


### 事务操作

1. start Transaction |begin 
2. commit 
3. rollback
4. savepoint idetifier
5. rollback to savepoint idetifier
6. release savepoint
7. set Transaction 



### 事务的隔离级别

1. read uncommit
2. read commit 
3. repeatable read
4. searializable




###  索引

1. 数据结构和算法



在B+树索引中，B+树索引只能找到某条记录所在的页，需要再根据二分查找来进一步找到记录所在的页的具体位置

二叉搜索树
平衡二叉树 AVL
B+树 是为了磁盘或其他直接存取辅助设备设计的一种平衡查找树，在B+树种
