package leetcode.com;

import java.util.Stack;

public class StackDemo {
    public static void main(String[] args) {
        Stack<Integer> stack = new Stack<Integer>();
        stack.push(1);
        Integer pop = stack.pop();
        System.out.println(pop);
        stack.push(10);
        stack.push(100);

        System.out.println(stack.peek());

    }
}
