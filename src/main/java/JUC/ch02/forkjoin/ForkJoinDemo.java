package JUC.ch02.forkjoin;


import java.util.concurrent.ExecutionException;
import java.util.concurrent.ForkJoinPool;
import java.util.concurrent.ForkJoinTask;
import java.util.concurrent.RecursiveTask;

class MyTask extends RecursiveTask<Integer> {
    private static final Integer Value = 10;

    private int begin;
    private int end;
    private int result;

    @Override
    protected Integer compute() {
        if ((end - begin) <= Value) {
            for (int i = begin; i <= end; i++) {
                result += i;
            }
        } else {
            int mid = (end + begin) / 2;
            MyTask m1 = new MyTask(begin, mid);
            MyTask m2 = new MyTask(mid, end);

            m1.fork();
            m2.fork();

            // 合并

            result = m1.join() + m2.join();

        }

        return result;

    }


    public MyTask(int begin, int end) {
        this.begin = begin;
        this.end = end;
    }
}

public class ForkJoinDemo {
    public static void main(String[] args) throws ExecutionException, InterruptedException {
        MyTask task = new MyTask(1, 100);
        ForkJoinPool pool = new ForkJoinPool();
        ForkJoinTask<Integer> submit = pool.submit(task);
        Integer integer = submit.get();

        System.out.println(integer);
        pool.shutdownNow();


    }
}
