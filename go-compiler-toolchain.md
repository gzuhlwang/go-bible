## Two official Go compiler toolchains

- gc, Go compiler written in Go

- gccgo, a more traditional compiler using the GCC back end

reminder: "gc" stands for "Go compiler" while uppercase "GC", which stands for garbage collection. They have littile in common.

reminder: A **toolchain** is the **set of tools** that compiles source code into executables that can run on your target device, and includes **a compiler**,  **an assembler**, **a linker**, **runtime libraries**, and **a few useful utilities**.

我们可以在构建时指定编译器名字：

> go help build | grep compiler 

## Bootstrapping(自举)

The process of writing a compiler(or assembler) in the target programming language which it is intended to compile.

## Good Starting Point

[intro to the Go compiler](https://github.com/golang/go/tree/master/src/cmd/compile) [must read]

[Go 1.3+ Compiler Design Document](https://docs.google.com/document/d/1P3BLR31VA8cvLJLfMibSuTdwTuF7WWLux71CYD0eeD8/edit) by Russ Cox [must read]

[Go in Go](https://talks.golang.org/2015/gogo.slide#1) by Rob Pike [must read]

[gccgo](https://github.com/golang/gofrontend)

[Installing Go from source](https://golang.google.cn/doc/install/source)

[Bootstrapping Go 1.5](https://infoq.com/news/2015/01/golang-15-bootstrapped)