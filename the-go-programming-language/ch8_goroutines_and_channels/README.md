# 第8章 Goroutines和Channels

## 8.1 Goroutines
1. 在Go语言中，每一个并发的执行单元叫作`goroutine`。
2. 当一个程序启动时，其主函数会在一个单独的goroutine中运行，我们叫它`main goroutine`。
3. go语句会使其语句中的函数在一个新创建的goroutine中运行。而go语句本身会迅速地完成。
   ```go
   f()    // call f(); wait for it to return
   go f() // create a new goroutine that calls f(); don't wait
   ```
4. 主函数返回时，所有的goroutine都会被直接打断，程序退出。

## 8.2 示例：并发的Clock服务

## 8.3 示例：并发的Echo服务

## 8.4 Channels
   ```go
   // 创建channel
   ch = make(chan int)    // unbuffered channel, ch has type 'chan int'
   ch = make(chan int, 0) // unbuffered channel
   ch = make(chan int, 3) // buffered channel with capacity 3
   
   // 发送和接收
   ch <- x  // a send statement
   x = <-ch // a receive expression in an assignment statement
   <-ch     // a receive statement; result is discarded
   x, ok := <-naturals
   if !ok {
       break // channel was closed and drained
   }
   
   // 关闭channel
   close(ch)
   ```
1. `channels`是goroutine之间的通信机制，它可以让一个goroutine通过它给另一个goroutine发送值信息。
2. 每个channel都有一个特殊的类型，也就是channels可发送数据的类型。
3. 使用内置的make函数，我们可以创建一个channel。
4. 一个channel有发送和接受两个主要操作，且都使用<-运算符。
5. Channel还支持close操作，用于关闭channel，随后对基于该channel的任何发送操作都将导致panic，但对基于该channel的接收操作依然可以接受到之前已经成功发送的数据。
6. 一个基于无缓存Channels的发送操作将导致发送者goroutine阻塞，直到另一个goroutine在相同的Channels上执行接收操作，当发送的值通过Channels成功传输之后，两个goroutine可以继续执行后面的语句。
7. 基于无缓存Channels的发送和接收操作将导致两个goroutine做一次同步操作。
8. Channels也可以用于将多个goroutine连接在一起，一个Channel的输出作为下一个Channel的输入。这种串联的Channels就是所谓的`管道`（pipeline）。
9. 试图重复关闭一个channel将导致panic异常，试图关闭一个nil值的channel也将导致panic异常。
10. Go语言提供了单方向的channel类型，分别用于只发送或只接收的channel。类型`chan<- int`表示一个只发送int的channel，`<-chan int`表示一个只接收int的channel。
11. 只有在发送者所在的goroutine才能调用close函数，而对一个只接收的channel调用close将导致一个编译错误。
12. 向缓存Channel的发送操作就是向内部缓存队列的尾部插入元素，接收操作则是从队列的头部删除元素。

## 8.5 在并行中循环

## 8.6 示例：并发的Web服务

## 8.7 基于select的多路复用
1. `channel`的零值是nil。
2. 对一个值为nil的channel发送和接收操作会永远阻塞；在select语句中操作值为nil的channel永远都不会被select到。

## 8.8 示例：并发的目录遍历

## 8.9 退出

## 8.10 示例：聊天服务