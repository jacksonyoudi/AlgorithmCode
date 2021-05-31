package leetcode.com.linkedlist;

// 实现单链表

import java.util.List;

class MyLinkedList {

    // 链表结点的定义

    class ListNode {

        // val用来存放链表中的数据

        public int val = 0;

        // next指向下一个结点

        public ListNode next = null;

        public ListNode() {
        }

        public ListNode(int x) {

            val = x;

        }

    }

    /**
     * code here: 初始化链表
     */
    private ListNode dummy = new ListNode();
    private ListNode tail = dummy;
    private int length = 0;


    public MyLinkedList() {

    }

    public void addAtTail(int val) {

        /* code here: 将值为 val 的结点追加到链表尾部*/
        tail.next = new ListNode(val);
        tail = tail.next;
        length++;

    }

    public void addAtHead(int val) {

        /* code here: 插入值val的新结点，使它成为链表的第一个结点*/
        ListNode node = new ListNode(val);
        node.next = dummy.next;
        dummy.next = node;
        // 注意动静结合原则，添加结点时，注意修改tail指针 为空的情况
        if (tail == dummy) {
            tail = node;
        }

        length++;


    }

    public int get(int index) {

        /* code here: 获取链表中第index个结点的值。如果索引无效，则返回-1。*/

        // index从0开始。
        if (index < 0 || index > length) {
            return -1;
        }

        return getPrevNode(index).next.val;

    }

    public void addAtIndex(int index, int val) {

        // code here:

        // 在链表中的第 index 个结点之前添加值为 val  的结点。

        // 1. 如果 index 等于链表的长度，则该结点将附加到链表的末尾。

        // 2. 如果 index 大于链表长度，则不会插入结点。

        // 3. 如果index小于0，则在头

        if (index > length) {
            return;
        } else if (index == length) {
            addAtTail(val);
        } else if (index <= 0) {
            addAtHead(val);
        } else {
            ListNode prevNode = getPrevNode(index);
            ListNode node = new ListNode(val);
            node.next = prevNode.next;
            prevNode.next = node;
            length++;

        }


    }

    public void deleteAtIndex(int index) {

        /* code here: 如果索引index有效，则删除链表中的第index个结点。*/

        if (index < 0 || index > length) {
            return;
        }
        ListNode prevNode = getPrevNode(index);
        if (tail == prevNode.next) {
            tail = prevNode;
            tail.next = null;
        }

        prevNode.next = prevNode.next.next;
        length--;


    }

    private ListNode getPrevNode(int index) {

        /*返回index结点的前驱结点，如果index不存在，那么返回dummy*/

        // 初始化front与back，分别一前一后

        ListNode front = dummy.next;

        ListNode back = dummy;

        // 在查找的时候，front与back总是一起走

        for (int i = 0; i < index && front != null; i++) {

            back = front;

            front = front.next;

        }

        // 把back做为prev并且返回

        return back;

    }


}
