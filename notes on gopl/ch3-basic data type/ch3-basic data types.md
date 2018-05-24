# 3 基础数据类型
Go语言将数据类型分为四类。

    基础类型
    复合类型（数组、s结构体）
    引用类型（指针、切片、map、函数、通道）
    接口类型
## 3.1 整型

    有符号整数：int   int8(byte)    int16   int32(rune)   int64
    无符号整数：uint  uint8   uint16  uint32  uint64
    其他：uintptr（没有指定具体的bit大小但是足以容纳指针。uintptr类型只有在底层编程时才需要，特别是Go语言和C语言函数库或操作系统接口相交互的地方）
    
在计算机中，有符号整数采用2的补码形式表示，也就是最高bit为用作表示符号位。
    
    一个n-bit的有符号数的值域落在[-2^(n-1),2^(n-1)-1]。
无符号整数的所有bit位都用于表示非负数，值域落在[0,2^n-1]。
例如，int8类型整数的值域是从[-128,127],而uint8类型整数的值域是从[0,255]

二元运算符
    
    算术运算符：+   -   *   /    %   
    逻辑运算符：&   ^   ||  &&   &^   <<  >>  |
    比较运算符：==  !=   <   <=   >   >=

算术运算符+、-、*和/可以适用于整数、浮点数和复数。但取模运算符%仅用于整数间的运算。在Go语言中，%取模运算符的符号和被取模数的符号总是一致的，因此
-5%3和-5%-3的结果都是-2。

除法运算符/的行为则依赖于操作数是否全为整数：
    
    a,b:=10.0,5.0
    fmt.Printf("a/b=%f\n",a/b)  //"a/b=2.000000"
    
    c,d:=100,11
    fmt.Printf("c/d=%d\n",c/d)  //"c/d=9"
    
    //e,f:=100.0,10
    //fmt.Printf("e/f=%d\n",e/f)  //"invalid operation: e / f (mismatched types float64 and int)"
    
如果一个算术运算的结果，不管是有符号或者无符号，如果需要更多的bit位才能正确表示的话，就说明计算结果是溢出了。超过高位的bit位部分将被丢弃。如果原始的数值是有
符号类型，而且最左边的bit为1的话，那么最终结果可能是负的，例如int8的例子：
    
    var u uint8=255
    fmt.Println(u,u+1,u*u)        //"255 0 1"
    
    var i int8=127
    fmt.Println(i,i+1,i*i)        //"127 -128 1"
    
两个相同的整数类型可以使用下面的二元比较运算符进行比较，比较表达式的**结果**为布尔类型(true or false)。

    ==  !=  <   <=  >   >=

**两个相同类型的值可以用==和!=进行比较**。  

一元的加法和减法运算符：
    
    +：一元加法（无效果）
    -：负数

对于整数，+0是0+x的简写，-x则是0-x的简写；对于浮点数和复数，+x就是x，-x则是x的负数。

位操作运算符

    &   位运算 AND
    |   位运算 OR
    ^   位运算 XOR
    &^  位清空 （AND NOT）
    <<  左移
    >>  右移
    
**无符号数**往往只有在位运算或其他特殊的运算场景才会使用，就像bit集合、分析二进制文件格式或者哈希和密码学等。
    
许多整数之间的相互转换并不会改变数值：它们只是告诉编译器如何解释这个值。（例如，分配更大的内存空间，存相同的数）。

任何大小的整数字面值都可以用以0开始的八进制格式书写，例如0666；或用户以0x或0X开头的十六进制格式书写，例如0xdeadbeef。

十六进制数字则更强调数字值的bit位模式。**密码学里经常用十六进制来表示常数**。我们不妨拿SHA-256举例。

    const (
    	chunk     = 64
    	init0     = 0x6A09E667
    	init1     = 0xBB67AE85
    	init2     = 0x3C6EF372
    	init3     = 0xA54FF53A
    	init4     = 0x510E527F
    	init5     = 0x9B05688C
    	init6     = 0x1F83D9AB
    	init7     = 0x5BE0CD19
    	init0_224 = 0xC1059ED8
    	init1_224 = 0x367CD507
    	init2_224 = 0x3070DD17
    	init3_224 = 0xF70E5939
    	init4_224 = 0xFFC00B31
    	init5_224 = 0x68581511
    	init6_224 = 0x64F98FA7
    	init7_224 = 0xBEFA4FA4
    )


