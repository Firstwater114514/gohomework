package main

import (
	"fmt"
	"time"
)

type 钱 struct {
	money int
	all   int
}

var (
	sum      = 0
	a        = 10000
	泡泡收钱 钱
)

func main() {
	泡泡的钱包 := make(chan 钱, a)
	网友的钱 := 钱{
		money: 50,
	}
	go func() {
		for i := 0; i < a; i++ {
			泡泡的钱包 <- 网友的钱
			泡泡收钱 = <-泡泡的钱包
			sum += 泡泡收钱.money
		}
	}()
	fmt.Println("收钱中")
	time.Sleep(2 * time.Second)
	泡泡收钱.all = sum
	fmt.Printf("一共收了%v元\n", 泡泡收钱.all)
}
