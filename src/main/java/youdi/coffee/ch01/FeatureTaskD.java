package youdi.coffee.ch01;

import java.util.concurrent.ExecutionException;
import java.util.concurrent.FutureTask;

public class FeatureTaskD {
    public static void main(String[] args) throws ExecutionException, InterruptedException {
        FutureTask<Integer> task = new FutureTask<>(
                () -> {
                    System.out.println("hello");
                    return 100;
                }
        );

        new Thread(task, "task").start();

        Integer integer = task.get();
        System.out.println(integer);


    }
}
