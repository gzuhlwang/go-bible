# 方法

一个方法则是一个和特定类型关联的函数。面向对象编程就是使用方法来描述每个数据结构的属性和操作（操作就是算法）。

## 方法声明

在函数声明时，在其名字之前放上一个变量，即是一个方法。这个附加的参数会将该函数附加到这种类型上，即相当于为这种类型定义了一个独占的方法。

    package "geometry"
    
    import "math"
    type Point struct{ X,Y float64}
    
    //普通的函数
    func Distance(p,q Point) float64{
        return math.Hypot(q.X-p.X,q.Y-p.Y)
    }
    
    //Point类型的方法
    func (p Point) Distance(q Point) float64{
        return math.Hypot(q.X-p.X,q.Y-p.Y)
    }
    
附加的参数p称为方法的接收者（receiver），它源自早先的面向对象语言，用来描述主调方法就像向对象发送消息。

接收者的命名原则：简短且方法内唯一。最常用的方法就是使用其类型的首字母，就像Point的p。

调用方法的时候，接收者在方法名前面。这样就和方法声明保持一致。
    
    p:=Point{1,2}
    q:=Point{4,6}
    
    fmt.Println(Distance(p,q))    //"5",函数调用
    fmt.Println(p.Distance(q))    //"5",方法调用

第一个Distance的调用实际上用的是包级别的函数geometry.Distance。

    func (p Point) X(q Point) {
    	p.X=q.X
    	fmt.Println(p.X)
    }
    //"type Point has both field and method named X"

由于方法和字段都是在同一命名空间，所以如果我们在这里声明一个X方法的话，编译器会报错，因为在调用p.X时会有歧义。

因为每一个类型都有命名空间，所以我们能够在其他不同的类型中使用名字Distance作为方法名。定义一个Path类型表示一条线段，同样
也使用Distance作为方法名。因此在给别人讲方法时，需要明确指出其类型。

    //Path是连接多个点的直线段
    type Path []Point
    
    func (path Path) Distance() float64{
        sum:=0.0
        for i:=range path{
            if i>0{
                sum+=path[i-1].Distance(path[i])
            }
        }
        return sum
    }

Path是一个命名的slice，而非Point这样的结构体类型。Go和许多其他面向对象的语言不同，它可以将方法绑定到任何类型上。可以很方便
地为简单的类型（如数值、字符串、slice、map，甚至函数等）定义附加行为。同一个包下的任何类型都可以声明方法，只要不是一个指针或者
一个接口类型。

类型拥有的所有方法名都必须是唯一的，但不同的类型可以使用相同的方法名。比如Point和Path类型的Distance方法。

使用方法的第一个好处：命名可以比函数更简单。

## 指针接收者的方法

motivation
    
    当调用一个函数时，会对其每一个实参变量进行拷贝，如果一个函数需要更新一个变量，或者如果一个实参太大而我们希望避免这种默认的拷贝。
    这种情况下我们必须使用指针来传递变量的地址。这也同样适用于更新接收者，比如*Point。
    
    func (p *Point) ScaleBy(factor float64){
        p.X *=factor
        p.Y *=factor
    }

这个方法的名字是(*Point).ScaleBy【因此在向别人描述代码的时候，需要带上接收者】。圆括号是必需的；没有圆括号，表达式可能被理解为*(Point.ScaleBy)。

在现实的程序中，一般会约定：如果Point这个类型有一个指针接收者，那么所有的Point的方法都必须有一个指针接收者，即使那些不需要这个指针接收者
的方法。

只有类型(Point)和指向它们的指针(*Point),才是可能会出现在接收者声明里的两种接收者（类型）。此外，为了避免歧义，在声明方法时，如果一个类型
名本身是一个指针的话，是不允许其出现在接收者中，比如下面这个例子。

    type P *int
    func (P) f(){/*.....*/}  //"invalid receiver type P (P is a pointer type)"
    
## 通过结构体内嵌组成类型（组合）
匿名字段类型可以是指向命名类型的指针，这个时候，字段和方法间接地来自于所指向的对象。

结构体类型可以拥有多个匿名字段。

## 方法值与方法表达式

方法表达式写成T.f或者(*T).f，其中T是类型，会返回一个函数值（函数变量）。这种函数会将第一个参数用作接收者。

    p:=Point{1,2}
    q:=Point{4,6}
    
    distance:=Point.Distance     //方法表达式
    fmt.Println(distance(p,q))   //"5"
    fmt.Println("%T\n",distance) //"func(Point,Point) float64"
    
    scale:=(*Point).ScaleBy
    scale(&p,2)
    fmt.Println(p)               //"{2 4}"
    fmt.Printf("%T\n",scale)     //"func(*Point,float64)"





