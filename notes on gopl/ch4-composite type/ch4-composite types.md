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
一个短变量声明操作符在一次操作中完成两件事情：**声明一个变量**，并**初始化**。

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

## 结构体
结构类型通过组合一系列固定且唯一的字段来声明。结构里每个字段都会用一个已知类型声明。
这个已知类型可以是内置类型，也可以是其他用户定义的类型。
### 声明一个结构类型
    
    type person struct{
        name string
        age int
    }
    
    type Transaction struct {
    	data txdata
    	// caches
    	hash atomic.Value
    	size atomic.Value
    	from atomic.Value
    }
    
    type txdata struct {
    	AccountNonce uint64          `json:"nonce"    gencodec:"required"`
    	Price        *big.Int        `json:"gasPrice" gencodec:"required"`
    	GasLimit     uint64          `json:"gas"      gencodec:"required"`
    	Recipient    *common.Address `json:"to"       rlp:"nil"` // nil means contract creation
    	Amount       *big.Int        `json:"value"    gencodec:"required"`
    	Payload      []byte          `json:"input"    gencodec:"required"`
    
    	// Signature values
    	V *big.Int `json:"v" gencodec:"required"`
    	R *big.Int `json:"r" gencodec:"required"`
    	S *big.Int `json:"s" gencodec:"required"`
    
    	// This is only used when marshaling to JSON.
    	Hash *common.Hash `json:"hash" rlp:"-"`
    }
    
一旦声明了类型，就可以使用这个类型创建值。
    
    //声明person类型的变量p
    var p person
    fmt.Println(p)             //"{ 0}"  
任何时候，创建一个变量并初始化为其零值，习惯是使用关键字var。这种用法是为了更明确地表示一个变量被设置为零值。
如果变量被初始化为某个非零值，就配合结构字面量和短变量声明操作符来创建变量。

### 结构体字面值
结构体字面量使用一对大括号括住内部字段的初始值。

    //使用结构字面量创建结构类型的值
    p:=person{
        name:"gzuhlwang",
        age:20,
    }
    fmt.Println(p)                //"{gzuhlwang 20}"
    fmt.Printf("%T\n",p)          //"main.person"
第一种形式更常用，以成员名字和相应的值来初始化，可能包含部分或全部的成员。在这种形式的结构体字面值写法中，
如果成员被忽略的话将默认用零值。因为提供了成员的名字，所有成员出现的顺序并不重要。
   
    //不使用字段名，只声明对应的值
    p:=person{"gzuhlwang",20}
第二种形式一般只在定义结构体的包内部使用，或者是在较小的结构体中使用。
## 空结构体（empty struct）

如果结构体没有任何成员就是空结构体，写作struct{}。它的大小为0，也不包含任何信息，但是有时候依然是有价值的。

    package main
    
    import "fmt"
    
    func main() {
    
    	var p struct {
    		X int
    		Y int
    	}
    	fmt.Printf("%T\n", p)       //"struct { X int; Y int }"
    	fmt.Println(p)              //"{0 0}"
    	
    	seen := make(map[string]struct{})
    	s := "hello"
    	if _, ok := seen[s]; !ok {
    		seen[s] = struct{}{}
    	}
    
    	fmt.Println(seen) //"map[hello:{}]"
    
    	var pp struct{}
    
    	fmt.Printf("%T\n", pp)     //"struct {}"
    	fmt.Println(pp)            //"{}"
    }


## 使用结构体

### 结构体比较
如果结构体的全部成员是可以比较的，那么结构体也是可以比较的。
相等比较运算符==将比较两个结构体的每个成员。因此下面两个比较的表达式是等价的。

    type Point struct{X, Y int}
    
    p:=Point{1,2}
    q:=Point{2,1}
    fmt.Println(p.X==q.X&&p.Y==q.Y)     //"false"
    fmt.Println(p==q)                   //"false"
可比较的结构体类型和其他类型一样，可以用作map的key类型。

    type address struct{
        hostname string
        port int
    }   
    
    hits:=make(map[address]int)
    hits[address{"golang.org",443}]++
### 在函数间传递结构体  
结构体可以作为函数的参数和返回值。在Go语言中，所有的函数传参都是值拷贝传入的，函数参数将不再是函数调用时
的原始变量。如果考虑效率的话，较大的结构体通常会用指针的方式传入和返回。

### 结构体嵌入和匿名成员

匿名成员的数据类型必须是命名的类型或指向一个命名的类型的指针。

    type Circle struct{
        Point
        Radius int
    }

