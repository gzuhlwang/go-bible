package main

import "fmt"

func main(){
	fmt.Println(basename("a/b/c.go"))  //"c"
	fmt.Println(basename("c.d.go"))    //"c.d"
	fmt.Println(basename("abc"))       //"abc"
}

func basename(s string) string{
	//丢弃最后一个	\及其以前的所有东西
	for i:=len(s)-1;i>=0;i--{
		if s[i]=='/'{
			s=s[i+1:]
			break
		}
	}
	//保留最后一个.之前的所有东西
	for i:=len(s)-1;i>=0;i--{
		if s[i]=='.'{
			s=s[:i]
			break
		}
	}
	return s
}