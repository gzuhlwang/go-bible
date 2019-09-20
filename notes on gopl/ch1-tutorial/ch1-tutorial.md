# 入门（intro）
本章会先通过几个小例子对Go有个first expression。

## Hello，World
	package main
	import "fmt"
	func main(){
		fmt.Println("Hello,世界")
	}

Go是一门编译型语言，Go语言的工具链将源代码及其依赖转换成计算机的机器指令，即常说的**静态编译**（相对的概念是动态编译）。Go语言提供的工具都通过一个单独的命令**go**（**go相当于Tendermint Core的tendermint根命令或祖先命令**）调用，go命令有一系列子命令。最简单的一个子命令就是run（**相当于Tendermint Core中的init，node等子命令**）。

$ go run helloworld.go
输出：Hello,世界
Go语言原声支持Unicode，它可以处理全世界任何语言的文本。
如果想保存编译结果以备将来之用，可用使用build子命令：
$ go build helloworld.go
这个命令会生成一个名为helloworld的可执行文件（linux和mac上是helloworld，在Win上是helloworld.exe），之后你可以随时运行它，不需要任何处理（因为是静态编译）。
$ ./helloworld
$ helloworld.exe ?

来探究一下程序本身。Go语言的代码通过包（package）组织，包类似于其他语言里的库（libraries）或者模块（modules）。一个包由位于单个目录下的一个或多个.go源代码文件组成，目录定义包的作用(当我们阅读go-ethereum的源码时，就可以通过包名来大概了解代码组织和功能)。每个源文件都以一条package声明语句开始，这个例子就是package main，表示该文件属于哪个包，紧跟着一系列导入（import）的包，之后是存储在这个文件里的程序语句。

Go的标准库提供了100多个包，以支持常见功能，如输入、输出、排序以及文本处理。

main包比较特殊。**它定义了一个独立可执行的程序，而不是一个库**。在main包里的main函数也很特殊，它是整个程序执行时的入口。main函数所做的事情就是程序做的。当然，main函数一般调用其他包里的函数完成很多工作，比如fmt.Println。

必须告诉编译器源文件需要哪些包，这就是import声明以及随后的package声明扮演的角色。hello world例子只用了一个包（fmt），大多数程序需要导入多个包。

必须恰当导入需要的包，缺少了必要的包或者导入了不需要的包，程序都无法编译通过。这项严格要求避免了程序开发过程中引入未使用的包。

import声明必须放在源文件的package声明之后。随后，则是组成程序的函数、变量、常量、（自定义）类型的声明语句（分别由关键字func,var,const,type定义）。这些内容的声明顺序并不重要。

一个函数的声明由func关键字、函数名、参数列表、返回值列表（main函数的参数列表和返回值都是空的）以及包含在大括号里的函数体组成。

Go语言在代码格式上采取了很强硬的态度。gofmt工具把代码格式化为标准格式，并且go工具中的fmt子命令会对指定包中的所有.go源文件应用gofmt命令。很多文本编辑器（or IDE）都可以配置为保存文件时自动执行gofmt，这样源代码总会被恰当🉐️格式化。

对于大多数用户来说，下载、编译包、运行测试用例、查看Go语言的文档等等常用功能都可以用go的工具完成。

## 命令行参数
大多数的程序都是处理输入，产生输出；这也正是“计算”的定义。但是，程序如何获取要处理的输入数据呢？一些程序生成自己的数据，但通常情况下，输入来自于程序外部：文件、网络连接、其他程序的输出、敲键盘的用户、命令行参数或其他类似输入源。

下面几个例子会讨论其中几个输入源，首先是命令行参数。
os包以跨平台的方式，提供了一些与操作系统交互的函数和变量。程序的命令行参数可从os包的Args变量获取；os包外部使用os.Args访问该变量。

os.Args变量是一个字符串（string）的切片（slice），切片是Go语言的基础概念。和大多数编程语言类似，区间索引时，Go语言里也采用左闭右开形式，即区间包括第一个索引元素，不包括最后一个，因为这样可以简化逻辑（怎么简化？）。

