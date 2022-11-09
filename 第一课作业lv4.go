package main

import (
	"fmt"
	"math/rand"
	"time"
)

// /////////前排提示，v6增加的玩法在最后一行有说明
func main() {
	maxNum := 100
	rand.Seed(time.Now().UnixNano())
	secretNumber := rand.Intn(maxNum)
	var sw, n int8
	sw = int8(secretNumber) / 10
	n = 0
	fmt.Println("Please input your guess")
	for {
		var guess int
		_, err := fmt.Scanf("%d", &guess)
		if err != nil {
			fmt.Println("Invalid input. Please enter an integer value")
			continue
		}
		fmt.Println("You guess is", guess)
		if guess > secretNumber {
			fmt.Println("Your guess is bigger than the secret number. Please try again")
		} else if guess < secretNumber {
			fmt.Println("Your guess is smaller than the secret number. Please try again")
		} else {
			fmt.Println("Correct, you Legend!")
			break
		}
		n++
		fmt.Printf("已经猜了%d/6次了", n)
		if n == 4 {
			fmt.Printf("提示：十位是%d\n", sw) //十位提示
		}
		if n == 6 { //6次猜测限制
			fmt.Println("次数达到上限，你已经寄了")
			return
		}
	}
} //  v6说明：增加了次数限制为6次，增加了在第四次猜测时十位数的提示。
