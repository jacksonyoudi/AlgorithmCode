package JUC.ch02.lock.hashset;

import java.util.HashSet;
import java.util.UUID;


/**
 * java.util.ConcurrentModificationException
 *
 */

public class D01 {
    public static void main(String[] args) {
        HashSet list = new HashSet();


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
