package main

import (
	"fmt"
	"time"
)

type 钱 struct {
	money int
}

var (
	a          = 10000
	泡泡收钱   钱
	泡泡总收入 钱
)

func main() {
	泡泡的钱包 := make(chan 钱, a)
	网友的钱 := 钱{
		money: 50,
	}
	泡泡总收入.money = 0
	for i := 0; i < a; i++ {
		go func() {
			泡泡的钱包 <- 网友的钱
		}()
	}
	go func() {
		for {
			泡泡收钱 = <-泡泡的钱包
			泡泡总收入.money = 泡泡收钱.money + 泡泡总收入.money
		}
	}()
	fmt.Println("收钱中")
	time.Sleep(2 * time.Second)
	fmt.Printf("一共收了%v元\n", 泡泡总收入.money)
}
