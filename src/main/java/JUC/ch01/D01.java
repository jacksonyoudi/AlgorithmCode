package JUC.ch01;


/**
 * 虚假唤醒
 *
 *
 */
class Share {
    private int number = 0;

    // 判断， 干活 通知
    public synchronized void incr() throws InterruptedException {
        while (number != 0) {
            this.wait(); // 哪里睡，就在哪里唤醒
        }
        number++;

        System.out.println(Thread.currentThread().getName() + number);
        this.notify();
    }

    public synchronized void decr() throws InterruptedException {
        while (number != 1) {
            this.wait();
        }

        number--;
        System.out.println(Thread.currentThread().getName() + number);
        this.notify();

    }

}


public class D01 {
    public static void main(String[] args) {
        Share share = new Share();
        new Thread(new Runnable() {
            @Override
            public void run() {
                try {
                    share.incr();
                } catch (InterruptedException e) {
                    e.printStackTrace();
                }
            }
        }, "a").start();

        new Thread(new Runnable() {
            @Override
            public void run() {
                try {
                    share.decr();
                } catch (InterruptedException e) {
                    e.printStackTrace();
                }
            }
        }, "bb").start();
    }

}
