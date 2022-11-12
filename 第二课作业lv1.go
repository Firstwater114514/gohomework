package main

import (
	"fmt"
	"time"
)

type 解除封印 func(让我看看 string) string

var (
	让我看看 = "阿伟被登dua郎力（悲"
)

func main() {
	go 打电动()
	go 碰到杰哥()
	到杰哥家()
	好康的 := 进入杰哥房间(让我看看, 拳击)
	fmt.Println(好康的)
}
func 打电动() {
	fmt.Println("阿嫲：“阿伟，休息一下吧，去看个书好不好?”")
	time.Sleep(1 * time.Second)
	fmt.Println("阿伟：“死了啦，都是你害的啦，拜托。”")
}

func 碰到杰哥() {
	time.Sleep(2 * time.Second)
	fmt.Println("阿伟走在路上~")
	time.Sleep(1 * time.Second)
	fmt.Println("碰到了杰哥")
	fmt.Println("杰哥：“欢迎你们来我家van,玩累了~就~直接睡觉。")
	fmt.Println("前往杰哥家ing")
}
func 到杰哥家() {
	time.Sleep(8 * time.Second)
	fmt.Println("杰哥：“我的房间里有一些好康的，可以教你，登dua郎。（窃喜”")
	time.Sleep(1 * time.Second)
	fmt.Println("进入杰哥房间")
	time.Sleep(1 * time.Second)
}
func 拳击(让我看看 string) string {
	return 让我看看
}
func 进入杰哥房间(让我看看 string, 解除封印 解除封印) string {
	return 解除封印(让我看看)
}