不幸的是，结构体字面值并没有简短表示匿名成员的语法，因此下面的语句都不能编译通过：
    
    c:=Circle{X:8,Y:8,Radius:5}     //compile error:unknown field
结构体字面量必须遵循类型声明时的形状。
    
    c:=Circle{
        Point:Point{X:8,Y:8},       //成员名字只能是Point，成员名字是由其类型隐式地决定的
        Radius:5,
    }
    fmt.Printf("%#v\n",c)           //"main.Circle{Point:main.Point{X:8, Y:8}, Radius:5}"
所有匿名成员也有可见性的规则约束。

**但是为什么要嵌入一个没有任何子成员类型的匿名成员类型呢？**

答案是匿名类型的方法集。简短的点运算符语法可以用于选择匿名成员类型的成员，也可以用于访问它们的方法。实际上，
外层的结构体不仅仅是获得了匿名成员类型的所有成员，而且也获得了该类型可导出的全部方法。这个机制可以用于将一个
有简单行为的对象组合成有复杂行为的对象。组合（composition）是Go语言中面向对象编程的核心。

    package main
    
    import "fmt"
    
    type person struct{
    	name string
    	email string
    }
    
    func (p *person) notify(){
    	fmt.Printf("Sending user email to %s<%s>\n",p.name,p.email)
    }
    type admin struct{
    	p person
    	level string
    }
    //嵌入类型
    //type admin struct{
    //   	person
    //    	level string
    //}
    
    func main(){  
    	ad:=admin{
    		p:person{
    			name:"gzuhlwang",
    			email:"gzuhlwang@gzu.edu.cn",
    		},
    		level:"root",
    	}
    
    	ad.p.notify()
        
        //不采用嵌入类型，提示该方法未定义！
    	ad.notify()                 //"ad.notify undefined (type admin has no field or method notify)"
  
    }
    
因此，通过上面的例子可以看出，Go语言中外部类型复用内部类型的成员和方法是通过**结构体嵌入**实现的。必须遵循这一语言设计规范，只有这样你想要的效果才能出来。    

ps：Go语言允许用户扩展或者修改已有类型的行为。这个功能对代码复用很重要，在修改已有类型以符合新类型的时候也很重要。**这个功能是通过嵌入类型(type embedding)完成的**。嵌入类型是将**已有的类型**直接声明在新的结构类型里。被嵌入的类型被称为新的**外部类型**的**内部类型**。
   **通过嵌入类型，与内部类型相关的标识符会提升到外部类型上**。这些被提升的标识符就像直接声明在外部类型里的标识符一样，也是外部类型的一部分。这样外部类型就**组合**了内部类型包含的所有属性（成员），并且可以添加新的字段和方法。
   外部类型也可以通过声明与内部类型标识符同名的标识符来覆盖内部标识符的字段或者方法。这就是扩展或者修改已有类型的方法。

由于内部类型的提升，**内部类型实现的接口**会自动提升到外部类型。
这意味着由于内部类型的实现，外部类型也同样实现了这个接口。
    
    这个知识点可以参阅《Go语言实战》第5.5节的内容。
