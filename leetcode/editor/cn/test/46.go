package main

import "fmt"

//func permute(nums []int) (res [][]int) {
//	var dfs func(start int, temp []int, values []int)
//	dfs = func(start int, temp []int, values []int) {
//		if len(temp) == len(nums) {
//			path := make([]int, len(temp))
//			copy(path, temp)
//			res = append(res, path)
//			return
//		}
//		for i := 0; i < len(nums); i++ {
//			temp = append(temp, values[i])
//			value1 := values[:i]
//			value2 := values[i+1:]
//			dfs(0, temp, append(value1, value2...))
//			temp = temp[:len(temp)-1]
//		}
//	}
//	dfs(0, []int{}, nums)
//	return
//}
//
func permute(nums []int) (res [][]int) {
	var dfs func(path []int)
	visited := make(map[int]bool)
	dfs = func(path []int) {
		if len(path) == len(nums) {
			temp := make([]int, len(nums))
			copy(temp, path)
			res = append(res, temp)
			return
		}
		for _, n := range nums {
			if v, ok := visited[n]; ok && v {
				fmt.Println(ok, "visited[n]=", visited[n])
				continue
			}

			//if visited[n] {
			//	fmt.Println("continue visited[n]=", visited[n])
			//	continue
			//}
			path = append(path, n)
			visited[n] = true
			dfs(path)
			path = path[:len(path)-1]
			visited[n] = false
		}
	}
	dfs([]int{})
	return

	//res := [][]int{}
	//visited := map[int]bool{}
	//
	//var dfs func(path []int)
	//dfs = func(path []int) {
	//	if len(path) == len(nums) {
	//		temp := make([]int, len(path))
	//		copy(temp, path)
	//		res = append(res, temp)
	//		return
	//	}
	//	for _, n := range nums {
	//		if visited[n] {
	//			continue
	//		}
	//		path = append(path, n)
	//		visited[n] = true
	//		dfs(path)
	//		path = path[:len(path)-1]
	//		visited[n] = false
	//	}
	//}
	//
	//dfs([]int{})
	//return res
}

func main() {
	fmt.Println(permute([]int{1, 2, 3}))
}
