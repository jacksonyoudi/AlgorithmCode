package leetcode.com.topk;

import org.apache.hadoop.yarn.webapp.hamlet.Hamlet;

import java.util.Comparator;
import java.util.HashMap;
import java.util.Map;
import java.util.PriorityQueue;

public class TopkFrequentElements {
    public int[] topKFrequent(int[] nums, int k) {
        HashMap<Integer, Integer> map = new HashMap<Integer, Integer>();

        for (int num : nums) {
            Integer i = map.get(num);
            if (i != null) {
                map.put(num, i + 1);
            } else {
                map.put(num, 1);
            }
        }

        PriorityQueue<Eleme> queue = new PriorityQueue<>(
                new Comparator<Eleme>() {
                    @Override
                    public int compare(Eleme o1, Eleme o2) {
                        return o2.getCnt() - o1.getCnt();
                    }
                }
        );


        for (Map.Entry<Integer, Integer> entry : map.entrySet()) {
            Eleme elem = new Eleme(entry.getKey(), entry.getValue());
            queue.offer(elem);
        }


        int[] result = new int[k];

        for (int i = 0; i < k; i++) {
            int e = queue.poll().getE();
            result[i] = e;
        }


        return result;


    }
}

class Eleme {
    private int e;
    private int cnt;

    public Eleme(int e, int cnt) {
        this.e = e;
        this.cnt = cnt;
    }

    public int getE() {
        return e;
    }

    public void setE(int e) {
        this.e = e;
    }

    public int getCnt() {
        return cnt;
    }

    public void setCnt(int cnt) {
        this.cnt = cnt;
    }
}
