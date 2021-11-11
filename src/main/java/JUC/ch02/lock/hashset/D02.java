package JUC.ch02.lock.hashset;

import java.util.HashSet;
import java.util.UUID;
import java.util.concurrent.CopyOnWriteArraySet;


/**
 * java.util.ConcurrentModificationException
 * 
 */

public class D02 {
    public static void main(String[] args) {
        CopyOnWriteArraySet list = new CopyOnWriteArraySet<>();


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
