package main

import "fmt"

type movie struct {
	name     string
	director string
	kind     string
	country  string
	language string
	upDay    string
	timeLong string
	score    string
}

func main() {
	var m movie
	m.name = "小黄人大眼萌"
	m.director = "凯尔 巴尔达/皮埃尔 柯芬"
	m.kind = "喜剧/动画"
	m.country = "美国"
	m.language = " 英语 / 西班牙语"
	m.upDay = "2015-09-13(中国大陆) / 2015-07-10(美国)"
	m.timeLong = "91 minutes"
	m.score = "7.7 star"
	fmt.Printf("请输入你的命令\n1.电影名字\n2.电影导演\n3.电影种类\n4.制片地区\n5.电影语言\n6.上映日期\n7.电影时长\n8.豆瓣评分\n9.退出程序\n")
	var a int8
	for {
		a = 0
		fmt.Scanf("%d", &a)
		switch a {
		case 1:
			fmt.Println(m.name)
		case 2:
			fmt.Println(m.director)
		case 3:
			fmt.Println(m.kind)
		case 4:
			fmt.Println(m.country)
		case 5:
			fmt.Println(m.language)
		case 6:
			fmt.Println(m.upDay)
		case 7:
			fmt.Println(m.timeLong)
		case 8:
			fmt.Println(m.score)
		case 9:
			fmt.Println("欢迎下次使用")
			return
		}
	}
}
