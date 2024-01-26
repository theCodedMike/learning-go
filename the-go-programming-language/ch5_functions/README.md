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
2. 如果一个函数在声明时，包含返回值列表，该函数必须以`return`语句结尾，除非函数明显无法运行到结尾处。
3. 函数的类型被称为函数的签名。如果两个函数形式参数列表和返回值列表中的变量类型一一对应，那么这两个函数被认为有相同的类型或签名。
4. 在函数体中，函数的形参作为局部变量，被初始化为调用者提供的值。函数的形参和有名返回值作为函数最外层的局部变量，被存储在相同的词法块中。

## 5.2 递归
1. 函数可以是递归的，这意味着函数可以直接或间接地调用自身。

## 5.3 多返回值
1. 在Go中，一个函数可以返回多个值。
   ```go
   links, err := findLinks(url)
   links, _ := findLinks(url) // errors ignored
   
   func findLinksLog(url string) ([]string, error) {
       log.Printf("findLinks %s", url)
       return findLinks(url)
   }
   
   func Size(rect image.Rectangle) (width, height int)
   func Split(path string) (dir, file string)
   func HourMinSec(t time.Time) (hour, minute, second int)
   
   // bare return
   func CountWordsAndImages(url string) (words, images int, err error) {
       resp, err := http.Get(url)
       if err != nil {
           return
       }
       doc, err := html.Parse(resp.Body)
       resp.Body.Close()
       if err != nil {
           err = fmt.Errorf("parsing HTML: %s", err)
           return
       }
       words, images = countWordsAndImages(doc)
       return
   }
   ```
2. 调用多返回值函数时，返回给调用者的是一组值，调用者必须显式的将这些值分配给变量。如果某个值不被使用，可以将其分配给blank identifier。
3. 如果一个函数所有的返回值都有显式的变量名，那么该函数的return语句可以省略操作数。这称之为`bare return`

## 5.4 错误
1. 内置的error是接口类型
2. io包保证任何由文件结束引起的读取失败都返回同一个错误`io.EOF`
3. 错误处理策略
   ```go
   // 1. 传播错误
   resp, err := http.Get(url)
   if err != nil{
       return nil, err
   }
   
   // 2. 如果错误的发生是偶然性的，或由不可预知的问题导致的，可以尝试重试
   // WaitForServer attempts to contact the server of a URL.
   // It tries for one minute using exponential back-off.
   // It reports an error if all attempts fail.
   func WaitForServer(url string) error {
       const timeout = 1 * time.Minute
       deadline := time.Now().Add(timeout)
       for tries := 0; time.Now().Before(deadline); tries++ {
           _, err := http.Head(url)
           if err == nil {
               return nil // success
           }
           log.Printf("server not responding (%s);retrying…", err)
           time.Sleep(time.Second << uint(tries)) // exponential back-off
       }
       return fmt.Errorf("server %s failed to respond after %s", url, timeout)
   }
   
   // 3. 如果错误发生后，程序无法继续运行，则输出错误信息并结束程序。这种策略只应在main中执行
   // (In function main.)
   if err := WaitForServer(url); err != nil {
       fmt.Fprintf(os.Stderr, "Site is down: %v\n", err)
       os.Exit(1)
   }
   
   // 4. 有时我们只需要输出错误信息就足够了，不需要中断程序的运行
   if err := Ping(); err != nil {
       log.Printf("ping failed: %v; networking disabled",err)
   }
   
   // 5. 直接忽略掉错误
   dir, err := ioutil.TempDir("", "scratch")
   if err != nil {
       return fmt.Errorf("failed to create temp dir: %v",err)
   }
   // ...use temp dir, ignore errors; $TMPDIR is cleaned periodically
   os.RemoveAll(dir) // 操作系统会定期地清理临时目录，故无须处理这里出现的错误
   ```

## 5.5 函数值
1. 在Go中，函数被看作第一类值（first-class values）：函数像其他值一样，拥有类型，可以被赋值给其他变量，传递给函数，从函数返回。对函数值（function value）的调用类似函数调用。
   ```go
   func square(n int) int { return n * n }
   func negative(n int) int { return -n }
   func product(m, n int) int { return m * n }
   
   f := square
   fmt.Println(f(3)) // "9"
   
   f = negative
   fmt.Println(f(3))     // "-3"
   fmt.Printf("%T\n", f) // "func(int) int"
   
   f = product // compile error: can't assign func(int, int) int to func(int) int
   ```
2. 函数类型的零值是nil。调用值为nil的函数值会引起panic错误。
3. 函数值可以与nil比较，但是函数值之间是不可比较的，也不能用函数值作为map的key。

## 5.6 匿名函数
1. 拥有函数名的函数只能在包级语法块中被声明
2. 函数字面量（function literal）是一种表达式，它的值被称为**匿名函数**（anonymous function）
3. Go使用**闭包**（closures）技术实现函数值
4. 当匿名函数需要被递归调用时，需要先声明一个变量，然后再将匿名函数赋值给这个变量。即必须分2步
5. 函数值中记录的是循环变量的内存地址，而不是循环变量某一时刻的值。
   ```go
   var rmdirs []func()
   for _, dir := range tempDirs() {
       //os.MkdirAll(dir, 0755) // NOTE: incorrect!
       dir := dir // declares inner dir, initialized to outer dir
       rmdirs = append(rmdirs, func() {
           os.RemoveAll(dir)
       })
   }
   // ...do some work…
   for _, rmdir := range rmdirs {
       rmdir() // clean up
   }
   ```

## 5.7 可变参数
1. 参数数量可变的函数称为可变参数函数。
2. 在声明可变参数函数时，需要在参数列表的最后一个参数类型之前加上省略符号"..."，这表示该函数会接收任意数量的该类型参数。
3. 可变参数函数和以切片作为参数的函数是不同的，它们只是行为上看起来很像。
   ```go
   func sum(vals ...int) int {
       total := 0
       for _, val := range vals {
           total += val
       }
       return total
   }
   fmt.Println(sum())           // 0
   fmt.Println(sum(3))          // 3
   fmt.Println(sum(1, 2, 3, 4)) // 10
   values := []int{1, 2, 3, 4}
   fmt.Println(sum(values...))  // 10
   
   func f(...int) {}
   func g([]int) {}
   fmt.Printf("%T\n", f) // func(...int)
   fmt.Printf("%T\n", g) // func([]int)
   ```
4. 可变参数函数经常被用于格式化字符串。

## 5.8 Deferred函数

## 5.9 Panic异常

## 5.10 Recover捕获异常
