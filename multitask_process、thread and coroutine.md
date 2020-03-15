> 以下内容选自《许式伟的架构课：多任务、线程与协程》

## 多任务(multitask)与执行体

多任务需求，既可以单核CPU系统上通过分时系统实现，也可以通过多核CPU，多核心运行不同程序实现。

执行体是指可被CPU赋予执行权的对象，它至少包含下一个执行位置（获取执行权后会从这里开始执行）以及其他的运行状态。



执行体的上下文，就是一堆寄存器的值。要切换执行体，只需要保存（保存针对的是正在运行的执行体）和恢复（恢复指的是）寄存器的值就好了。无论是进程、线程还是协程，都是如此。



协程切换并不需要系统调用。协程切换只是寄存器的值保存和恢复，所以可以在用户态下自己实现。

多线程用的是异步回调IO，例如libevent。

多协程用的是同步IO，例如Go。

## 进程 vs 线程

Q1：不用的执行体究竟有何不同？

Q2：为何会出现不同种类的执行体？

| 执行体 | 地址空间                 | 调度方       | 时间片调度               | 主动调度                      |
| ------ | ------------------------ | ------------ | ------------------------ | ----------------------------- |
| 进程   | 不同执行体有不同地址空间 | 操作系统内核 | 基于时钟中断（硬中断？） | 系统调用(syscall)（软中断？） |
| 线程   | 不同执行体共享地址空间   | 操作系统内核 | 基于时钟中断（硬中断？） | 系统调用(syscall)             |
| 协程   | 不同执行体共享地址空间   | 用户态       | 一般不支持               | 包装系统调用                  |

Q：Why processes？

进程是操作系统从安全性角度来说的隔离单位，不同进程之间基于最低授权的原则。

Q：Why thread？

线程的出现则是因为操作系统发现同一个软件内还是会有多任务的需求，这些任务处在相同的地址空间，彼此之间相互可以信任。

早起操作系统没有线程的概念，线程是80年代才引入的。



## 协程  vs goroutine

协程并不是操作系统内核提供的，它有时候也被称为用户态线程。这是因为协程是用户态下实现的。

Q:但为什么会出现协程呢？看起来它要应对的需求与线程一样，但是功能比线程弱很多？

答案是因为实现高性能的网络服务器的需求。对于常规的桌面应用程序来说，进程+线程绰绰有余。但对于一个网络服务器，我们可以用下面这个简单的模型看它：

【省略图】

对网络服务器来说，大量的来自客户端的请求包和服务器的返回包，都是网络IO；在响应请求的过程中，往往需要访问存储来保存和读取自身的状态，这也涉及本地网络IO。



如果这个网络服务器有很多客户，那么整个服务器就充斥着大量并行的IO请求。

操作系统提供的标准网络IO有以下这些成本：

- 系统调用(syscall)机制产生的开销；
- 数据多次拷贝的开销（数据总是先写到操作系统缓存再到用户存入的内存）；
- 因为没有数据而阻塞，（当解除阻塞时）产生调度重新获得执行权，产生的时间成本；
- 线程的空间成本和时间成本（标准IO请求都是同步调用，要想IO请求并行只能使用更多线程）。



系统调用虽然比函数调用多做了一点点事情，比如查询了中断向量表（这类似编程语言中的虚函数）。比如改变CPU的执行权限（从用户态跃迁到内核态再回到用户态）。



回过头来看，我们为什么希望减少线程数量？因为线程的成本高？我们分析一下。



首先，我们看下**时间成本**。它可以拆解为：

- 执行体切换本身的开销，它主要是寄存器保存和恢复的成本，可腾挪的余地非常有限【切换开销】；
- 执行体的调度开销，它主要是如何在大量已准备好的执行体中选出谁获得执行权【调度开销】；
- 执行体之间的同步与互斥成本。



我们再看线程的**空间成本**。它可以拆解为：

- 执行体的执行状态；
- TLS（thread local storage，线程局部存储）；
- 执行体的堆栈。

