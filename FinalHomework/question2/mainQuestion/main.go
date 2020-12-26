package main

import (
	"fmt"
	"sync"
)

var(
	result []int
	ENDTip=make(chan int,0)
	END=123456
	lock sync.Mutex
)

func main(){
	for i:=2;i<=END;i++{
		go SumApproximateNumber(i)
	}

	<-ENDTip
	fmt.Printf("%v\n",result)
}

func SumApproximateNumber(number int){
	sum:=1
	for i:=2;i*i<number;i++{
		if number%i==0{
			sum+=i+number/i
		}
	}
	if sum==number{
		lock.Lock()
		result = append(result, number)
		lock.Unlock()
	}

	//所有数已判断完毕
	if number==END{
		ENDTip<-1
	}
}