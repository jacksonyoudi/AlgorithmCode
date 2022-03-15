package JUC.ch02.lock;

import java.util.ArrayList;
import java.util.Collections;
import java.util.List;
import java.util.UUID;

public class D05 {
    public static void main(String[] args) {


        List<Object> list = Collections.synchronizedList(new ArrayList<>());

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
