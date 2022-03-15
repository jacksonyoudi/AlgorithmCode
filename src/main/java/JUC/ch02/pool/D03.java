package JUC.ch02.pool;

import java.util.concurrent.ExecutorService;
import java.util.concurrent.Executors;


/**
 *
 */
public class D03 {
    public static void main(String[] args) {
        ExecutorService pool = Executors.newCachedThreadPool();


        try {
            for (int i = 0; i < 10; i++) {
                pool.execute(new Runnable() {
                    @Override
                    public void run() {
                        System.out.println(Thread.currentThread().getName() + " 办理业务");
                    }
                });
            }
        } catch (Exception e) {
            e.printStackTrace();
        } finally {
            pool.shutdownNow();
        }
    }

}
