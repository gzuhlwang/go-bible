package main

import (
	"os"
	"fmt"
)

func main(){
	var s,sep string
	fmt.Println("os.Args:",os.Args)  //输出整个os.Args
	fmt.Println(os.Args[0])            //os.Args中第一个参数是该程序的路径
	for i:=1;i<len(os.Args);i++{       //注意i从1开始，因为我们只需要打印从命令行传进来的参数，索引为0的不符合要求
		s+=sep+os.Args[i]
		sep=" "
	}
	fmt.Println(s)
}