值得一提的是，这个知识点在tendermint的启动中用到了。这个知识点的具体用法可以参见我的文章[Tendermint源码分析——启动流程分析](https://blog.csdn.net/keencryp/article/details/80149953)
。  

**其实反过头来觉得，把一些基础知识掌握好对阅读别人的代码是有好处的，否则你会陷入迷茫之中。**

## JSON
JavaScript对象表示法（JSON）是一种用于发送和接收结构化信息的标准协议。类似的协议还有XML、ASN.1和Google的Protocol Buffers。它们各有特色，但由于简洁性、可读性和流行程度
等原因，JSON是应用最广泛的一个。

JSON是对JavaScript中各种类型的值——字符串、数字、布尔值和对象——Unicode文本编码。

基本的JSON类型有数字（十进制或科学计数法）、布尔值（true or false）、字符串。其中字符串是以双引号包含的Unicode字符序列，支持和Go
语言类似的反斜杠转义特性，不过JSON使用的是UTF-16编码，而不是Go语言的rune类型。

    boolean     true
    number      -273.15
    string      "she said \"Hello,BF\"
    array       ["gold","silver","bronze"]
    object      {"year":1980,
                 "event":"archery",
                 "medals":["gold","silver","bronze"]}

    type Movie struct{
        Title string
        Year  int       `json:"released"`
        Color bool      `json:"color,omitempty"`
        Actors []string 
    }

在结构体声明中，Year和Color成员后面的字符串字面值是结构体成员tag（field tag）。
    
    var movies=[]Movie{
        {Title: "Casablanca", Year: 1942, Color: false,
            Actors: []string{"Humphrey Bogart", "Ingrid Bergman"}}, 
        {Title: "Cool Hand Luke", Year: 1967, Color: true,
            Actors: []string{"Paul Newman"}}, 
        {Title: "Bullitt", Year: 1968, Color: true,
            Actors: []string{"Steve McQueen", "Jacqueline Bisset"}},
        // ...
    }
这样的数据结构特别适合JSON格式，并且在两种之间相互转换也很容易。将一个Go语言中类似movies的
结构体slice转为JSON的过程叫编组（marshaling），也叫序列化(serialization)。编组通过
调用json.Marshal函数完成：
    
    data,err:=json.Marshal(movies)
    if err!=nil{
        log.Fatalf("JSON marshaling failed:%s",err)
    }
    fmt.Printf("%s\n",data)
    //输出
    [{"Title":"Casablanca","released":1942,"Actors":["Humphrey Bogart","Ingrid Bergman"]},{"Title":"Cool Hand Luke","released":1967,"color":true,"Actors":["Paul Newman"]},{"Title":"Bullitt","released":1968,"color":true,"Actors":["Steve McQueen","Jacqueline Bisset"]}]
这种紧凑的表示形式虽然包含了全部的信息，但是很难阅读。为了生成便于阅读的格式，另一个json.MarshalIndent函数将产生
整齐缩进的输出。该函数有两个额外的字符串参数用于表示每一行输出的前缀和每一层级的缩进：

    data,err:=json.MarshalIndent(movies,""," ")
    if err!=nil{
    	log.Fatalf("JSON marshaling failed:%s",err)
    }
    fmt.Printf("%s\n",data)
    //输出
    [
     {
      "Title": "Casablanca",
      "released": 1942,
      "Actors": [
       "Humphrey Bogart",
       "Ingrid Bergman"
      ]
     },
     {
      "Title": "Cool Hand Luke",
      "released": 1967,
      "color": true,
      "Actors": [
       "Paul Newman"
      ]
     },
     {
      "Title": "Bullitt",
      "released": 1968,
      "color": true,
      "Actors": [
       "Steve McQueen",
       "Jacqueline Bisset"
      ]
     }
    ]

在编码时，默认使用Go语言结构体成员名字作为JSON的对象（通过reflect反射计数），只有导出
的结构体成员才会被编码，这也就是我们为什么选择大写字母开头的成员名称。

一个结构体成员tag是和在编译阶段关联到该成员的元数据字符串。
    
     Year  int       `json:"released"`
     Color bool      `json:"color,omitempty"`

结构体的成员tag可以是任意的字符串字面值，但通常是一系列用空格分隔的key:"value"键值对序列。因为
值包含双引号字符，因此成员tag一般用原生字符串字面值的形式书写。以json开头的key对应的值
用于控制encoding/json包的编解码的行为，并且encoding/...下面其他的包也遵循这样的约定。成员tag
中json对应值的第一部分用于指定JSON对象的名字。比如Color成员的tag，还带了一个额外的omitempty选项，
表示当Go语言结构体成员为空或零值时不生成JSON对象（这里false为零值）。

编码的逆操作是解码，对应将JSON数据解码为Go语言的数据结构，Go语言中一般叫unmarshaling，也叫反序列化（deserialization），通过json.Unmarshal函数完成。
下面的代码将JSON格式的电影数据解码为一个结构体slice，结构体中只有Title成员。通过定义合适的Go语言数据结构，我们可以选择性
地解码JSON中感兴趣的成员。当Unmarshal函数调用返回，slice将只含有Title信息值填充，其他JSON
成员将被忽略。
    
    var titles []struct{Title string}         //注意学习这种写法
    if err:=json.Unmarshal(data,&titles);err!=nil{
        log.Fatalf("JSON Unmarshaling failed:%s",err)
    }
    fmt.Println(titles)                       //"[{Casablanca} {Cool Hand Luke} {Bullitt}]"

## 文本和HTML模版

todo
 
    
##章小结

    1、数组是构造切片和映射的基石。
    2、内置函数make可以创建切片和映射，并指定原始的长度和容量（容量是针对slice）。也可以直接使用切片和映射字面量，或者使用字面量作为变量的初始值。
    3、切片有容量限制，不过可以使用内置的append函数扩展容量。
    4、映射的增长没有容量或任何限制。
    5、内置函数len可以用来获取切片或者映射的长度。
    6、内置函数cap只能用于切片。

## 参考资料

1、《Go语言实战》中相关章节的内容（主要集中在chapter 4和chapter 5.5）