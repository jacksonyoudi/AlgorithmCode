package JUC.ch02.itc;

import java.util.concurrent.CountDownLatch;

public class D01 {
    public static void main(String[] args) throws InterruptedException {
        CountDownLatch count = new CountDownLatch(6);


        for (int i = 0; i < 6; i++) {

            new Thread(() -> {
                System.out.println(Thread.currentThread().getName() + " 号学生离开教室");

                count.countDown();
            }, String.valueOf(i)).start();

        }

        count.await();
        System.out.println("close down");


    }
}
