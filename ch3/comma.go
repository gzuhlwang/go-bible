package main

import (
	"fmt"
)

func main(){
	s:=comma("12")
	fmt.Println(s)

}

func comma(s string) string{
	n:=len(s)

	if n<=3{
		return s
	}

	return comma(s[:n-3])+","+s[n-3:]
}
