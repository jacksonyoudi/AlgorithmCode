package JUC.ch02.callable;


import java.util.concurrent.Callable;
import java.util.concurrent.ExecutionException;
import java.util.concurrent.FutureTask;


/**
 * 1. FutureTask
 */
class M1 implements Runnable {
    @Override
    public void run() {
        System.out.println("M1");
    }
}


class M2 implements Callable {
    @Override
    public Object call() throws Exception {
        return 200;
    }
}

public class D01 {
    public static void main(String[] args) throws ExecutionException, InterruptedException {

        Thread t1 = new Thread(new M1(), "aa");
        t1.start();

        FutureTask task = new FutureTask<>(new M2());


        FutureTask<Integer> task1 = new FutureTask<>(new Callable<Integer>() {
            @Override
            public Integer call() throws Exception {
                System.out.println(Thread.currentThread().getName() + "come in callable");
                Thread.sleep(1000);
                return 200;
            }
        });

        new Thread(task1, "lucy").start();

        while (!task1.isDone()) {
            System.out.println("waiting....");
        }

        System.out.println(task1.get());
        System.out.println(task1.get());


        // futuretask

    }
}
