//You are given an undirected weighted graph of n nodes (0-indexed), represented
// by an edge list where edges[i] = [a, b] is an undirected edge connecting the no
//des a and b with a probability of success of traversing that edge succProb[i]. 
//
// Given two nodes start and end, find the path with the maximum probability of 
//success to go from start to end and return its success probability. 
//
// If there is no path from start to end, return 0. Your answer will be accepted
// if it differs from the correct answer by at most 1e-5. 
//
// 
// Example 1: 
//
// 
//
// 
//Input: n = 3, edges = [[0,1],[1,2],[0,2]], succProb = [0.5,0.5,0.2], start = 0
//, end = 2
//Output: 0.25000
//Explanation: There are two paths from start to end, one having a probability o
//f success = 0.2 and the other has 0.5 * 0.5 = 0.25.
// 
//
// Example 2: 
//
// 
//
// 
//Input: n = 3, edges = [[0,1],[1,2],[0,2]], succProb = [0.5,0.5,0.3], start = 0
//, end = 2
//Output: 0.30000
// 
//
// Example 3: 
//
// 
//
// 
//Input: n = 3, edges = [[0,1]], succProb = [0.5], start = 0, end = 2
//Output: 0.00000
//Explanation: There is no path between 0 and 2.
// 
//
// 
// Constraints: 
//
// 
// 2 <= n <= 10^4 
// 0 <= start, end < n 
// start != end 
// 0 <= a, b < n 
// a != b 
// 0 <= succProb.length == edges.length <= 2*10^4 
// 0 <= succProb[i] <= 1 
// There is at most one edge between every two nodes. 
// Related Topics 图 
// 👍 49 👎 0
package main

//leetcode submit region begin(Prohibit modification and deletion)
type Edge struct {
	val  int
	prob float64
}


func maxProbability(n int, edges [][]int, succProb []float64, start int, end int) float64 {
	nums:=make([][]Edge,n)
	for i,edge:=range edges{
		from,to:=edge[0],edge[1]
		if nums[from]==nil{
			nums[from] = []Edge{{val: to,prob: succProb[i]}}
		}else {
			nums[from] = append(nums[from],Edge{val: to,prob: succProb[i]})
		}
		if nums[to]==nil{
			nums[to] =[]Edge{{val: from,prob: succProb[i]}}
		}else {
			nums[to] = append(nums[to],Edge{val: from,prob: succProb[i]})
		}
	}
	prob:=make([]float64,n)
	prob[start] = 1.0
	nodes:=make([]Edge,0)
	nodes = append(nodes,Edge{val: start,prob: 1.0})
	for len(nodes)>0{
		node:=nodes[0]
		nodes=nodes[1:]
		if prob[node.val]>node.prob{
			continue
		}
		for _,num:=range nums[node.val]{
			if prob[num.val]<prob[node.val]*num.prob{
				prob[num.val] = prob[node.val]*num.prob
				nodes = append(nodes,Edge{num.val,prob[num.val]})
			}
		}
	}
	return prob[end]
}
//leetcode submit region end(Prohibit modification and deletion)
/**
type Edge struct {
	val  int
	prob float64
}

func maxProbability(n int, edges [][]int, succProb []float64, start int, end int) float64 {
	// 构造图
    nums:=make([][]Edge,n)
    for i,edge:=range edges{
    	from,to:=edge[0],edge[1]

		// 正向
    	if nums[from]==nil{
    		nums[from] = []Edge{{val: to,prob: succProb[i]}}
		}else {
			nums[from] = append(nums[from],Edge{val: to,prob: succProb[i]})
		}
		// 反向
		if nums[to]==nil{
			nums[to] =[]Edge{{val: from,prob: succProb[i]}}
		}else {
			nums[to] = append(nums[to],Edge{val: from,prob: succProb[i]})
		}
	}

	prob:=make([]float64,n)
	prob[start] = 1.0
	nodes:=make([]Edge,0)
	nodes = append(nodes,Edge{val: start,prob: 1.0})
	for len(nodes)>0{
		node:=nodes[0]
		nodes=nodes[1:]
		if prob[node.val]>node.prob{
			continue
		}
		for _,num:=range nums[node.val]{
			if prob[num.val]<prob[node.val]*num.prob{
				prob[num.val] = prob[node.val]*num.prob
				nodes = append(nodes,Edge{num.val,prob[num.val]})
			}
		}
	}
	return prob[end]
}

作者：xuxin
链接：https://leetcode-cn.com/problems/path-with-maximum-probability/solution/go-can-kao-guan-fang-ti-jie-dijkstra-by-xuxin/
来源：力扣（LeetCode）
著作权归作者所有。商业转载请联系作者获得授权，非商业转载请注明出处。
 */