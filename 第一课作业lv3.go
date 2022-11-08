package main

import (
	"fmt"
	"math/rand"
	"time"
)

type hero interface {
	say()
}
type ax struct {
	hp  int
	pp  int
	atk int
	def int
}
type gl struct {
	hp  int
	pp  int
	atk int
	def int
}
type rz struct {
	hp  int
	pp  int
	atk int
	def int
}
type yi struct {
	hp  int
	pp  int
	atk int
	def int
}

func (h ax) say() {
	fmt.Println("艾希：你想来一发吗？")
}
func (h gl) say() {
	fmt.Println("盖伦：生命不息，战斗不止！")
}
func (h rz) say() {
	fmt.Println("瑞兹：以意志之名！")
}
func (h yi) say() {
	fmt.Println("易：我的剑就是你的剑!")
}
func show(myHero hero) {
	myHero.say()
}
func main() {
	var (
		hp, pp, atk, def, hpm, ppm int
	)
way2:
	fmt.Printf("英雄选择界面：\n请选择你的英雄\n1.寒冰射手 艾希   2.德玛西亚之力 盖伦   3.流浪法师 瑞兹   4.无极剑圣 易\n5.我不响丸辣\n")
	fmt.Printf("温馨提示：本游戏为单人游戏(pve)    请将运行界面调到半屏以上 :)\n-------------------------------------------------------------------------------\n")
	var op int8
	fmt.Scanf("%d", &op)
way1:
	fmt.Scanf("%d", &op)
	switch op {
	case 1:
		var mh ax
		mh.hp = 570
		mh.pp = 280
		mh.atk = 109
		mh.def = 26
		hp = mh.hp
		pp = mh.pp
		atk = mh.atk
		def = mh.def
		show(mh)
	case 2:
		var mh gl
		mh.hp = 620
		mh.pp = 200
		mh.atk = 121
		mh.def = 36
		hp = mh.hp
		pp = mh.pp
		atk = mh.atk
		def = mh.def
		show(mh)
	case 3:
		var mh rz
		mh.hp = 575
		mh.pp = 300
		mh.atk = 109
		mh.def = 22
		hp = mh.hp
		pp = mh.pp
		atk = mh.atk
		def = mh.def
		show(mh)
	case 4:
		var mh yi
		mh.hp = 599
		mh.pp = 251
		mh.atk = 117
		mh.def = 33
		hp = mh.hp
		pp = mh.pp
		atk = mh.atk
		def = mh.def
		show(mh)
	case 5:
		fmt.Printf("你看这个彬彬，他就是逊啦\n")
		return
	default:
		fmt.Printf("你在敲啥啊，给爷重输\n")
		goto way1
	}
	fmt.Printf("确认你的英雄：\n1.确认   2.重新选择\n-------------------------------------------------------------------------------\n")
	var sure int8
	fmt.Scanf("%d", &sure)
way3:
	fmt.Scanf("%d", &sure)
	switch sure {
	case 1:
		fmt.Printf("战斗开始！你的敌人是亚索！\n")
	case 2:
		goto way2
	default:
		fmt.Printf("快选吧，我已经等不及了（喜）\n")
		goto way3
	}
	hpm = hp
	ppm = pp
	var hp_, jf, def_, atk_, hpm_ int
	hp_ = 511
	jf = 0
	def_ = 32
	atk_ = 110
	hpm_ = 511
	var money, money_ int
	money = 0
	money_ = 0
	var xp, lp, xp_ int8
	xp = 0
	lp = 0
	xp_ = 0
	var x rune
	maxNum := 2
	rand.Seed(time.Now().UnixNano())
	secretNumber := rand.Intn(maxNum)
	if secretNumber == 0 {
		fmt.Println("你很走运，你先手")
		fmt.Printf("按下回车继续\n-------------------------------------------------------------------------------\n")
		fmt.Scanf("%c", &x)
		fmt.Scanf("%c", &x)
	} else {
		fmt.Printf("敌方血量：%d/%d 疾风值：%d 攻击力：%d 护甲值：%d\n", hp_, hpm_, jf, atk_, def_)
		fmt.Printf("你的血量：%d/%d 魔法值：%d/%d 攻击力：%d 护甲值：%d\n", hp, hpm, pp, ppm, atk, def)
		fmt.Println("哦吼，你后手了")
		fmt.Printf("你受到了%d点攻击\n", atk_)
		hp = hp - atk_ + def
		money_ += 25
		jf += 34
		money += 10
		fmt.Printf("按下回车继续\n-------------------------------------------------------------------------------\n")
		fmt.Scanf("%c", &x)
		fmt.Scanf("%c", &x)
	}
	for {
		fmt.Printf("敌方血量：%d/%d 疾风值：%d 攻击力：%d 护甲值：%d\n", hp_, hpm_, jf, atk_, def_)
		fmt.Printf("你的血量：%d/%d 魔法值：%d/%d 攻击力：%d 护甲值：%d\n", hp, hpm, pp, ppm, atk, def)
		fmt.Printf("选择你的战斗方式:\n1.普通攻击   2.使用技能（耗蓝125,加5攻击力）   3.购买护甲（50）   4.使用血瓶（有%d个）   5.使用蓝瓶（有%d个）\n6.购买血瓶（100）   7.购买蓝瓶（75）\n114514.投降逃跑                              你的经济：%d$\n", xp, lp, money)
		var op2 int32
		fmt.Scanf("%d", &op2)
	way4:
		fmt.Scanf("%d", &op2)
		switch op2 {
		case 1:
			fmt.Printf("敌方受到了%d点攻击\n", atk)
			hp_ = hp_ - atk + def_
			money += 25
			if hp_ <= 0 {
				fmt.Printf("哎呦你很勇嘛，打败他了嚎\n你赢力\n")
				return
			}
		case 2:
			if pp >= 50 {
				fmt.Printf("你使用了技能，攻击力+5，蓝量-125\n")
				money += 15
				atk = atk + 5
				pp = pp - 125
				fmt.Printf("敌方受到了%d点攻击\n", atk)
				hp_ = hp_ - atk + def_
				if hp_ <= 0 {
					fmt.Printf("亚索:你干嘛~哈哈~呦~\n好耶，你赢了！\n")
					return
				}
			} else {
				fmt.Printf("你的能量不足力，快选择其他的，不然就要被撅力（悲）\n")
				goto way4
			}
		case 3:
			if money >= 50 {
				def = def + 20
				money = money - 50
				fmt.Println("护甲值增加20！")
			} else {
				fmt.Println("奇怪的护甲值增加了!(因为钱不够，其实啥都没加)")
				goto way4
			}
		case 4:
			if xp >= 1 {
				xp--
				hp = hp + 220
				if hp > hpm {
					hp = hpm
					fmt.Println("血量已回满！")
				} else {
					fmt.Println("恢复了220血量！")
				}
			} else {
				fmt.Println("已经没有血瓶了哼哼哼啊啊啊啊啊啊")
				goto way4
			}
		case 5:
			if lp >= 1 {
				lp--
				pp = pp + 100
				if pp > ppm {
					pp = ppm
					fmt.Println("蓝量已经恢复满了!")
				} else {
					fmt.Println("恢复了100点蓝量")
				}
			} else {
				fmt.Println("你是一个一个一个蓝瓶都没有啊啊啊啊啊啊")
				goto way4
			}
		case 6:
			if money >= 100 {
				money = money - 100
				xp++
				fmt.Printf("获得血瓶，已有%d瓶血瓶\n", xp)
			} else {
				fmt.Println("钱包：我指腚是不行了")
				goto way4
			}
		case 7:
			if money >= 75 {
				money = money - 75
				lp++
				fmt.Printf("获得蓝瓶，已有%d瓶蓝瓶\n", lp)
			} else {
				fmt.Println("钱包:鸭蛋摸鸭蛋~牡蛎摸牡蛎~")
				goto way4
			}
		case 114514:
			fmt.Printf("你给路达油~~\n你逃跑力\n")
			return
		}
		money += 10
		money_ += 10
		fmt.Printf("按下回车继续\n-------------------------------------------------------------------------------\n")
		fmt.Scanf("%c", &x)
		fmt.Scanf("%c", &x)
		fmt.Printf("敌方血量：%d/%d 疾风值：%d 攻击力：%d 护甲值：%d\n", hp_, hpm_, jf, atk_, def_)
		fmt.Printf("你的血量：%d/%d 魔法值：%d/%d 攻击力：%d 护甲值：%d\n", hp, hpm, pp, ppm, atk, def)
		fmt.Printf("敌方回合。（按下回车继续）\n-------------------------------------------------------------------------------\n")
		fmt.Scanf("%c", &x)
		fmt.Scanf("%c", &x)
		if hp_ <= hpm_/2 {
			if xp_ >= 1 {
				xp_--
				hp_ += 220
				if hp_ > hpm_ {
					hp_ = hpm_
				}
				fmt.Println("敌方使用血瓶，血量+220")
			} else if jf >= 100 {
				fmt.Printf("亚索使用技能\n亚索:哈撒给！（他变得更快乐了）\n")
				jf -= 100
				atk_ += 10
				def_ += 10
				money_ += 40
				hp_ += 50
				if hp_ > hpm_ {
					hp_ = hpm_
				}
				fmt.Printf("你受到了%d的伤害\n", atk_+50)
				hp = hp - atk_ - 50 + def
				money += 10
				if hp <= 0 {
					fmt.Printf("鉴定为“只因”\n你寄力,被撅力（大悲）\n")
					return
				}
			} else {
				fmt.Printf("敌方对你进行了攻击，受到%d点伤害\n", atk_)
				money_ += 25
				hp = hp - atk_ + def
				money += 10
				jf += 34
				if hp <= 0 {
					fmt.Printf("亚索:啊？乖乖站好，我要搞死你哈？\n你输力\n")
					return
				}
				goto way5
			}
		} else if money >= 100 {
			money_ -= 100
			xp_++
			fmt.Printf("敌方已做出行动\n")
		} else if jf >= 100 {
			fmt.Printf("亚索使用技能\n亚索:嗖赖！嗖赖！（他变得更快乐了）\n")
			jf -= 100
			atk_ += 10
			def_ += 10
			money_ += 40
			hp_ += 50
			if hp_ > hpm_ {
				hp_ = hpm_
			}
			fmt.Printf("你受到了%d的伤害\n", atk_+50)
			hp = hp - atk_ - 50 + def
			money += 10
			if hp <= 0 {
				fmt.Printf("死了啦,都是你害啦,拜托\n")
				return
			}
		} else {
			fmt.Printf("敌方对你进行了攻击，受到%d点伤害\n", atk_)
			money_ += 10
			hp = hp - atk_ + def
			money += 25
			jf += 34
			if hp <= 0 {
				fmt.Printf("噗唧啪（发出了憋不住笑的声音）\n已经结束嘞！\n")
				return
			}
		}
		pp += 10
	way5:
		fmt.Printf("按下回车继续\n-------------------------------------------------------------------------------\n")
		fmt.Scanf("%c", &x)
		fmt.Scanf("%c", &x)
	}
}
