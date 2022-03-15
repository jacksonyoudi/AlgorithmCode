package leetcode.com.greedy;


import java.util.ArrayList;
import java.util.Comparator;
import java.util.List;

public class RemoveCoveredInterval {

    public int removeCoveredIntervals(int[][] intervals) {
        if (intervals.length == 0) {
            return 0;
        }

        ArrayList<int[]> list = new ArrayList<int[]>();


        for (int[] ints : intervals) {
            list.add(ints);
        }
        list.sort(
                new Comparator<int[]>() {
                    @Override
                    public int compare(int[] o1, int[] o2) {
                        return o1[1] - o2[1];
                    }
                }
        );
        int sum = 0;
        int start = 0;
        int end = list.get(0)[1];
        for (int i = 1; i < list.size(); i++) {
            start = list.get(i)[0];
            if (start < end) {
                sum++;
                continue;
            } else {
                end = list.get(i)[1];
            }
        }

        return sum;
    }


    public static void main(String[] args) {
        System.out.println("hello");
    }
}


class Tuple implements Comparable {
    public int[] data;

    public Tuple(int[] data) {
        this.data = data;
    }


    public int[] getData() {
        return data;
    }

    public void setData(int[] data) {
        this.data = data;
    }

    @Override
    public int compareTo(Object o) {
        Tuple o1 = (Tuple) o;
        return this.getData()[1] - o1.getData()[1];
    }
}