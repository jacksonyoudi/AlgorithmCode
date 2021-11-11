package JUC.ch02.lock.hashmap;

import java.util.HashMap;
import java.util.UUID;
import java.util.concurrent.ConcurrentHashMap;


/**
 * ConcurrentModificationException
 */

public class D02 {
    public static void main(String[] args) {
        ConcurrentHashMap list = new ConcurrentHashMap();


        for (int i = 0; i < 30; i++) {
            new Thread(new Runnable() {
                @Override
                public void run() {
                    list.put(UUID.randomUUID().toString().substring(0, 8), UUID.randomUUID().toString().substring(0, 8));
                    System.out.println(list);
                }
            }, String.valueOf(i)).start();
        }
    }
}
