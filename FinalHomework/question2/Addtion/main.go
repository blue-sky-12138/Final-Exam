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
		go IfPrimeNumber(i)
	}

	<-ENDTip
	fmt.Printf("%v\n",result)
}

func IfPrimeNumber(number int){
	//手动处理小于6的数
	if number==1 || number==2 || number==3 || number==5{
		AddResult(number)

		//所有数已判断完毕
		if number==END{
			ENDTip<-1
		}
		return
	}else if number==4 || number==6{

		//所有数已判断完毕
		if number==END{
			ENDTip<-1
		}
		return
	}

	//排除2,3,5的倍数
	if number%2==0 || number%3==0 || number%5==0{

		//所有数已判断完毕
		if number==END{
			ENDTip<-1
		}
		return
	}

	//上述之后，这个数只可能被被6n+1或6n+5的数整除
	//因数的平方必有一个小于本身，可减小一半判断
	for i:=6;i*i<number;i+=6{
		if number%(i+1)==0 || number%(i+5)==0{

			//所有数已判断完毕
			if number==END{
				ENDTip<-1
			}
			return
		}
	}
	AddResult(number)

	//所有数已判断完毕
	if number==END{
		ENDTip<-1
	}
}

func AddResult(number int){
	lock.Lock()
	result = append(result, number)
	lock.Unlock()
}