package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {

	rand.Seed(time.Now().UnixNano()) // 取纳秒时间戳，可以保证每次的随机数种子都不同
	for i := 0; i < 5; i++ {
		fmt.Println(rand.int(1, 4064)) // Intn(n)返回一个取值范围在[0,n)的伪随机int值
	}
}
