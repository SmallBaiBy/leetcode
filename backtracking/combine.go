package backtracking

var result [][]int
var path []int

//Solution01 77. 组合 ：给定两个整数 n 和 k，返回 1 ... n 中所有可能的 k 个数的组合。
func Solution01(n int, k int) [][]int {
	backtracking01(n, k, 1)
	return result
}

func backtracking01(n int, k int, startIndex int) {
	//终止条件
	if len(path) == k {
		//收集结果
		count := make([]int, 0)
		for i := 0; i < len(path); i++ {
			count = append(count, path[i])
		}
		result = append(result, count)
		return
	}
	//循环遍历
	for i := startIndex; i <= n; i++ {
		//处理节点
		path = append(path, i)
		//递归
		backtracking01(n, k, i+1)
		//回溯
		path = path[:len(path)-1]

	}
}
