package JUC.ch02.lock;


import java.util.concurrent.locks.Condition;
import java.util.concurrent.locks.ReentrantLock;


/**
 * 通过 signal 和 await
 */
class ShareResource {
    private int flag = 1;

    private ReentrantLock lock = new ReentrantLock();

    private Condition c1 = lock.newCondition();
    private Condition c2 = lock.newCondition();
    private Condition c3 = lock.newCondition();


    public void print5(int loop) throws InterruptedException {
        lock.lock();

        try {
            while (flag != 1) {
                c1.await();
            }
            for (int i = 0; i < 5; i++) {
                System.out.println(Thread.currentThread().getName() +  "  次数:" +  i + "轮数:" + loop);
            }
            // 通知 c2
            flag = 2;
            c2.signal();

        } finally {
            lock.unlock();
        }
    }

    public void print10(int loop) throws InterruptedException {
        lock.lock();

        try {
            while (flag != 2) {
                c2.await();
            }
            for (int i = 0; i < 10; i++) {
                System.out.println(Thread.currentThread().getName() +  "  次数:" +  i + "轮数:" + loop);
            }
            // 通知 c2
            flag = 3;
            c3.signal();

        } finally {
            lock.unlock();
        }
    }

    public void print15(int loop) throws InterruptedException {
        lock.lock();

        try {
            while (flag != 3) {
                c3.await();
            }
            for (int i = 0; i < 15; i++) {
                System.out.println(Thread.currentThread().getName() +  "  次数:" +  i + "轮数:" + loop);
            }
            // 通知 c2
            flag = 1;
            c1.signal();

        } finally {
            lock.unlock();
        }
    }



}


public class D02 {
    public static void main(String[] args) {
        ShareResource resource = new ShareResource();

        new Thread(new Runnable(

        ) {
            @Override
            public void run() {
                for (int i = 1; i < 5; i++) {
                    try {
                        resource.print5(i);
                    } catch (InterruptedException e) {
                        e.printStackTrace();
                    }
                }
            }
        }, "aa").start();

        new Thread(new Runnable(

        ) {
            @Override
            public void run() {
                for (int i = 1; i < 5; i++) {
                    try {
                        resource.print10(i);
                    } catch (InterruptedException e) {
                        e.printStackTrace();
                    }
                }
            }
        }, "bb").start();

        new Thread(new Runnable(

        ) {
            @Override
            public void run() {
                for (int i = 1; i < 5; i++) {
                    try {
                        resource.print15(i);
                    } catch (InterruptedException e) {
                        e.printStackTrace();
                    }
                }
            }
        }, "cc").start();
    }
}
