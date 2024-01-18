# 第2章 程序结构

## 2.1 命名
1. Go语言中的函数名、变量名、常量名、类型名、语句标号和包名等所有的命名，必须以一个字母（Unicode字母）或下划线开头，后面可以跟任意数量的字母、数字或下划线
2. 大写字母和小写字母是不同的：heapSort和Heapsort是两个不同的名字
3. Go语言中的关键字共有25个；关键字不能用于自定义名字
   ```text
   <-关键字->
   break      default       func     interface   select
   case       defer         go       map         struct
   chan       else          goto     package     switch
   const      fallthrough   if       range       type
   continue   for           import   return      var
   
   <-30多个预定义的名字->
   内建常量: true false iota nil
   内建类型: int  int8  int16  int32  int64
           uint uint8 uint16 uint32 uint64 uintptr
           float32 float64 complex128 complex64
           bool byte rune string error
   内建函数: make len cap new append copy close delete
           complex real imag
           panic recover
   ```
4. 名字的开头字母的大小写决定了名字在包外的可见性。如果一个名字是大写字母开头的，那么它将是导出的
5. 在习惯上，Go语言推荐使用**驼峰式**命名。包本身的名字一般总是用小写字母

## 2.2 声明
1. Go语言主要有四种类型的声明语句：var、const、type和func，分别对应变量、常量、类型和函数实体对象的声明
2. 每个源文件以包的声明语句开始，说明该源文件是属于哪个包。包声明语句之后是import语句导入依赖的其它包，然后是包一级的类型、变量、常量、函数的声明语句

## 2.3 变量
1. 变量声明的一般语法如下，其中"类型"或"= 表达式"这两个部分可以省略其中的一个
   ```text
   var 变量名字 类型 = 表达式
   ```
2. 数值类型变量对应的零值是0，布尔类型变量对应的零值是false，字符串类型对应的零值是空字符串，接口或引用类型（包括slice、指针、map、chan和函数）变量对应的零值是nil
   ```text
   var s string                    // ""
   var i, j, k int                 // int, int, int
   var b, f, s = true, 2.3, "four" // bool, float64, string
   var f, err = os.Open(name)      // os.Open returns a file and an error
   i, j = j, i                     // 交换 i 和 j 的值
   ```
3. 在函数内部，有一种称为**短变量声明**（Short Variable Declaration）语句的形式可用于声明和初始化局部变量，它以`名字 := 表达式`形式声明变量，变量的类型根据表达式来自动推导
   ```text
   anim := gif.GIF{LoopCount: nframes}
   freq := rand.Float64() * 3.0
   t := 0.0
   i, j := 0, 1
   f, err := os.Open(name)
   ```
4. 一个指针的值是另一个变量的地址。一个指针对应变量在内存中的存储位置。并不是每一个值都会有一个内存地址，但是对于每一个变量必然有对应的内存地址
   ```text
   x := 1
   p := &x         // p, of type *int, points to x
   fmt.Println(*p) // "1"
   *p = 2          // equivalent to x = 2
   fmt.Println(x)  // "2"
   
   var p = f()     // 这样是可以的
   func f() *int {
       v := 1
       return &v
   }
   fmt.Println(f() == f()) // false，每次调用f函数都将返回不同的结果
   ```
5. 表达式`new(T)`将创建一个T类型的匿名变量，初始化为T类型的零值，然后返回变量地址，返回的指针类型为`*T`
   ```text
   p := new(int)   // p, *int 类型, 指向匿名的 int 变量
   fmt.Println(*p) // "0"
   *p = 2          // 设置 int 匿名变量的值为 2
   fmt.Println(*p) // "2"
   
   p := new(int)
   q := new(int)
   fmt.Println(p == q) // "false"
   ```
6. 变量的生命周期指的是在程序运行期间变量有效存在的时间间隔。对于在包一级声明的变量来说，它们的生命周期和整个程序的运行周期是一致的。而相比之下，局部变量的声明周期则是动态的

## 2.4 赋值
1. 使用赋值语句可以更新一个变量的值
   ```text
   x = 1                       // 命名变量的赋值
   *p = true                   // 通过指针间接赋值
   person.name = "bob"         // 结构体字段赋值
   count[x] = count[x] * scale // 数组、slice或map的元素赋值
   count[x] *= scale
   
   v := 1
   v++    // 等价方式 v = v + 1；v 变成 2
   v--    // 等价方式 v = v - 1；v 变成 1
   ```