字符字面值（rune literal）通过一对单引号直接包含对应字符（character）。

## 3.2 浮点数
    
    float32
    float64
二者的算术规范由IEEE754浮点数国际标准定义，该浮点数规范被所有现代的CPU支持。

浮点数的范围极限值可以在math包找到。这里就不过多介绍范围极限。

通常应该优先使用float64类型，因为float32类型的累计误差很容易扩散，并且float32能精确表示的正整数并不是很大（因为float32的有效bit位只有
23个，其它的bit位用于指数和符号；当整数大于23bit能表达的范围时，float32的表示将出现误差）。

## 3.3 复数

## 3.4 布尔型

一个布尔类型的值只有两种：true和false。

布尔值并不会隐式转化为数字值0或1，反之亦然。必须使用一个显式的if语句辅助转换：
    
    //布尔值——》数字值
    func btoi(b bool) int{
        if b{
            return 1
        }
        return 0
    }
    
    //数字值——》布尔值
    func itob(i int) bool{
        return i!=0
    }

if和for语句的条件部分都是布尔类型的值，并且==和<等比较操作也会产生布尔型的值。

## 3.5 字符串

一个字符串是一个**不可改变**的**字节序列**。（A string is an immutable sequence of bytes.） 
字符串可以包含任意的数据，包含byte值0，但是通常是用来包含人类可读的文本。文本字符串通常被解释为采用UTF8编码的Unicode码点（rune）序列。 

注意：内置的len函数返回一个字符串中的**字节数目**（不是rune字符数目），索引操作s[i]返回第i个字节的**字节值**。

    
    bc:="block chain"
    fmt.Println(len(bc))
    fmt.Println(bc[5],bc[6])   //"32 99" (' ' and 'c')

    
第i个字节**并不一定**是字符串的第i个字符，因为对于**非ASCII字符**的**UTF8编码**会要两个或多个字节。

子字符串操作s[i:j]基于原始的字符串（original string）s的第i个字节开始到第j个字节（并不包含j本身）生成一个新的字符串。生成的新字符串将包含j-i个字节（当然对非ASCII字符就不一定成立）。
    
    fmt.Println(bc[2:6])     //"ock"
    
如果索引超出字符串范围或者j小于i将导致panic异常。

    fmt.Println(bc[11:4])    //"invalid slice index: 11 > 4"
    fmt.Println(bc[3:12])    //"panic: runtime error: slice bounds out of range"
    
i和j都可以被忽略，这样i默认采用0作为开始位置，j默认采用len(bc)作为结束位置。

    fmt.Println(bc[:])       //"block chain"
    fmt.Println(bc[6:])      //"chain"
    fmt.Println(bc[:6])      //"block"
    
+操作符将两个字符串连接构造一个新字符串：

    fmt.Println("Hello"+bc[:])   //"Hello block chain"

字符串可以用==和<进行比较；比较通过**逐字节**比较完成。因此比较的结果是字符串自然编码的字典序。

字符串的值是不可变的：一个字符串包含的字节序列永远不会被改变。但我们可以给一个字符串变量分配一个新字符串值。
    
因为字符串是不可修改的，因此尝试修改字符串内部数据的操作也是被禁止的。  
   
    bc[0]='A'     //compile error：cannot assign to bc[0]
   
   
？ 不变性意味着如果两个字符串共享相同的底层数据的话也是安全的，这使得赋值任何长度的字符串代价是低廉的。 

### 3.5.1 字符串字面值（string literals）

字符串值也可以用字符串字面值方式编写，即将一系列字节序列包含在双引号：
    
    "Hello， 世界"
    
以前看视频的时候动不动就字面量，字面值。其实就是用双引号括起来的字节序列称为字面值。

