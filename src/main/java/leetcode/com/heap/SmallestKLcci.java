package leetcode.com.heap;


import java.util.Comparator;
import java.util.PriorityQueue;

public class SmallestKLcci {
    public int[] smallestK(int[] arr, int k) {
        int[] res = new int[k];
        if (k == 0) {
            return res;
        }

        PriorityQueue<Integer> priorityQueue = new PriorityQueue<Integer>(new Comparator<Integer>() {

            @Override
            public int compare(Integer o1, Integer o2) {
                return o1 - o2;
            }
        });


        for (int i : arr) {
            priorityQueue.offer(i);
        }

        for (int j = 0; j < k; j++) {
            res[j] = (int) priorityQueue.poll();
        }
        return res;
    }
}
