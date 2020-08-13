package com.logic.chap01;

public class ArrayDemo {
    public static void main(String[] args) {
        int arr[] = {1, 2, 3};
        int[] arr1 = new int[]{1, 2, 3};
        int[] arr2 = new int[3];

        arr2[0] = 1;
        arr2[1] = 2;
        arr2[2] = 3;


        // 每个元素都有一个默认值，这个默认值跟数组类型有关

        int[] arrA = {1, 2, 3};
        int[] arrB = {1, 2, 3, 4};
        arrA = arrB;
        System.out.println(arrA.toString());

    }
}
