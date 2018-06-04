# 函数

## 函数声明

    func name(parameter-list) (result-list){
        body
    }

函数的类型称作函数签名。

你可能会偶尔遇到没有函数体的函数声明，这表示该函数不是以Go实现的。这样的声明定义了函数的签名。
    
    pakcage math
    
    func Sin(x float64) float64  //汇编语言实现
    
## 5.2 递归
    
    $ ./fetch https://golang.org | ./findlinks1
其中|为管道符，把前面的命令执行结果输送到管道符号后的命令作为参数执行。

## 5.3 多返回值

    包名.函数名（fmt.Printf）   vs      实例.方法名(resp.Body.Close())
函数只能由包名调用。

不宜过渡使用bare return。

## 5.4 错误
   
在Go中，函数运行失败时会返回错误信息，这些错误信息被认为是一种预期的值（因为error是内建的接口，因此错误信息是值）而非异常（exception），这使得Go有别于
那些将函数运行失败看作是异常的语言。
### 5.4.1 错误处理策略

    fmt.Fprintf(os.Stderr,"Site is down:%v\n",err)
Fprintf中F代表file,其第一个参数是*os.File类型。

Go中大部分函数的代码结构几乎相同，首先是一系列的初始检查，防止错误发生，之后是函数的实际逻辑。
这一点可以在/core/types/tx_pool.go中的add方法可见一斑。

# 5.5 函数值（Function values）

在Go中，函数被看作第一类值（first-class values）：函数像其他值一样，拥有类型，可以被复制给其他变量，传递给函数，从函数返回（作为返回列表）。

    对函数值的调用类似函数调用。
    
函数是一种值。类比i:=8.

    i:=8   //编译器会自动推导出i的类型为int
    
    func square(n int) int {return n*n}
    
    f:=square               //f为函数变量，其类型为func(int) int
    fmt.Println(f(3))       //"9"
    fmt.Printf("%T\n",f)    //"func(int) int"
**数类型的零值是nil**。调用值为nil的函数会引发panic错误。

    var  f func(int) int
    f(3)                    //此处f的值为nil，会引发panic错误

函数值可以与nil比较：
    
    var f func(int) int
    if f!=nil{
        f(3)
    }
但是函数值之间是不可比较的（因为类型不一样，咋比较？），也不能用函数值作为map的key。

## 5.6 匿名函数
命名函数只能在包级别语法块中被声明，通过函数字面量（function literal），我们可以绕过这一限制，在任何表达式中表示一个函数值。
函数字面量的语法和函数声明相似，区别在于func关键字后面没有函数名。函数字面量是一种表达式，它的值称为匿名函数。
    
## 5.7 变长函数

尽管...int参数就像函数体内的slice，但变长函数的类型和一个带有普通slice参数的函数的类型是不同的。
    
    func f(...int){}
    func g([]int){}
    
    fmt.Printf("%T\n",f)       //"func(...int)"
    fmt.Printf("%T\n",g)       //"func([]int)"
    
interface{}类型意味着可以接受任何值。