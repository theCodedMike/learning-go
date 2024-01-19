# 第5章 函数
函数可以让我们将一个语句序列打包为一个单元，然后可以从程序中其它地方多次调用。
函数的机制可以让我们将一个大的工作分解为小的任务，这样的小任务可以让不同程序员在不同时间、不同地方独立完成。
一个函数同时对用户隐藏了其实现细节。由于这些因素，对于任何编程语言来说，函数都是一个至关重要的部分。

## 5.1 函数声明
1. 函数声明包括函数名、形式参数列表、返回值列表（可省略）以及函数体。
   ```go
   func name(parameter-list) (result-list) {
       body
   }
   
   func hypot(x, y float64) float64 {
       return math.Sqrt(x*x + y*y)
   }
   fmt.Println(hypot(3,4)) // "5"
   
   func f(i, j, k int, s, t string)                 { /* ... */ }
   func f(i int, j int, k int,  s string, t string) { /* ... */ }
   
   func add(x int, y int) int   {return x + y}
   func sub(x, y int) (z int)   { z = x - y; return}
   func first(x int, _ int) int { return x }
   func zero(int, int) int      { return 0 }
   fmt.Printf("%T\n", add)   // "func(int, int) int"
   fmt.Printf("%T\n", sub)   // "func(int, int) int"
   fmt.Printf("%T\n", first) // "func(int, int) int"
   fmt.Printf("%T\n", zero)  // "func(int, int) int"
   ```
2. 如果一个函数在声明时，包含返回值列表，该函数必须以 return语句结尾，除非函数明显无法运行到结尾处。
3. 函数的类型被称为函数的签名。如果两个函数形式参数列表和返回值列表中的变量类型一一对应，那么这两个函数被认为有相同的类型或签名。
4. 在函数体中，函数的形参作为局部变量，被初始化为调用者提供的值。函数的形参和有名返回值作为函数最外层的局部变量，被存储在相同的词法块中。

## 5.2 递归
1. 函数可以是递归的，这意味着函数可以直接或间接的调用自身。
2. 
## 5.3 多返回值

## 5.4 错误

## 5.5 函数值

## 5.6 匿名函数

## 5.7 可变参数

## 5.8 Deferred函数

## 5.9 Panic异常

## 5.10 Recover捕获异常
