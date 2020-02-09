# Communicating Sequential Goroutines in Go

by gzuhlwang

## History:

- 2020-02-01: draft v0.0.1
- 2020-02-05：更新goroutine状态；增加goroutine可能切换的点
- 2020-02-09：在take away部分更新Dave Sheney关于goroutine的论述；在参考资料部分添加[panjf2000](https://gocn.vip/people/4028)的文章

## 0x00 Computer Science 101

### Concepts

> computer science is about concepts, not languages.
>
> ​                                                                      《Computation and State Machines》 by Leslie Lamport 

- binary(or program)

  passive entity(静态实体),lifeless thing, sits on the disk, just assembly language(instruction plus some static data) ,also refer to program, large and significant binaries are called applications.

- process

  active entity,running program,smallest unit of resource allocation(CPU time,(virtualized) memory,file,I/O devices,thread(s)).

- thread

  smallest unit of execution/activity inside of a process, schedulable by an OS’s process scheduler.

- coroutine

  协程, see wikipedia

  two categories:

  → stackless

  → stackful(e.g. goroutine in go, fiber in crystal)

- fiber

  纤程，see wikipedia

  Crystal programming language has [fibers](https://crystal-lang.org/reference/guides/concurrency.html).

- concurrency vs parallelism

  Concurrency is about dealing with(应对) lots of things at once.

  Parallelism is about doing(做) lots of things at once.

  Not the same, but related.

  Concurrency is about structure, parallelism is about execution.

  Concurrency provides a way to structure a solution to solve a problem that may (but not necessarily) be parallelizable.

  For difference between concurrency and parallelism, see [stackoverflow](https://stackoverflow.com/questions/1050222/what-is-the-difference-between-concurrency-and-parallelism/1050257#1050257). For more on concepts above, Robert Love's book 《Linux System Programming》 and Rob Pike's [talk](For more, see https://talks.golang.org/2012/waza.slide#1) 

> linux只支持进程和线程。Windows还提供了Fiber的系统调用。crystal在语言层面实现了fiber。二者不太一样。协程分为两类。goroutine是stackful.

## More on Concurrency vs parallelism

1、What's the difference between concurrency and parallelism?

Explain it to a five year old.

Concurrent

Two queues and one coffee machine(咖啡机).

Parallel

Two queues and two coffee machines.

> from [Joe Armstrong](https://joearms.github.io/#2013-04-05%20Concurrent%20and%20Parallel%20Programming), Erlang's author

2、book 《The Crystal Programming Language》

“A concurrent system is one that can be in charge of (负责) many tasks, although not necessarily it is executing them at the same time. You can think of yourself being in the kitchen cooking: you chop an onion, put it to fry, and while it's being fried you chop a tomato, but you are not doing all of those things at the same time: you distribute your time between those tasks. Parallelism would be to stir fry onions with one hand while with the other one you chop a tomato.”

### Thread Model

- 1:1
- N:1
- N:M(hybrid)

> Go multiplexs(or schedules) N goroutines on M OS thread(called goroutines in go).
>
> For more on thread model, see book 《Linux System Programming》.

### 0x02 The Go Programming Language

### Into to Go

From 《go spec》, "Go is a general-purpose language designed with systems programming in mind. It is strongly typed and garbage-collected and has explicit support for concurrent programming."

> 公认的系统编程语言：C、C++、Rust

### Some Observations

1 From 《Introduction to Concurrent Programming》By Rob Pike，2000

"The world runs in parallel, but our usual model of software does not. Programming languages are sequential. This mismatch makes it hard to write systems software that provides the interface between a computer (or user) and the world."

世界并行运作，但通常我们的软件模型不是。编程语言都是串行的。这种错位使得编写系统软件变得困难。

For more, see http://herpolhode.com/rob/lec1.pdf

2 From 《Concurrency Oriented Programming in Erlang》by   Joe Armstrong，2003

"the real world, the world in which we live and breath and are born in and die is concurrent. Paradoxically, the programming languages which we use to write programs which interact with the real world are predominately sequential."   ​                                                                                              

For more, see http://citeseerx.ist.psu.edu/viewdoc/download?doi=10.1.1.116.1969&rep=rep1&type=pdf

3 《The Free Lunch Is Over: A Fundamental Turn Toward Concurrency in Software》by Herb Sutter, 2005

“We desperately need a higher-level programming model for concurrency than languages offer today; I'll have more to say about that soon.”

 For more, see https://www.cs.utexas.edu/~lin/cs380p/Free_Lunch.pdf

>   《免费午餐终结：软件并发的根本转向》，作者是ISO C++标准委员会主席，C++/CLI首席架构师。作者在文章提到主流的并发模型是lock-based programming.

### The Origins of Go Concurrency
picture omitted

see <http://go-lang.cat-v.org/talks/slides/emerging-languages-camp-2010.pdf> and  https://swtch.com/~rsc/thread/#8 for more 

## Two Models of Concurrency

1. Shared Memory

   - lock(mutex)

   Representative language: Java, Python,...

2. Message Passing(eg. CSP, Actor Model)

   Representative language: Erlang, go,...

> Go supports both traditional shared memory and  message passing, specifically csp.
>
> Whether communication by sharing memory or share memory by communicating in Go depends on concret case.

 ### Two threads(called goroutines in Go) sharing the use of the same object

Thread 1(Goroutine 1)

picture omitted

> go provides concurrency primitives at language level.
>
> ➔channel(messaging)
>
> →goroutine(the unit of concurrecy in Go)

### Why Called Goroutine?

"They're called goroutines because the existing terms—threads, coroutines, processes, and so on—convey inaccurate connotations."        —— 《Effective Go》

> https://golang.org/doc/effective_go.html#goroutines

### What goroutine?
"A goroutine is a lightweight thread managed by the Go runtime."   —— 《A Tour of Go》

strictly, A goroutine is a lightweight user level thread managed by the Go runtime.

>  https://golang.org/doc/effective_go.html#goroutines

### The State of Goroutine

- Creation

  using go keyword,e.g. go func()

- Runnable

  goready

- Running

- Exit

  goexit

- Waiting

  gopark



  ps:goroutine可能切换的点：

  1、I/O(net io,部分file io)

  2、显式地调用runtime.Gosched()让出资源

  3、channel（阻塞写或读通道的goroutine）

  4、select（默认阻塞）

  5、一个goroutine拿到锁，另一个goroutine会等待锁

  上面只是参考，不保证切换，不保证在其他地方不切换goroutine。

### Milestones of Go(routine) Scheduler

- go 1.0:G-M model
- go 1.1: G-P-M model，work-stealing algorithm
- go 1.2: 新增了抢占式调度（部分解决了“饿死”问题）
- go 1.4:much of the  runtime code has been translated to Go
- go 1.5:Go compiler and runtime entirely written in go,the runtime  now sets the default number of threads to run simultaneously, defined by GOMAXPROCS, to the number of cores available on the CPU
- ...

## Take Away

- Go is a concurrency oriented language.

- Go supports both communication by sharing memory and  shared memory by communicating.

- Concurrency is not parallelism, it's better.

- Never start a goroutine without knowing when it will stop. —— Dave Cheney

  （goroutine会占用资源：锁，变量，内存。释放这些资源可靠方法是停止相应的goroutine。）

## Reference

<http://courses.cs.vt.edu/cs5204/fall09-kafura/Presentations/CSP.pdf>

<http://hjemmesider.diku.dk/~vinter/xmp/lecture3.pdf>

<https://morsmachine.dk/go-scheduler>

http://web.mit.edu/6.031/www/sp17/classes/19-concurrency/#message_passing_example

[http://rtoal.github.io/csp-talk/#/](http://rtoal.github.io/csp-talk/) 【Must Read】
[https://swtch.com/~rsc/thread/#8](https://swtch.com/~rsc/thread/) 【Must Read】

<http://go-lang.cat-v.org/talks/slides/emerging-languages-camp-2010.pdf> 【Must Watch】

Hoare's CSP paper 【Must Read】

Operating Systems: Three Easy Pieces

[Goroutine并发调度模型深度解析之手撸一个协程池](https://segmentfault.com/a/1190000015464889):作者回顾了3大线程模型，Go调度模型G-M-P，是对本slide的细化。



