package JUC.ch02.itc;

import java.util.concurrent.locks.ReentrantLock;

public class WaitGroup {
    private int number = 6;

    public static void main(String[] args) {
        ReentrantLock lock = new ReentrantLock();

        WaitGroup group = new WaitGroup();

        for (int i = 0; i < 7; i++) {
            new Thread(()-> {
                lock.lock();

                try {
                    if (group.number <= 0) {
                        System.out.println("close door");
                    } else {
                        System.out.println("num--");
                        group.number--;
                    }



                }

                finally {
                    lock.unlock();
                }


            }, String.valueOf(i)).start();
        }



    }
}
