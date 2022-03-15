package leetcode.com.heap;

import java.util.Comparator;
import java.util.PriorityQueue;

public class Kthlargestelementinanarray {
    public int findKthLargest(int[] nums, int k) {
        PriorityQueue<Integer> priorityQueue = new PriorityQueue<>(new Comparator<Integer>() {
            @Override
            public int compare(Integer o1, Integer o2) {
                return o2-o1;
            }
        });

        for (int num : nums) {
            priorityQueue.offer(num);
        }

        int res = 0;
        for (int i = 0; i < k; i++) {
            res = (int) priorityQueue.poll();
        }
        return res;
    }
}
