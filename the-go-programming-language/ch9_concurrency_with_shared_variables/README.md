# 第9章 基于共享变量的并发

## 9.1 竞争条件
1. 一个函数在并发调用时没法工作，其原因可能有死锁（deadlock）、活锁（livelock）和饿死（resource starvation）等。
2. `竞争条件`是指程序在多个goroutine交叉执行操作时，没有给出正确的结果。
3. `数据竞争`是在两个以上的goroutine并发访问相同的变量且其中至少一个为写操作时发生。
4. 有三种方式可以避免数据竞争：
   ```go
   // 1. 避免写操作
   var icons = map[string]image.Image{
   "spades.png":   loadIcon("spades.png"),
   "hearts.png":   loadIcon("hearts.png"),
   "diamonds.png": loadIcon("diamonds.png"),
   "clubs.png":    loadIcon("clubs.png"),
   }
   
   // Concurrency-safe.
   func Icon(name string) image.Image { return icons[name] }
   
   
   // 2. 避免从多个goroutine访问变量
   
   // 3. 允许很多goroutine去访问变量，但是在同一个时刻最多只有一个goroutine在访问，这种方式被称为“互斥”
   ```

## 9.2 sync.Mutex互斥锁
1. 一个只能为1和0的信号量叫做`二元信号量`（binary semaphore）。
2. 在Lock和Unlock之间的代码段中的内容goroutine可以随便读取或者修改，这个代码段叫做`临界区`。
3. goroutine在结束后释放锁是必要的，无论以哪条路径都需要释放，即使在错误路径中也要释放。

## 9.3 sync.RWMutex读写锁
1. 允许多个只读操作并行执行，但写操作会完全互斥，这种锁叫作`多读单写`锁（multiple readers, single writer lock）

## 9.4 内存同步

## 9.5 sync.Once惰性初始化

## 9.6 竞争条件检测
1. Go的runtime和工具链为我们装备了一个复杂但好用的动态分析工具——竞争检查器（the race detector）。
2. 只要在`go build`，`go run`或者`go test`命令后面加上`-race`flag，就会使编译器创建一个能够记录所有运行期对共享变量访问工具的test，并且会记录下每一个读或者写共享变量的goroutine的身份信息。

## 9.7 示例：并发的非阻塞缓存

## 9.8 Goroutines和线程
1. 每一个OS线程都有一个固定大小的内存块（一般会是2MB）来做栈，这个栈会用来存储当前正在被调用或挂起（指在调用其它函数时）的函数的内部变量。
2. 一个goroutine会以一个很小的栈开始其生命周期，一般只需要2KB（可以动态伸缩，最大为1GB），会保存其活跃或挂起的函数调用的本地变量。
3. OS线程会被操作系统内核调度。而Go调度器并不是一个硬件定时器，而是被Go语言本身进行调度的。
4. Go的调度器使用了一个叫做`GOMAXPROCS`的变量来决定会有多少个OS线程同时执行Go的代码，其默认值是运行机器上的CPU的核心数。
5. goroutine没有ID。