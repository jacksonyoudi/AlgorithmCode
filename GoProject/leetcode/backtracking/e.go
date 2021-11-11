package main

import "fmt"

type Box struct {
	l []int
}

func (b *Box) push(i int) {
	b.l = append(b.l, i)
}

func (b *Box) pop() {
	if len(b.l) > 0 {
		b.l = b.l[1:]
	}
}

func (b *Box) print() {
	fmt.Printf("{")
	for _, v := range b.l {
		fmt.Printf("%d,", v)
	}
	fmt.Printf("}\n")
}

func solve3(nums []int, index int, b *Box) {
	ll := len(nums)
	b.print()  // 说出箱子的状态
	if index >= ll {
		return
	}
	b.push(nums[index]) // 将自己的选中的宝石放进去
	solve3(nums, index+1, b) // 交给下一个人
	b.pop()  // 把自己的宝石取出来

}

func print_backtracking(nums []int, index int) {
	ll := len(nums)
	if ll == 0 && index == 0 {
		fmt.Printf("{}\n")
		return
	}

	if index > ll {
		return
	}

	fmt.Printf("{")
	for i := 1; i < index+1; i++ {
		fmt.Printf("%d,", nums[i-1])
	}
	fmt.Printf("}\n")

	print_backtracking(nums, index+1)

}


//void backtrace(int[] A,
//int i, /*第i个人*/
//Box s, /*箱子*/
//answer/*存放所有的答案*/) {
//final int N = A == null ? 0 : A.length;
//if (状态满足要求) {
//answer.add(s);
//}
//
//if ([i, ...., 后面）的人都没有任何选项了) {
//return;
//}
//for 宝石 in {第i个人当前所有宝石选项} {
//s.push(宝石);
//backtrace(A, i + 1, s, answer);
//s.pop();
//}
//}





func main() {
	b := &Box{l: []int{}}
	solve3([]int{1, 2, 3, 4}, 0, b)
}