os.Args的第一个元素，os.Args[0]	是命令本身的名字；其它的元素则是程序启动时传给它的参数。

	Go文档里经常使用命令（command）这个词来指代可执行程序，如命令行应用程序。这会让新手在阅读文档时产生困惑。记住，在Go语言里，命令是指任何可执行程序。    ——from 《Go in Action》

下面是Unix里echo命令的一份实现，echo把它的命令行参数打印成一行。
	//ch1/echo1
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
输出结果是：
​        os.Args: [/private/var/folders/g0/857dqf294r3107f870tbqr1m0000gn/T/___go_build_echo1_go Hello Rob Pike]
​        /private/var/folders/g0/857dqf294r3107f870tbqr1m0000gn/T/___go_build_echo1_go
​        Hello Rob Pike

程序导入了两个包，用括号把它们括起来写成列表形式，而没有分开写成独立的import声明（**类似的写法还有变量组，常量组**）。两种形式都合法，列表形式习惯上用得多。包导入顺序并不重要；gofmt工具格式化时按照字母顺序对包名排序。

echo程序可以每循环一次输出一个参数，这个版本却是不断地把新文本追加到末尾来构造字符串。字符串s开始为空，即值为“”，每次循环会添加一些文本；第一次迭代之后，还会再插入一个空格，因此循环结束时每个参数中间都有一个空格。这是一种二次加工，当参数数量庞大时，开销很大，但是对于echo，这种情形不大可能出现。
注意：++和--都只能放在变量名后面，因此--i也非法。++i非法。

Go语言只有for循环这一种循环语句。for循环有多种形式，其中一种如下所示：

```
for initialization;condition;post{
        //zero or more statements
}
```

for循环三个部分不需括号包围。大括号强制要求，做大括号必须和post语句在同一行。
initialization语句是可选的，在循环开始前执行。initialization如果存在，必须是一条简单语句（simple statement），即短变量声明、自增语句、复制语句或函数调用。condition是一个布尔表达式。post语句在循环体执行结束后执行。

for循环的这三部分每个都可以省略，如果省略initialization和post，分号也可以省略：

```
   //传统的while循环
    for condition{
        //...
    }
```

如果连condition也省略了，像下面这样：

```
   //传统的无限循环
    for{
        //...
    }
```

这就变成一个无限循环，尽管如此，还可以用其他方式终止循环，如一条`break`或`return`语句。无限循环用得还是蛮多的，像`go-ethereum`的`miner`模块中著名的方法（self *worker）update就是for和select配合。

		for {
			// A real event arrived, process interesting content
			select {
	
			// Handle ChainHeadEvent
	
			case <-self.chainHeadCh:
	
				self.commitNewWork()
			// Handle ChainSideEvent
	
			case ev := <-self.chainSideCh:
	
				...
			}
		}

for循环的另一种形式，在某种数据类型的区间（range）上遍历，如字符串、切片、带缓冲的channel。echo的第二版本展示了这种形式：
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
每次循环迭代，range产生一对值：索引以及在该索引处的元素值。这个例子不需要索引，但range的语法要求，要处理元素，必须处理索引。一种思路是把索引赋给一个临时变量，如temp，然后忽略它的值，但Go语言不允许使用无用的局部变量，因为这会导致编译错误。

Go语言中这种情况的解决方法是使用空白标识符（blank identifier），即_(也就是下划线)。空白标识符可用于任何语法需要但程序逻辑不需要的时候，例如，在循环里丢弃不需要的循环索引，保留元素值。大多数的Go程序员都会像上面这样使用range和_写echo程序，因为隐式地而非显式地索引os.Args,容易写对。

声明一个变量有好几种方式，下面这些都等价：

    s:=""
    var s string
    var s = ""
    var s string=""

第一种形式是短变量声明，最简洁，但只能用在函数内部，而不能用于包级别变量。第二种形式依赖于字符串的默认初始化零值机制，被初始化为“”。第三种形式用得很少，除非同时声明多个变量。第四种形式显示地标明变量的类型，当变量类型与初始类型相同时，类型冗余，但如果两者类型不同，变量类型就必须了。**实践中一般使用💰两种形式的某个，初始值重要的话就显式地指定变量的类型，否则使用隐式初始化**。

