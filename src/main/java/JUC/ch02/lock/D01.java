package JUC.ch02.lock;


import java.util.concurrent.locks.Condition;
import java.util.concurrent.locks.Lock;
import java.util.concurrent.locks.ReentrantLock;

class Share {
    private int number = 0;
    // 创建condition

    ReentrantLock lock = new ReentrantLock();

    Condition con = lock.newCondition();


    public void incr() throws InterruptedException {
        lock.lock();
        try {
            // 判断 干活  通知
            while (number != 1) {
                con.await();
            }
            number++;
            System.out.println(Thread.currentThread().getName() + number);

            con.signalAll();


        } finally {
            lock.unlock();
        }

    }


    public void decr() throws InterruptedException {
        lock.lock();
        try {
            while (number != 0) {
                con.await();
            }

            number--;
            System.out.println(Thread.currentThread().getName() + number);
            con.signalAll();
        } finally {
            lock.unlock();
        }
    }


}


public class D01 {
    public static void main(String[] args) {
        
    }
}
