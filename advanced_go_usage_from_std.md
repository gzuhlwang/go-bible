2020.02.16 by gzuhlwang

1、struct{}类型

在标准库中，可以随处看到struct{}的身影。例如context包的Context接口。

```
type Context interface{
     // omit some method list
     Done() <-chan struct{}
}
```

`struct{}`的实例是`struct{}{}`，该实例的数据类型是`struct{}`。

```
fmt.Printf("%T\n", struct {}{})  // "struct{}"
```

struct{}也就是人们常说的empty struct，即空结构。它占0字节的内存空间。

```
fmt.Printf("%d\n",unsafe.Sizeof(struct{}{})) // "0"
```

那么指向空结构体的指针变量占几个字节的内存空间呢？我们不妨测试一下。

```
fmt.Printf("%d\n",unsafe.Sizeof(&struct{}{})) // "8"
```

struct{}常和chan搭配使用，即大家熟悉的chan struct{}。chan是个引用类型。

关于struct{}的更多用法，可以学习文章[The empty struct](https://dave.cheney.net/2014/03/25/the-empty-struct)。



2、 信号量

信号量常用来限制最大的并发数量。semephore是荷兰计算机科学家Dijkstra在1965年提出的概念。目前，go的标准库没有信号量(semaphore)原语的内置实现。

很多go实现的信号量是基于带缓冲的通道。例如开源项目[Hyperledger Fabric](https://github.com/hyperledger/fabric)。

首先，抽象出了一个信号量接口，其定义如下：

```
type Semaphore interface{
	Acquire(ctx context.Context) error
	Release()
}
```

该接口的实现是

```
type Semaphore chan struct{}
```

这里又见struct{}类型。我们这里就不展开实现细节了。在"go内存模型"一文的[通道通信](https://golang.google.cn/ref/mem#tmp_7)部分有提到基于缓冲通道实现信号量。

除了信号量，goroutine pool（即常说的线程池）也是控制并发量的常用手段。

有关信号量和同步问题，可以阅读开源书籍[《The little Book of Semaphores》](http://greenteapress.com/semaphores/LittleBookOfSemaphores.pdf)。

3、6维

读，写，阻塞（通道）；非阻塞（通道）；同步；异步