如前文所述，每次循环迭代字符串s的内容都会更新。+=连接原字符串、空格和下个参数，产生新字符串，并把它赋值给s。s原来的内容已经不再使用，将在适当时机对它进行垃圾回收。

如果连接涉及的数据量很大，这种方式代价高昂。一种简单且高效的解决方案是使用strings包的Join函数：
	package main
	
	import (
		"fmt"
		"strings"
		"os"
	)
	
	func main(){
		fmt.Println(strings.Join(os.Args[1:]," "))
	}
最后，如果不关心输出格式，只想看看那输出值，或许只是为了调试，可以用Println为我们格式化输出。

```
fmt.Println(os.Args[1:])
```

这条语句的输出结果跟strings.Join得到的结果很像，只是被放到了一对方括号里。


##1.3查找重复的行
​	对文件做拷贝、打印、搜索、排序、统计或类似事情的程序都有一个差不多的程序结构：一个处理输入的循环，在每个元素上执行计算处理，在处理的同时或最后产生输出。我们会展示一个名为dup的程序的三个版本；灵感来自于Unix的uniq命令，其寻找相邻的重复行。该程序使用的结构和包是个参考范例，可以方便地修改。

dup的第一个版本打印标准输入中多次出现的行，以重复次数开头。该程序将引入if语句，map数据类型以及bufio包。
    //ch1/dup1.go
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
    
    如何运行这个程序呢？至少有以下两种方法。
    第一种是在终端上，先build生成可执行文件，再运行“提示”输入文本，最后使用ctrl+d退出for循环，搞定！
    第二种是将文本提前写在某个文件里，例如input文件。这种情况下是cat input | go run dup1.go。

map存储了键/值（key/value）的集合，对集合元素，提供常数时间的存、取或测试操作。键可以是任意类型，只要其值能用==运算符比较，最常见的例子是字符串；值则可以是任意类型。这个例子中键是字符串，值是整数。内置函数make创建空map。

map的迭代算许并不确定，从实践来看，该顺序随机，每次运行都会变化。这种设计是有意为之的，因为能防止程序依赖特定遍历顺序，而这是无法保证的。

继续来看bufio包，它使处理输入和输出方便又高效。Scanner类型是该包最有用的特性之一，它读取输入并将其拆成行或单词；通常是处理行形式的输入最简单的方法。

程序使用短变量声明创建bufio.Scanner类型的变量input。

```
input:=bufio.NewScanner(os.Stdin)
```

input从程序的标准输入中读取内容。每次调用input.Scan(),即读入下一行，并移除行末的换行符；读取的内容可以调用input.Text()得到。Scan函数在读到一行时返回true，在无输入时返回false。

input从程序的标准输入中读取内容。每次调用input.Scan(),即读入下一行，并移除行末的换行符；读取的内容可以调用input.Text()得到。Scan函数在读到一行时返回true，在无输入时返回false。

fmt.Printf函数对一些表达式产生格式化输出。该函数的首个参数是格式字符串，指定**后续参数**被如何格式化。各个参数的格式取决于“转换字符”（conversion character），形式为百分号后跟一个字母。例如，%d表示以十进制形式打印一个整型操作数，而%s则表示把字符串型操作数的值展开。

Printf有一大堆这种特性，GO程序员称之为动词（verb）。下面是一些常用的特性：
    %d				十进制整数
    %x,%0,%b		十六进制（默认按小写字母打印），八进制，二进制整数
    %X              十六进制（按大写字母打印）
    %f,%g,%e		浮点数:3.141593  3.14159265  3.141593e+00
    %t				布尔:true或false
    %c				字符(rune)(Unicode码点)
    %s				字符串
    %q				带双引号的字符串“abc”或带单引号的字符'c'
    %v				变量的自然形式（natural format）
    %T				变量的类型
    %%				字面上的百分号标志（无操作数）
    %p				变量的地址

dup1的格式字符串中还含有制表符\t和换行符\n。字符串字面上可能含有这些代表不可见字符的转义字符（escape sequences）。按照惯例，以字母f结尾的格式化函数，如log.Printf和fmt.Printf,都采用fmt.Printf的格式化准则。而以ln结尾的格式化函数，则遵循Println的方式，以跟%v差不多的方式格式化参数，并在最后添加一个换行符。 其中后缀f指format，ln指line。

