package main

import (
	"bufio"
	"os"
	"fmt"
)

func main(){
	counts:=make(map[string]int)
	input:=bufio.NewScanner(os.Stdin)

	for input.Scan(){   //quit loop by typing Ctrl+D
		counts[input.Text()]++
	}
	//filter duplicate line and its content
	for line,n := range counts{
		if n>1{
			fmt.Printf("The num of dup line is %d and its corresponding content are %s\n",n,line)
		}
	}
	fmt.Println("The result of input:",counts)
}
