package main

import (
	"fmt"
	"unicode/utf8"
)

func main(){
	a,b:=10.0,5.0
	fmt.Printf("a/b=%f\n",a/b)  //"a/b=2.000000"

	c,d:=100,11
	fmt.Printf("c/d=%d\n",c/d)  //"c/d=9"

	//e,f:=100.0,10
	//fmt.Printf("e/f=%d\n",e/f)  //"invalid operation: e / f (mismatched types float64 and int)"

	var u uint8=255
	fmt.Println(u,u+1,u*u)        //"255 0 1"

	var i int8=127
	fmt.Println(i,i+1,i*i)        //"127 -128 1"

	fmt.Println(1==2)

	bc:="block\xe4\xb8\x96\xe7\x95\x8cchain"
	fmt.Println(bc)       //"block世界chain"

	fmt.Println(HasPrefix(bc,"block"))

	fmt.Println(HasSuffix(bc,"chain"))

	fmt.Println(Contains(bc,"chain"))

	//fmt.Println(bc[5],bc[6])   //"32 99" (' ' and 'c')

	//fmt.Println(bc[4:12])
	//
	//fmt.Println(bc[:])       //"block chain"
	//fmt.Println(bc[6:])
	//fmt.Println(bc[:6])
	//
	//
	//fmt.Println("Hello "+bc[:]) //"Hello block chain"

	s:="Hello, 世界\xAB"
	bb:=[]byte(s)
	s2:=string(bb)

	fmt.Println(&s,&bb,&s2)
	fmt.Println(len(s))                      //"13"
	fmt.Println(utf8.RuneCountInString(s))   //"9"

	for i:=0;i<len(s);{
		r,size:=utf8.DecodeRuneInString(s[i:])
		fmt.Printf("%d\t%c\n",i,r)
		i+=size
	}

	for i,r:=range s{
		fmt.Printf("%d\t%q\t%d\n",i,r,r)
	}

	n:=0

	for range s{
		n++
	}

	fmt.Println(n)
}




func HasPrefix(s,prefix string) bool{
	return len(s)>=len(prefix)&&s[:len(prefix)]==prefix
}

func HasSuffix(s,suffix string) bool{
	return len(s)>=len(suffix)&&s[len(s)-len(suffix):]==suffix
}

func Contains(s,substr string) bool{
	for i:=0;i<len(s);i++{
		if HasSuffix(s[i:],substr){
			return true
		}

	}
	return false
}