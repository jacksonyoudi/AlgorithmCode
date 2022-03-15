package JUC.qqread;

import java.util.concurrent.RecursiveTask;

/**
 * @program: AlgorithmCode
 * @description:
 * @author: changyouliang
 * @date: 2022/02/21
 **/
public class AccumulateTask extends RecursiveTask<Integer> {
    private static final int THRESHOLD = 2;

    private int start;
    private int end;

    public AccumulateTask(int start, int end) {
        this.start = start;
        this.end = end;
    }

    @Override
    protected Integer compute() {
        int sum = 0;
        boolean b = (end - start) <= THRESHOLD;
        if (b) {
            for (int i = start; i <= end; i++) {
                sum += i;
            }
            System.out.println("执行任务，计算" + start + " 到 " + end + " 的和，结果是：" + sum);
        } else {
            System.out.println("切割任务： 将 " + start + " 到 " + end + " 的和一分为二");
            int mid = (start + end) / 2;
            AccumulateTask preTask = new AccumulateTask(start, mid);
            AccumulateTask endTask = new AccumulateTask(mid + 1, end);
            preTask.fork();
            endTask.fork();

            Integer lr = preTask.join();
            Integer rr = endTask.join();
            sum = lr + rr;
        }
        return sum;
    }
}
