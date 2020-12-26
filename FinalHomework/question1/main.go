package main

import (
	"fmt"
	"math/rand"
	"strconv"
	"time"
)

const (
	//黑桃
	Spade = 0
	//红桃
	Hearts = 1
	//梅花
	Club = 2
	//方块
	Diamond = 3
)

type Poker struct {
	Num int
	Flower int
}
type Pokers []Poker

func init() {
	//利用系统时间生成种子
	rand.Seed(time.Now().Unix())
}

func main(){
	player1:= CreatePokers()
	player1.RandPoker()
	player1.Print()

	player1.Sort()
	player1.Print()
}

func (p Pokers)RandPoker(){
	//通过随机生成下标，从后往前交换牌
	//由于随机下标生成的区间是[0,n)，若从前往后交换，可能会造成某些牌未交换仍在原位
	for i := len(p) - 1; i > 0; i-- {
		num := rand.Intn(i + 1)
		p[i], p[num] = p[num], p[i]
	}
}

//排序
func (p Pokers) Sort(){
	lenPoker:=len(p)

	//数值排序
	for i:=0;i<lenPoker;i++{
		for j:=i+1;j<lenPoker;j++{
			if  p[i].Num > p[j].Num{
				p[i], p[j] = p[j], p[i]
			}
		}
	}

	//花色排序
	for i:=0;i<lenPoker;i++{
		for j:=i+1;j<lenPoker;j++{
			if  p[i].Flower > p[j].Flower{
				p[i], p[j] = p[j], p[i]
			}
		}
	}}

//洗牌
func (p Poker)PokerSelf()string  {
	var buffer string

	switch p.Flower {
	case Spade:
		buffer += "♤"
	case Hearts:
		buffer += "♡"
	case Club:
		buffer += "♧"
	case Diamond:
		buffer += "♢"
	}
	switch p.Num {
	case 13:
		buffer += "2"
	case 12:
		buffer += "A"
	case 11:
		buffer += "K"
	case 10:
		buffer += "Q"
	case 9:
		buffer += "J"
	default:
		buffer += strconv.Itoa(p.Num+2)
	}

	return buffer
}

func CreatePokers()(pokers Pokers)  {
	for i := 1; i < 14; i++ {
		for j := 0; j < 4; j++ {
			pokers = append(pokers,Poker{
				Num:    i,
				Flower: j,
			})
		}
	}
	return
}

func (p Pokers)Print()  {
	for _, i2 := range p {
		fmt.Print(i2.PokerSelf()," ")
	}
	fmt.Println()
}