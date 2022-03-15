package JUC.ch01;

import java.util.ArrayList;

public class D06 {
    byte[] bytes = new byte[1 * 1024 * 1024];


    public static void main(String[] args) {


        ArrayList<D06> list = new ArrayList<D06>();

        int count = 0;

        try {
            while (true) {

                list.add(new D06());
                count++;
            }
        } catch (Exception e) {
            e.printStackTrace();
        }


    }
}
