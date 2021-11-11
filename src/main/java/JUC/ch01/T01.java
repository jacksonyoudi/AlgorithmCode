package JUC.ch01;

public class T01 {
    public static void main(String[] args) {
        Thread thread = new Thread(new Runnable() {
            @Override
            public void run() {
                System.out.println(
                        Thread.currentThread().getName() + "::" + Thread.currentThread().isDaemon()
                );
                while (true) {

                }
            }
        }, "aa");

        thread.setDaemon(true);

        thread.start();



        System.out.println(Thread.currentThread().getName() + "over");
    }
}
