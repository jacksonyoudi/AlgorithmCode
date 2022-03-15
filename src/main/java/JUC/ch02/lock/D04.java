package JUC.ch02.lock;

import java.util.ArrayList;
import java.util.UUID;
import java.util.Vector;

public class D04 {
    public static void main(String[] args) {
        Vector list = new Vector();


        for (int i = 0; i < 10; i++) {
            new Thread(new Runnable() {
                @Override
                public void run() {
                    list.add(UUID.randomUUID().toString().substring(0, 8));
                    System.out.println(list);
                }
            }, String.valueOf(i)).start();
        }


    }
}
