package main

import "fmt"

// 复杂了
//func generateMatrix(n int) [][]int {
//	if n <= 0 {
//		// 注意return可以是nil
//		return nil
//	}
//
//	res := make([][]int, n)
//	for i := 0; i < n; i++ {
//		res[i] = make([]int, n)
//	}
//
//	directions := [][]int{
//		{0, 1},
//		{1, 0},
//		{0, -1},
//		{-1, 0},
//	}
//
//	start := 1
//
//	x := 0
//	y := 0
//	for i := 1; i < (n+1)/2; i++ {
//		for _, direct := range directions {
//			for j := 0; j < n - 1; j++ {
//				x = x + direct[0]
//				y = y + direct[1]
//				res[x][y] = start
//				start++
//			}
//		}
//	}
//	return res
//}

func generateMatrix(n int) [][]int {
	if n <= 0 {
		// 注意return可以是nil
		return nil
	}

	res := make([][]int, n)
	for i := 0; i < n; i++ {
		res[i] = make([]int, n)
	}

	top, bottom := 0, n-1
	left, right := 0, n-1

	start := 1
	for top <= bottom {
		for i := left; i <= right; i++ {
			res[top][i] = start
			start++
		}
		top++

		for i := top; i <= bottom; i++ {
			res[i][right] = start
			start++
		}
		right--

		for i := right; i >= left; i-- {
			res[bottom][i] = start
			start++
		}
		bottom--

		for i := bottom; i >= top; i-- {
			res[i][left] = start
			start++
		}
		left++
	}
	//
	//for y <= right {
	//	res[x][y] = start
	//	y++
	//	start++
	//}
	//top++
	//
	//for x < bottom {
	//	res[x][y] = start
	//	x++
	//	start++
	//}
	//right--
	//
	//for y > left {
	//	res[x][y] = start
	//	y--
	//	start++
	//}
	//bottom--
	//
	//for x > top {
	//	res[x][y] = start
	//	x--
	//	start++
	//}
	//left++
	//}
	return res
}

func main() {
	fmt.Println(generateMatrix(3))
}

//func generateMatrix(n int) [][]int {
//	result:=[][]int{}
//	for i:=0;i<n;i++{
//		result=append(result,make([]int,n))
//	}
//
//	size:=n*n
//	a,b:=0,n-1      //左右边界
//	c,d:=0,n-1      //上下边界
//
//	idx:=1
//	for idx<=size{
//		//向右走
//		for i:=a;i<=b;i++{
//			result[c][i]=idx
//			idx++
//		}
//		c++     //上边界收缩
//		//向下走
//		for i:=c;i<=d;i++{
//			result[i][b]=idx
//			idx++
//		}
//		b--     //右边界收缩
//		//向左走
//		for i:=b;i>=a;i--{
//			result[d][i]=idx
//			idx++
//		}
//		d--     //。。。
//		//向上走
//		for i:=d;i>=c;i--{
//			result[i][a]=idx
//			idx++
//		}
//		a++     //。。。
//	}
//	return result
//}
//
//作者：jin-ji-de-miao-miao
//链接：https://leetcode-cn.com/problems/spiral-matrix-ii/solution/jian-dan-ming-liao-by-jin-ji-de-miao-miao/
//来源：力扣（LeetCode）
//著作权归作者所有。商业转载请联系作者获得授权，非商业转载请注明出处。
