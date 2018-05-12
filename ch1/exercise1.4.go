package main

import (
	"os"
	"fmt"
	"bufio"
)

func main(){
	counts:=make(map[string]int)
	countsFiles:=make(map[string][]string)
	files:=os.Args[1:]
	if len(files)==0{
		countLines(os.Stdin,counts,countsFiles)   //从标准输入读取
	}else{
		for _,arg:=range files{       			  //从具名文件读取
			f,err:=os.Open(arg)
			if err!=nil{
				fmt.Fprintf(os.Stderr,"dup2:%v\n",err)
				continue
			}
			countLines(f,counts,countsFiles)
			f.Close()
		}
	}

	for line,n :=range counts{
		if n>1{
			fmt.Printf("%s\t%d\t%s\n",countsFiles[line],n,line)
		}
	}
}

func countLines(f *os.File,counts map[string]int,countFiles map[string][]string){
	filename:=f.Name()
	input:=bufio.NewScanner(f)
	for input.Scan(){
		counts[input.Text()]++
		if !arrayContains(countFiles[input.Text()],filename){
			countFiles[input.Text()]=append(countFiles[input.Text()],filename)
		}

	}
	//ignore potential errors
}

func arrayContains(array []string,value string) bool{
	for _,v :=range array {
		if v==value{
			return true
		}
	}
	return false
}

