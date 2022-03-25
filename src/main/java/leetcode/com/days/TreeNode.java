package leetcode.com.days;

/**
 * @program: AlgorithmCode
 * @description:
 * @author: changyouliang
 * @date: 2022/03/25
 **/
public class TreeNode {
    int val;
    TreeNode left;
    TreeNode right;

    TreeNode() {
    }

    TreeNode(int val) {
        this.val = val;
    }

    TreeNode(int val, TreeNode left, TreeNode right) {
        this.val = val;
        this.left = left;
        this.right = right;
    }
}