Go语言的源码文件总是用**UTF8编码**，并且Go语言的文本字符串也以UTF8编码的方式处理，因此我们可以将Unicode码点也写到字符串字面值中。

在一个双引号包含的字符串字面值中，可以用以反斜杠\开头的转义序列（escape sequence）插入任意的数据。下面列举一些常见的ASCII控制代码的转义方式：

    \a   响铃
    \b   退格
    \f   换页
    \n   换行
    \r   回车
    \t   制表符
    \v   垂直制表符
    \'   单引号（只用在'\''形式的rune符号字面值中）
    \"   双引号（只用在"..."形式的字符串字面值中）
    \\   反斜杠
右边这些含义描述就是打印效果。我们不妨举例试一下：

    bc:="block\\bchain"
    fmt.Println(bc)       //"block\bchain"
        
可以通过十六进制或八进制转义在字符串字面值包含任意的字节。一个十六进制的转义形式是\xhh，其中两个h表示十六进制数字(大写或小写都可以)。
一个八进制转义形式是，包含三个八进制的o数字\ooo(0到7)，但是不能超过\377(译注:对应一个字节的范围，十进制为 255)。 **每一个单一的字节表达一个特定的值**。
    
    bc:="block\xe4\xb8\x96\xe7\x95\x8cchain"
    fmt.Println(bc)       //"block世界chain"


一个原生的字符串字面值(raw string literal)形如`...`,使用反引号代替双引号。在原生的字符串字面值中，没有转义操作：
全部的内容都是字面的意思，包含退格和换行。因此一个程序中的原生字符串字面值可能跨越多行（在原生字符串字面值内部是无法直接写字符的，可以用八进制或十六进制转义
连接字符串常量完成）。唯一的特殊处理是会删除回车以保证在所有平台上的值都是一样的，包括那些把回车也放入文本文件的系统（Windows系统会把
回车和换行一起放入文本文件中）。

原生字符串字面值用于编写正则表达式会很方便，因为正则表达式往往会包含很多反斜杠。原生字符串字面值同时广泛用于HTML模版、JSON字面值、命令行提示信息以及那些需要扩展到多行的场景。

### 3.5.2 Unicode

早期计算机世界只有一个ASCII字符集，它使用7bit来表示128个字符：包含英文字母的大小写（计52个）、数字（计10个）、各种标点符号和设备控制符。


随着互联网的发展，混合多种语言的数据变得很常见。如何有效处理这些包含了各种语言的丰富多彩的文本数据呢？

    答案就是使用Unicode（http://unicode.org）.

Unicode码点（code point）对应Go语言中的rune整数类型（rune是int32的等价类型）。


我们可以将一个符文序列（sequence of runes）表示为一个int32序列。这种编码方式叫UTF-32或UCS-4，每个Unicode码点
都使用同样大小的int32来表示。优点是简单统一。缺点是浪费很多存储空间。因为大多数计算机可读的文本是ASCII字符，本来每个
ASCII字符只需要8bit或1字节就能表示。而且即便是常用的字符也远少于65536个，也就是说用16bit编码方式就能表达常用字符。
    
    但是，还有其他更好的编码方法吗？
    
### 3.5.3 UTF-8

UTF-8是一个将Unicode码点编码为**字节序列**的**变长编码**。UTF-8编码由Go语言之父Ken Thompson和Rob Pike共同发明。
现在已经是Unicode的标准。UTF-8编码使用1～4个字节来表示每个Unicode码点，ASCII部分字符只使用1个字节，常用字符部分使用
2～3个字节表示。

变长的编码无法直接用过索引来访问第n个字符，但是UTF-8编码有很多额外的优点。
    
    1、UTF-8编码比较紧凑，完全兼容ASCII码并且可以自动同步。
    2、它是一个前缀编码。当从左向右解码时不会产生任何歧义也不需要向前查看。
    3、UTF-8编码的顺序和Unicode码点的顺序一致
    4、没有任何字符的编码是其他字符编码的字串。


