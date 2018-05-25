# 复合类型
    
    数组
    切片
    map
    结构体

## 数组

数组是一个固定长度的特定类型元素组成的序列，一个数组可以由**零个**或多个元素组成。

### 数组的声明

    var a [0]int   //"[]" 空数组
    var b [3]int   //"[0 0 0]"
    c:=[1]int{}
    fmt.Println(c) //"[0]"  注意区分a和c
    
### 数组的初始化——数组字面值

    var q [3]int=[3]int{1,2,3}
    var r [3]int=[3]int{1,2}
    fmt.Println(r[2])            //"0"

数组的长度是数组类型的一部分。

    q:=[...]int{1,2,3}
    fmt.Printf("%T\n",q)         //[3]int
    
    q=[4]int{1,2,3,4}            //compile error:cannot use [4]int literal (type [4]int) as type [3]int in assignment
    
如果一个数组的元素类型是可比较的，那么数组类型是可相互比较的。只有当两个数组的所有元素
对应相等的时候数组才是相等的。

    a:=[...]int{1,2}
    b:=[...]int{2,1}
    fmt.Println(a==b)           //"false"
    
    c:=[3]int{1,2}
    fmt.Println(a==c)           //compile error:invalid operation: a == c (mismatched types [2]int and [3]int)
    
### 在函数间传递数组——数组指针 

    var ptr *[32]int            //"<nil>"   数组指针
    var a    [32]*int           //指针数组
在传递大的数组时，我们有必要显式地传入一个数组指针。数组很少用作函数参数。

    b:=[16]byte{}
    fmt.Println(b)             //"[0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0]"

数组字面值[32]byte{}可以生成一个32字节的数组。而且每个数组的元素都是零值初始化，也就是0。

    bb:=[...]bool{}
    fmt.Println(bb)             //"[]"

### 数组的使用范围

数组宜适用于长度已知的场景。例如SHA256等。

## slice

slice是变长序列。一个slice类型一般写作[]T,其中T代表slice中元素的类型。slice的语法和数组很像，只是没有固定长度而已。
一个slice由三分部分构成：指针、长度和容量。指针指向第一个slice元素对应的底层数组元素的地址。要注意：slice的第一个元素
并不一定就是数组的第一个元素。长度对应slice中元素的数目；长度不能超过容量，容量一般是从slice的开始位置到底层数据的结尾
位置。

### 创建和初始化

1、使用内置函数make

当使用make创建时，**必须**指定切片的长度。
   
   //使用长度声明一个字符串切片
    slice:=make([]string,5)
    fmt.Println(len(slice),cap(slice))   //"5 5"

如果指指定长度，那么切片的容量和长度相等。也可以分别指定长度。

    //使用长度和容量声明整型切片
    slice:=make([]int,3,5)          //slice:=make([]int,5)[:3]
    
    for i,v:=range slice{
    	fmt.Printf("i=%d\t slice[%[1]d]=%d\n",i,v)
    }
    //输出
    i=0	 slice[0]=0
    i=1	 slice[1]=0
    i=2	 slice[2]=0
    
    fmt.Println(slice[3])       //"panic: runtime error: index out of range"
分别指定长度和容量时，创建的切片，底层数组的长度是指定的容量，**但是初始化后并不能访问所有的数组元素**。上面的例子中我们可以访问3个元素。

2、使用切片字面量语法

    slice:=[]string{"block","dag","hashgraph"}
    fmt.Println(len(slice),cap(slice))            //"3 3"


## 声明数组 vs 声明切片

如果在[]运算符里指定了一个值，那么创建的就是数组而不是切片。只有不指定值的时候，才会创建切片。
    array:=[3]int{10,20,30}
    fmt.Printf("%T\n",array)           //"[3]int"
    
    slice:=[]int{10,20,30}
    fmt.Printf("%T\n",slice)           //"[]int"
    
## nil和空切片