2. 元组赋值是另一种形式的赋值语句，它允许同时更新多个变量的值
   ```text
   i, j, k = 2, 3, 5
   x, y = y, x
   a[i], a[j] = a[j], a[i]
   
   // 计算两个整数值的的最大公约数(GCD: Greatest Common Divisor)
   func gcd(x, y int) int {
       for y != 0 {
           x, y = y, x%y
       }
       return x
   }
   
   // 计算斐波纳契(Fibonacci)数列的第N个数
   func fib(n int) int {
       x, y := 0, 1
       for i := 0; i < n; i++ {
           x, y = y, x+y
       }
       return x
   }
   
   // function call returns two values
   f, err = os.Open("foo.txt")
   
   v = m[key]                // map查找，失败时返回零值
   v = x.(T)                 // type断言，失败时panic异常
   v = <-ch                  // 管道接收，失败时返回零值（阻塞不算是失败）
   
   v, ok = m[key]            // map返回2个值
   _, ok = mm[""], false     // map返回1个值
   _ = mm[""]                // map返回1个值，_表示丢弃该值
   ```
3. 隐式的赋值行为：函数调用会隐式地将调用参数的值赋值给函数的参数变量，一个返回语句会隐式地将返回操作的值赋值给结果变量等
   ```text
   medals := []string{"gold", "silver", "bronze"}
   ```
4. 对于两个值是否可以用`==`或`!=`进行相等比较的能力也和可赋值能力有关系：对于任何类型的值的相等比较，第二个值必须是对第一个值类型对应的变量是可赋值的，反之亦然

## 2.5 类型
1. 一个类型声明语句创建了一个新的类型名称(即给已有的类型**取别名**)
   ```text
   type 类型名字 底层类型
   ```

## 2.6 包和文件
1. 一个包的源代码保存在一个或多个以.go为文件后缀名的源文件中，通常一个包所在目录路径的后缀是包的导入路径，例如包`gopl.io/ch1/helloworld`对应的目录路径是`$GOPATH/src/gopl.io/ch1/helloworld`
2. 每个包都对应一个独立的名字空间，例如在`image`包中的`Decode`函数和在`unicode/utf16`包中的`Decode`函数是不同的
3. 在Go语言中，如果一个名字是大写字母开头的，那么该名字是导出的（汉字、日语等语言不区分大小写，故是不可导出的）
4. 每个源文件都是以包的声明语句开始，用来指明包的名字
5. 在每个源文件的包声明前紧跟着的注释是包注释，包注释的第一句应该先是包的功能概要说明
6. 包的初始化首先是解决包级变量的依赖顺序，然后按照包级变量声明出现的顺序依次初始化
   ```text
   var a = b + c // a 第三个初始化, 为 3
   var b = f()   // b 第二个初始化, 为 2, 通过调用 f (依赖c)
   var c = 1     // c 第一个初始化, 为 1
   
   func f() int { return c + 1 }
   ```
7. 每个文件都可以包含多个init初始化函数，这个特殊的init函数可以用来简化初始化工作。init初始化函数除了不能被调用或引用外，其他行为和普通函数类似
   ```text
   func init() { 
       /* ... */
   }
   ```
8. 初始化工作是自下而上进行的，main包最后被初始化

## 2.7 作用域
1. 声明语句的作用域是指源代码中可以有效使用这个名字的范围
2. 不要将作用域和生命周期混为一谈。声明语句的作用域对应的是一个源代码的文本区域，它是一个编译时的属性。一个变量的生命周期是指程序运行时变量存在的有效时间段，在此时间区域内它可以被程序的其他部分引用，是一个运行时的概念
3. **句法块**是由花括弧所包含的一系列语句，**词法块**是指在代码中并未显式地使用花括号包裹起来的声明
4. 一个程序可能包含多个同名的声明，只要它们在不同的词法域就没有关系
   ```text
   // 一个在函数体词法域，一个在for隐式的初始化词法域，一个在for循环体词法域
   func main() {
       x := "hello"
       for _, x := range x {
           x := x + 'A' - 'a'
           fmt.Printf("%c", x) // "HELLO" (one letter per iteration)
       }
   }
   ```