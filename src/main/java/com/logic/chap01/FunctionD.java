package com.logic.chap01;

public class FunctionD {
    public static void main(String[] args) {

    }

    public static int sum(int a, int b) {
        return a + b;
    }

    public static int max(int min, int... b) {
        int max = min;

        for (int i = 0; i < b.length; i++) {
            if (max < b[i]) {
                max = b[i];
            }
        }
        return max;
    }
}


