package JUC.ch02.itc;

import java.util.concurrent.CyclicBarrier;

public class D02 {
    private static final int number = 7;

    public static void main(String[] args) {
        CyclicBarrier barrier = new CyclicBarrier(number, new Runnable() {
            @Override
            public void run() {
                System.out.println("集齐7颗龙珠，就可以召唤神龙");
            }
        });

        // 集齐7颗龙珠的过程

        for (int i = 0; i < 7; i++) {
            new Thread(() -> {
                System.out.println(Thread.currentThread().getName() + " 星龙珠收集");
                try {
                    barrier.await();

                } catch (Exception e) {
                    e.printStackTrace();
                }
            }, String.valueOf(i)).start();
        }
    }
}
