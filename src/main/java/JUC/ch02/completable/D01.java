package JUC.ch02.completable;

import java.util.concurrent.CompletableFuture;
import java.util.concurrent.ExecutionException;


/**
 *  异步一般使用MQ
 */
public class D01 {
    public static void main(String[] args) throws ExecutionException, InterruptedException {
        CompletableFuture<Void> runAsync = CompletableFuture.runAsync(() -> {
            System.out.println(Thread.currentThread().getName() + " : CompletableFuture");
        });

        runAsync.get();


        // 异步调用
        CompletableFuture<Integer> async = CompletableFuture.supplyAsync(() -> {
            System.out.println(Thread.currentThread().getName() + " : CompletableFuture");
            return 1024;
        });

        async.whenComplete((t, u) -> {
            System.out.println("---t, " + t);
            System.out.println("---u, " + u);
        }).get();

    }
}
