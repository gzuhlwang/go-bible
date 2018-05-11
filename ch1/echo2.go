package main

import (
	"os"
	"fmt"
)

func main(){
	s,sep:="",""
	for _,v :=range os.Args[1:]{   //注意os.Args[1:]不能写成os.Args or os.Args[:]
		s+=sep+v
		sep=" "
	}
	fmt.Println(s)
}
