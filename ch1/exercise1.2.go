package main

import (
	"os"
	"fmt"
)

func main(){
	s,sep:="",""
	for i,v:=range os.Args[1:]{
		//fmt.Printf("The index is %d and its value is %s\n",i,os.Args[i])  //这种写法会打印第一个参数
		fmt.Printf("The index is %d and its value is %s\n",i,v)       //这种写法才是对的，注意这个小细节
		s+=sep+v
		sep=" "
	}
	fmt.Println(s)
}