很对程序要么从标准输入中读取数据要么从一系列命名文件中读取数据。dup程序的下个版本读取标准输入或是使用os.Open打开某个具名文件，并操作它们。
    /ch1/dup2.go
    package main
    
    import (
    	"os"
    	"fmt"
    	"bufio"
    )
    
    func main(){
    	counts:=make(map[string]int)
    	files:=os.Args[1:]
    	if len(files)==0{
    		countLines(os.Stdin,counts)
    	}else{
    		for _,arg:=range files{
    			f,err:=os.Open(arg)
    			if err!=nil{
    				fmt.Fprintf(os.Stderr,"dup2:%v\n",err)
    				continue
    			}
    			countLines(f,counts)
    			f.Close()
    		}
    	}
    
    	for line,n :=range counts{
    		if n>1{
    			fmt.Printf("%d\t%s\n",n,line)
    		}
    	}
    }
    
    func countLines(f *os.File,counts map[string]int){
    	input:=bufio.NewScanner(f)
    	for input.Scan(){
    		counts[input.Text()]++
    	}
    	//ignore potential errors
    }
    
    os.Open函数返回两个值。第一个值是被打开的文件（*os.File）,其后被Scanner读取。os.File是io.Reader的实例。
    
    dup的前两个版本以“流”模式读取输入，并根据需要拆分成多个行。理论上，这些程序可以处理任意数量的输入数据。还有另一个方法，就是一口气把全部输入数据读到内存中，一次分割为多行，然后处理它们。
    /ch1/dup3.go
    package main
    
    import (
        "os"
        "io/ioutil"
        "fmt"
        "strings"
    )
    
    func main(){
        counts:=make(map[string]int)
        for _,filename:=range os.Args[1:]{
            //returned an byte[]
            data,err:=ioutil.ReadFile(filename)
            if err!=nil{
                fmt.Fprintf(os.Stderr,"dup3:%v\n",err)
                continue
            }
            //conversion byte[] to string
            for _,line:=range strings.Split(string(data),"\n"){
                counts[line]++
            }
        }
    
        for line,n :=range counts{
            if n>1{
                fmt.Printf("%d\t%s\n",n,line)
            }
        }
    }

ReadFile函数返回一个字节切片（byte slice），必须把它转换成string，才能用strings.Split分割。

实现上，`bufio.Scanner`、`ioutil.ReadFile`和`ioutil.WriteFile`都使用`*os.File`的Read和Write方法，但是，大多数程序员很少需要直接调用那些低级（lower-level）函数。高级函数（higher-level），像bufio和io/ioutil包中所提供的那些，用起来要容易点。


练习1.3和1.4都比较有实际意义。


    package main
    
    import (
        "os"
        "fmt"
        "time"
        "net/http"
        "io/ioutil"
        "io"
    )
    
    func main(){
        start:=time.Now()
        ch:=make(chan string)
        for _,url :=range os.Args[1:] {
            go fetch(url, ch)
        }
        for range os.Args[1:] {
            fmt.Println(<-ch)
        }
        fmt.Printf("%.2fs elapsed\n",time.Since(start).Seconds())
    
    }
    
    func fetch(url string,ch chan<-string){
        start:=time.Now()
        resp,err:=http.Get(url)
        if err!=nil{
            ch<-fmt.Sprint(err)
            return
        }
        nBytes,err:=io.Copy(ioutil.Discard,resp.Body)
        resp.Body.Close()
        if err!=nil{
            ch <-fmt.Sprintf("while reading %s:%v",url,err)
            return
        }
        secs:=time.Since(start).Seconds()
        ch<-fmt.Sprintf("%.2fs	%7d%s",secs,nBytes,url)
    
    }

io.Copy会把响应的Body内容拷贝到ioutil.Discard输出流中（译注：可以把这个变量看作一个垃圾桶，可以向里面写一些不需要的数据），因为我们需要这个方法返回的字节数，但是又不想要其内容。每当请求返回内容时，fetch函数都会往ch这个channel里写入一个字符串，由main函数里的第二个for循环来处理并打印channel里的这个字符串。





