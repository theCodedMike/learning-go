# 第11章 测试

## 11.1 go test
1. `go test`命令是一个按照一定的约定和组织来测试代码的程序。
2. 在包目录内，所有以`_test.go`为后缀名的源文件在执行`go build`时不会被构建成包的一部分，它们是`go test`测试的一部分。
3. 在`*_test.go`文件中，有三种类型的函数：**测试函数**、**基准测试（benchmark）函数**、**示例函数**。
4. **测试函数**(test function)是以*Test为函数名前缀*的函数，用于测试程序的一些逻辑行为是否正确；go test命令会调用这些测试函数并报告测试结果是PASS或FAIL。
5. **基准函数**(benchmark function)是以*Benchmark为函数名前缀*的函数，它们用于衡量一些函数的性能；go test命令会多次运行基准测试函数以计算一个平均的执行时间。
6. **示例函数**(example function)是以*Example为函数名前缀*的函数，提供一个由编译器保证正确性的示例文档。
7. `go test`命令会遍历所有的`*_test.go`文件中符合上述命名规则的函数，生成一个临时的main包用于调用相应的测试函数，接着构建并运行、报告测试结果，最后清理测试中生成的临时文件。

## 11.2 测试函数
1. 参数`-v`可用于打印每个测试函数的名字和运行时间。
2. 参数`-run`对应一个正则表达式，只有测试函数名被它正确匹配的测试函数才会被`go test`测试命令运行。
3. **黑盒测试**只需要测试包公开的文档和API行为，内部实现对测试代码是透明的。
4. **白盒测试**有访问包内部函数和数据结构的权限，因此可以做到一些普通客户端无法实现的测试。
5. 使用**外部测试包**的方式解决循环依赖的问题。
![](../assets/import_cycles.png)
![](../assets/external_test_packages.png)
6. 可以用`go list`命令查看包对应目录中哪些Go源文件是产品代码，哪些是包内测试，还有哪些是外部测试包。
   ```text
   go list -f={{.GoFiles}} fmt       // fmt包的go源文件, 也就是go build命令要编译的部分
   go list -f={{.TestGoFiles}} fmt   // fmt包的内部测试代码，以_test.go为后缀文件名
   go list -f={{.XTestGoFiles}} fmt  // fmt包的外部测试包，也就是fmt_test包
   ```
7. 如果一个测试仅仅对程序做了微小变化就失败则称为**脆弱**。

## 11.3 测试覆盖率
```text
// 可以测试覆盖率
go test -coverprofile=c.out ./ch11_testing/11_2_test_functions/word2
// 以html打开c.out文件
go tool cover -html=c.out
```
1. 对待测程序执行的测试的程度称为**测试的覆盖率**。

## 11.4 基准测试
   ```text
   // 执行基准测试
   go test -bench=BenchmarkIsPalindrome ./ch11_testing/11_2_test_functions/word2
   // 报告内存分配
   go test -bench=BenchmarkIsPalindrome -benchmem ./ch11_testing/11_2_test_functions/word2
   ```
1. **基准测试**是测量一个程序在固定工作负载下的性能。
2. `-benchmem`命令行参数会在报告中指出内存的分配数据统计。

## 11.5 剖析
   ```text
   // CPU剖析数据标识了最耗CPU时间的函数
   go test -cpuprofile=cpu.out
   // 阻塞剖析则记录阻塞goroutine最久的操作
   go test -blockprofile=block.out
   // 堆剖析则标识了最耗内存的语句
   go test -memprofile=mem.out
   
   // 1. 收集数据
   go test -run=NONE -bench=Benchmark -cpuprofile=cpu.log ./ch11_testing/11_2_test_functions/word2
   // 2. 展示数据
   go tool pprof -text -nodecount=10 ./word2.test cpu.log
   ```

## 11.6 示例函数
1. **示例函数**有三个用处，其中最主要的一个是作为文档。一个示例函数可以方便地展示属于同一个接口的几种类型或函数之间的关系。