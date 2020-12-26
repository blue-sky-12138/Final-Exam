package main

import "fmt"

func main(){
	sum:=1
	for i:=2;i*i<1000;i++{
		if 1000%i==0{
			sum+=i+1000/i
			fmt.Println(sum,1000)
		}
	}
	if sum==6{
		fmt.Println(sum)
	}


}