有时，程序可能需要声明一个值为nil的切片（也称nil切片）。只要在声明时不做任何初始化，就会创建一个nil切片。
    
    //创建nil切片
    var slice  []int                  //"nil"
    fmt.Println(len(slcie),cap(slice),slice==nil)   //"0 0 true"
    
    //使用make创建空的整型切片
    slice:=make([]int,0)
    
    //使用切片字面量创建空的整型切片
    slice:=[]int{}
    fmt.Println(len(slice),cap(slice),slice==nil)   //"0 0 false"

空切片在底层数组包含0个元素，也没有分配任何存储空间。想表示空集合时空切片很有用。注意，如果我们像测试一个slice是否为空，应该使用len(s)==0来判断，
而不应该使用s==nil来判断。**除了和nil相等比较外，一个nil值的slice的行为和其他任意0长度的slice一样**。

## 使用切片

### append函数
内置的append函数用于想slice追加元素。

### copy函数
todo

### len
todo

### cap
map中没有cap操作！

## 比较操作
todo

### 迭代切片
for range

关键字range总是会从切片头部开始迭代。如果想对迭代做更多的控制，依旧可以使用传统的for循环。
    
    slice:=[]int{10,20,30,40}
    
    for index:=2;index<len(slice);index++{
    	fmt.Printf("Index:%d\t Value:%d\n",index,slice[index])
    }
 

### 在函数间传递切片

切片的尺寸很小，在函数间复制和传递切片成本也很低。
    
在64位架构的机器上，一个切片需要24字节的内存：指针字段需要8字节，长度和容量字段分别需要8字节。由于与切片关联
的数据包含在底层数组里，不属于切片本身，所以将切片复制到任意函数的时候，对底层数组大小都不会有影响。
在函数间传递24字节的数据会非常快速、简单。这也是切片效率高的地方。不需要传递指针和处理复杂的语法，只需要复制
切片。

## 总结

slice的用处比数组广泛，需要好好掌握。

## map

### 创建和初始化

1、使用内置的make

    ages:=make(map[string]int)
    ages1:=map[string]int{}
    fmt.Println(ages==ages1)

2、使用map字面值语法

    ages:=map[string]int{
        "alice":13,
        "bob":14,
    }

创建映射时更常用的方法是使用map字面值。

## nil和空map
nil映射也就是没有引用任何哈希表。

    //nil映射
    var colors map[string]string    //只是声明，零值机制初始化
    fmt.Println(colors==nil)        //"true"
    
    //空映射
    colors:=map[string]string{}     //注意{}是字面量语法要求的，不是string{}
    //colors:=map[string]string     //与colors:=map[string]string{}这种写法等价
    fmt.Println(colors==nil)        //"false"

## 使用map

### delete
    
    delete(ages,"bob") 
    delete(ages,"John")             //删除一个不在map中的元素也是安全的

### 查找

    if age,ok:=ages["bob"];!ok{/*...*/}

第二个是一布尔值，用于报告元素是否真的存在。
### len

    fmt.Println(len(ages))    //"2"
map无cap操作！

### 迭代map

    for k,v:=range ages{
        fmt.Printf("Key:%s\tValue:%d\n",k,v)
    }
    //输出
    Key:alice  Value:13
    Key:bob  Value:14

### 比较操作
和slice一样，map之间也不能进行相等比较；唯一的例外是和nil进行比较。


### 在函数间传递映射
和切片类似，传引用，不是传拷贝。保证可以用很小的成本来复制映射。


#章小结

    1、数组是构造切片和映射的基石。
    2、内置函数make可以创建切片和映射，并指定原始的长度和容量（容量是针对slice）。也可以直接使用切片和映射字面量，或者使用字面量作为变量的初始值。
    3、切片有容量限制，不过可以使用内置的append函数扩展容量。
    4、映射的增长没有容量或任何限制。
    5、内置函数len可以用来获取切片或者映射的长度。
    6、内置函数cap只能用于切片。

## 参考资料

1、《Go语言实战》ch4