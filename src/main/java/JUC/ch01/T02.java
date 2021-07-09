package JUC.ch01;


class Saleser {
    Integer count;


    public Saleser(Integer count) {
        this.count = count;
    }

    public synchronized void sales() {
        if (count > 0) {
            System.out.println(Thread.currentThread().getName() + ": sales");
            count--;
        }
    }

}


public class T02 {
    public static void main(String[] args) {

        Saleser saleser = new Saleser(30);

        new Thread(new Runnable() {
            @Override
            public void run() {
                for (int i = 0; i < 10; i++) {
                    saleser.sales();
                }
            }
        }, "aa").start();
    }
}
