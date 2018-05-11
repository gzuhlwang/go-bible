package main

import (
	"os"
	"fmt"
)

func main(){
	var s,sep string
	fmt.Println("os.Args:",os.Args)  //输出整个os.Args
	fmt.Println(os.Args[0])            //os.Args中第一个参数是该程序的路径
	for i:=1;i<len(os.Args);i++{
		s+=sep+os.Args[i]
		sep=" "
	}
	fmt.Println(s)
}
