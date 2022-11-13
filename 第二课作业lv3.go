package main

import (
	"fmt"
	"github.com/eiannone/keyboard"
	"time"
)

func main() {
	var (
		re   = make(chan struct{})
		ai   = make(chan struct{})
		fi   = make(chan struct{})
		stop = make(chan struct{})
		r    = 0
		a    = 0
		f    = 1
		//s    rune
	)
	go ready(re)
	go ready(re)
	go ready(re)
	go ready(re)
	go ready(re)
	go ready(re)
	go ready(re)
	go ready(re)
	go ready(re)
	go ready(re)
	go aim(ai)
	go aim(ai)
	go aim(ai)
	go aim(ai)
	go aim(ai)
	go fire(fi)
	go fire(fi)
	go fire(fi)
	go func() {
	way:
		s, _, err := keyboard.GetSingleKey()
		if err != nil {
			panic(err)
		}
		if s == 'q' {
			stop <- struct{}{}
		} else {
			goto way
		}
	}()
attack:
	for {
		select {
		case _ = <-re:
			if f == 1 {
				r = 1
				a = 0
				f = 0
				fmt.Printf("装弹 ->")
			}
		case _ = <-ai:
			if r == 1 {
				r = 0
				a = 1
				f = 0
				fmt.Printf("瞄准 ->")
			}
		case _ = <-fi:
			if a == 1 {
				r = 0
				a = 0
				f = 1
				fmt.Printf("发射！\n")
			}
		case _ = <-stop:
			break attack
		}
	}
}
func ready(re chan struct{}) {
	for {
		time.Sleep(1 * time.Second)
		re <- struct{}{}
	}
}
func aim(ai chan struct{}) {
	for {
		time.Sleep(1 * time.Second)
		ai <- struct{}{}
	}
}
func fire(fi chan struct{}) {
	for {
		time.Sleep(1 * time.Second)
		fi <- struct{}{}
	}
}
