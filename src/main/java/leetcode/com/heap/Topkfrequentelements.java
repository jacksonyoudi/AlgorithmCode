package leetcode.com.heap;

import org.apache.hadoop.yarn.webapp.hamlet.Hamlet;

import java.util.Comparator;
import java.util.HashMap;
import java.util.Map;
import java.util.PriorityQueue;


public class Topkfrequentelements {
    public int[] topKFrequent(int[] nums, int k) {
        Map<Integer, Integer> map = new HashMap<Integer, Integer>();
        for (int num : nums) {
            Integer res = map.get(num);
            if (res == null) {
                map.put(num, 1);
            } else {
                map.put(num, (int) res + 1);
            }
        }
        PriorityQueue<Eleme> priorityQueue = new PriorityQueue<>(new Comparator<Eleme>() {
            @Override
            public int compare(Eleme o1, Eleme o2) {
                return o2.getCnt() - o1.getCnt();
            }
        });

        for (Map.Entry<Integer, Integer> entry : map.entrySet()) {
            Eleme eleme = new Eleme(entry.getKey(), entry.getValue());
            priorityQueue.offer(eleme);
        }

        int[] res = new int[k];

        for (int i = 0; i < k; i++) {
            res[i] = priorityQueue.poll().getEleme();
        }
        return res;


    }
}


class Eleme {
    private int eleme;
    private int cnt;

    public Eleme(int eleme, int cnt) {
        this.eleme = eleme;
        this.cnt = cnt;
    }

    public int getEleme() {
        return eleme;
    }

    public void setEleme(int eleme) {
        this.eleme = eleme;
    }

    public int getCnt() {
        return cnt;
    }

    public void setCnt(int cnt) {
        this.cnt = cnt;
    }
}