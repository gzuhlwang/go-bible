package main

import (
	"fmt"
	"os"
)

func main(){
	fmt.Println(os.Args[1:])   //输出结果为一个slice值
}
