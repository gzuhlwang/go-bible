# program structure
函数以源文件和包的方式被组织。
## 命名（names）
Go语言中的函数名、变量名、常量名、类型名、包名、语句标号等所有的命名，都遵循一个简单的命名规则：一个名字必须以一个字母（Unicode字母）或下划线开头，后面可以跟任意数量的字母、数字或下划线。Go语言中变量名大小写敏感，例如heapSort和Heapsort是两个不同的名字。

Go的关键字比较少，只有25个，比26个字母还少一个。它们是
流程控制： for switch break continue goto default fallthrough if else case 			 range
复合类型： map chan func interface struct
基础类型： int
并发功能： go	select
其他： import type var return	const defer

此外，还有大约30多个预定义的名字，比如int和true等，主要对应内建的常量、类型和函数。
内建常量：true false iota nil
内建类型： int int8 int16 int32 int64
			 uint uint8 uint16 uint32 uint64 uintptr
			 float32 float64 complex128 complex64
			 bool	byte rune string error
内建函数：make len cap new append copy close delete
        complex real imag
        panic recover
        
名字的首字母的大小写决定了名字在包外的可见性。如果一个名字是大写字母开头的（**必须是在函数外部定义的包级别名字；包级别函数名也是包级别名字**），那么它是可导出的（exported），也就是说可以被外部的包访问，例如fmt包的Println函数就是可导出的。包本身的名字一般总是用小写字母。

Go语言的风格是尽量使用短小的名字，对于局部变量尤其如此。通常来说，如果一个名字的作用域比较大，生命周期也比较长，那么用长的名字将会更有意义。

在习惯上，GO语言程序员推荐使用驼峰式命名。你会在标准库中看到QuoteRuneToASCII和parseRequestLine这样的函数命名。但一般不会用quote_rune_to_ASCII和parse_request_line这样的命名。而像ACSII和HTML这样的缩略词则避免使用大小写混合的写法，它们可能被称为htmlEscape、HTMLEscape，但不会是escapeHtml。

##声明（declarations）
声明语句定义了程序的各种实体对象以及部分或全部的属性（常见于结构体）。Go语言主要有四种类型的声明语句：var、const、type和func，分别对应变量、常量、类型和函数实体对象的声明。

一个GO语言编写的程序对应一个或多个以.go为文件后缀名的源文件中。每个源文件以包的声明语句开始，说明该源文件是属于哪个包。包声明语句之后是import语句导入依赖的其他包，然后是包一级的类型、变量、常量、函数的声明语句，包一级的各种类型的声明语句的顺序无关紧要。

pakcage 
import
type struct
const
func

一个函数的声明由一个函数名字、参数列表（由函数的调用者提供参数变量的具体值）、一个可选的返回值列表和包含函数定义的函数体组成。执行函数从函数的第一个语句开始，依次顺序执行**直到遇到return返回语句**，如果没有返回语句则是执行到函数末尾，然后返回到函数调用者。





	

