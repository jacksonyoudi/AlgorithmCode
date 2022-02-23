package JUC.qqread;

import java.util.concurrent.*;

/**
 * @program: AlgorithmCode
 * @description:
 * @author: changyouliang
 * @date: 2022/02/21
 **/
public class ForkJoinTest {
    public static void main(String[] args) throws ExecutionException, InterruptedException, TimeoutException {
        ForkJoinPool pool = new ForkJoinPool();

        AccumulateTask task = new AccumulateTask(1, 100);
        ForkJoinTask<Integer> submit = pool.submit(task);
        Integer sum = submit.get(1, TimeUnit.SECONDS);
        System.out.println(sum);
    }
}