空间成本是第一根稻草。默认情况下Linux线程在数MB左右，其中最大的成本是堆栈（虽然线程的堆栈大小是可以设置的，但是出于线程执行安全性的考虑，线程的堆栈不能太小）。

我们可以算一下，如果一个线程1MB，那么有1000个线程就已经到GB级别了，消耗太快。



执行体的调度开销，以及执行体之间的同步与互斥成本，也是一个不可忽略的成本。虽然单位成本看起来还好，但是盖不住次数实在太多。



我们想象一下：系统中有大量的IO请求，大部分的IO请求并未命中而发生调度。另外，网络服务器的存储是个共享状态，也必然伴随着大量的同步与互斥操作。



综上，协程就是为了这样两个目标而来：

- 回归到同步IO的编程模式；
- 降低执行体的空间成本和时间成本。

一个成熟的协程库包括：

- 协程的创建
- (CPU)执行权的切换
- 协程的调度
- 协程的同步、互斥与通讯
- 协程的系统调用包装，尤其是网络IO请求的包装。

一个完备的协程库可以把它理解为用户态的操作系统，而协程就是用户态操作系统里面的“进程”。

Erlang和Go语言实现了完备的协程库。Erlang语言它基于虚拟机，但是道理上是一致的。Go语言里面的用户态“进程”叫goroutine。它（指goroutine）有这样一些重要设计：

- 堆栈开始很小（只有4K），但可按需自动增长；

- 坚决干掉了“线程局部存储（TLS）”特性的支持，让执行体更加精简；

- 提供了同步、互斥和其他常规执行体间的通讯手段，包括大家非常喜欢的channel；

- 提供了几乎所有重要的系统调用（尤其是IO请求）的包装。


## 阅读资料

- 腾讯开源的协程库[libco](https://github.com/Tencent/libco)

- [操作系统简史](http://www.personal.kent.edu/~rmuhamma/OpSystems/Myos/osHistory.htm)

- 异步回调IO网络库[libevent](https://github.com/libevent/libevent)

- [Why have processes？](https://web.cs.wpi.edu/~cs3013/c07/lectures/Section03-Processes.pdf) p.3

  1 resource sharing（logical （files） and physical(hardware)）

  2 computation speedup - taking advantage of multiprogramming — i.e. example of a customer/server database system

  补充：[Multitasking or Time Sharing System](http://www.idc-online.com/technical_references/pdfs/information_technology/Evolution_of_Operating_System_I.pdf):
  ● Multiprogramming **didn't provide the user interaction** with the computer system.
  ● Time sharing or Multitasking is a logical extension of Multiprogramming that **provides user**
  **interaction**.
  ● There are more than one user interacting the system at the same time
  ● The switching of CPU between two users is so fast that it gives the impression to user that he is
  only working on the system but actually it is shared among different users.
  ● CPU bound is divided into different time slots depending upon the number of users using the
  system.
  ● just as multiprogramming allows the processor to handle multiple batch jobs at a time,
  multiprogramming can also be used to handle multiple interactive jobs. In this latter case, the
  technique is referred to as time sharing, **because processor time is shared among multiple users**
  ● A multitasking system uses CPU scheduling and multiprogramming to provide each user with a
  small portion of a time shared computer. Each user has at least one separate program in
  memory.
  ● Multitasking are **more complex** than multiprogramming and must **provide a mechanism for jobs
  synchronization and communication** and it may ensure that system does not go in deadlock.

  3 modularity for protection

- [jobs vs processes](https://www.cl.cam.ac.uk/teaching/1011/OpSystems/os1a-slides.pdf) p.16

  On batch system, refer to *jobs*

  On interactive system, refer to *processes*

- 量化上下文切换的直接和间接成本。[Quantifying The Cost of Context Switch](https://www.usenix.org/legacy/events/expcs07/papers/2-li.pdf)

- 度量线程切换的直接成本。[Measuring context switching and memory overheads for Linux threads](https://eli.thegreenplace.net/2018/measuring-context-switching-and-memory-overheads-for-linux-threads/)