得益于UTF-8编码优良的设计（尽管我还没能理解设计上的优雅之处），诸多字符串操作都不需要解码操作。

    //测试一个字符串是否是另一个字符串的前缀
    func HasPrefix(s,prefix string) bool{
    	return len(s)>=len(prefix)&&s[:len(prefix)]==prefix
    }
    
    //后缀测试
    func HasSuffix(s,suffix string) bool{
    	return len(s)>=len(suffix)&&s[len(s)-len(suffix):]==suffix
    }
    
    //子串测试
    func Contains(s,substr string) bool{
    	for i:=0;i<len(s);i++{
    		if HasSuffix(s[i:],substr){
    			return true
    		}
    
    	}
    	return false
    }

对于UTF-8编码后的文本的处理和原始的字符处理逻辑是一样的。但是对应很多其他编码则并不是这样。


    import "unicode/utf8"
    
    s:="Hello, 世界"
    fmt.Println(len(s))                      //"13"
    fmt.Println(utf8.RuneCountInString(s))   //"9"


为了处理这些真实的字符，我们需要一个UTF8解码器。unicode/utf8包提供了该功能，我们可以这样使用：

    for i:=0;i<len(s);{
        r,size:=utf8.DecodeRuneInString(s[i:])
        fmt.Printf("%d\t%c\n",i,r)
        i+=size
    }
    

    for i,r :=range s{
        fmt.Printf("%d\t%q\t%d\n",i,r,r)
    }
    //输出
    0	'H'	72
    1	'e'	101
    2	'l'	108
    3	'l'	108
    4	'o'	111
    5	','	44
    6	' '	32
    7	'世'	19990
    10	'界'	30028

我们可以使用一个简单的循环来统计字符串中字符的数目，像这样：

    n:=0
    for _,_ =range s{
        n++
    }
    //输出 "9"

像其他形式的循环那样，我们也可以忽略不需要的变量：

    n:=0
    for range s{
        n++
    }
    //输出 "9"

或者我们可以直接调用utf8.RuneCountInString(s)函数。

每一个UTF-8字符解码，不管是显式地调用utf8.DecodeRuneInString解码或是在range循环中隐式地解码，如果遇到一个错误
的UTF-8编码输入，将生成一个特别的Unicode字符。在印刷中这个符号通常是一个黑色六角或钻石形状，里面包含一个白色的问号"�"。
当程序遇到这样的一个字符，通常是一个危险信号，说明输入并不是一个完美没有错误的UTF-8字符串。

UTF8字符串作为交换格式是非常方便的，但是在程序内部采用rune序列可能更方便，因为rune大小一致，支持数组索引和方便切割。

### 3.5.4 字符串和Byte切片

标准库中有四个包对字符串处理尤为重要：bytes、strings、strconv和unicode包。strings包提供了许多如字符串的查询（search）、替换（replace）、
比较（compare）、截断（trim）、拆分（split）和合并（join）等功能。


一个字符串由一个**只读**字节**数组**组成，一旦创建，是不可变的。一个字节slice的元素则可自由修改地修改。

字符串和字节slice之间可以相互转换：

    s:="abc"
    b:=[]byte(s)
    s2:=string(b)

bytes包还提供了Buffer类型用于字节slice的缓存。一个Buffer开始是空的，但是随着string、byte或[]byte等类型数据的写入可以动态增长，一个bytes.Buffer
变量并不需要初始化，因为零值也是有效的。

### 3.5.5 字符串和数字的转换

todo

## 3.6 常量

常量表达式的值在编译器计算，而不是在运行期。每种常量的底层类型都是基础类型：boolean、string或数字。

一个常量的声明语句定义了常量的名字，常量的值不可修改，这样可以防止在运行期被意外或恶意的修改。

    const pi=3.1415  
    //pi=12               //compile error:cannot assign to pi
    fmt.Printf("%T",pi)   //"float64"

### 3.6.1 iota常量生成器

    type  Weekday int
    
    const(
        Sunday Weekday=iota
        Monday 
        Thuesday
        Wednesday
        Thursday
        Friday
        Saturday
    )
    
### 3.6.2 无类型常量

todo