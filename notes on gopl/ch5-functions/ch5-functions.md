# 函数

## 函数声明

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