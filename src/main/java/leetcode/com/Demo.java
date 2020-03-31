package leetcode.com;

public class Demo {
    public int findJudge(int N, int[][] trust) {
        int[] cnt = new int[N + 1];//统计出入度
        for (int[] index : trust) {
            cnt[index[0]]--;//出度--
            cnt[index[1]]++;//入度++
        }
        for (int i = 1; i <= N; i++) {
            if (cnt[i] == N - 1) return i;
        }
        return -1;
    }

}
