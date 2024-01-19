# 第4章 复合数据类型
我们主要讨论四种类型：数组、slice、map和结构体

## 4.1 数组
1. 数组是一个由固定长度的特定类型元素组成的序列，一个数组可以由零个或多个元素组成
2. 数组的长度必须是常量表达式，其长度需要在编译阶段确定。数组的长度是数组类型的一个组成部分
   ```go
   var a [3]int             // array of 3 integers
   fmt.Println(a[0])        // print the first element
   fmt.Println(a[len(a)-1]) // print the last element, a[2]
   
   // Print the indices and elements.
   for i, v := range a {
       fmt.Printf("%d %d\n", i, v)
   }
   
   // Print the elements only.
   for _, v := range a {
       fmt.Printf("%d\n", v)
   }
   
   var q [3]int = [3]int{1, 2, 3}
   var r [3]int = [3]int{1, 2}
   fmt.Println(r[2]) // "0"
   
   q := [...]int{1, 2, 3} // 数组的长度由初始化值的个数来推断
   fmt.Printf("%T\n", q)  // "[3]int"
   
   q := [3]int{1, 2, 3}
   q = [4]int{1, 2, 3, 4} // compile error: cannot assign [4]int to [3]int
   
   // 数组作为函数参数传递
   func zero(ptr *[32]byte) {
       for i := range ptr {
           ptr[i] = 0
       }
   }
   ```
3. 也可以指定一个索引和对应值列表的方式初始化数组
   ```go
   type Currency int
   const (
       USD Currency = iota // 美元
       EUR                 // 欧元
       GBP                 // 英镑
       RMB                 // 人民币
   )
   
   symbol := [...]string{USD: "$", EUR: "€", GBP: "￡", RMB: "￥"}
   fmt.Println(RMB, symbol[RMB]) // "3 ￥"
   ```
4. 和数组对应的类型是Slice（切片），它是可以增长和收缩的动态序列
5. 如果一个数组的元素类型是可以相互比较的，那么数组类型也是可以相互比较的
   ```go
   a := [2]int{1, 2}
   b := [...]int{1, 2}
   c := [2]int{1, 3}
   fmt.Println(a == b, a == c, b == c) // "true false false"
   d := [3]int{1, 2}
   fmt.Println(a == d) // compile error: cannot compare [2]int == [3]int
   ```

## 4.2 Slice
1. Slice（切片）代表变长的序列，序列中每个元素都有相同的类型。一般写作`[]T`，其中T代表slice中元素的类型；slice的语法和数组很像，只是没有固定长度而已
2. 一个slice由三个部分构成：指针、长度和容量，其中指针指向第一个slice元素对应的底层数组元素的地址，长度对应slice中元素的数目，容量一般是从slice的开始位置到底层数据的结尾位置
   ```text
   // 注意这里的[]里没有`...`，表明创建的是切片（动态数组）
   months := []string{1: "January", /* ... */, 12: "December"}
   Q2 := months[4:7]
   summer := months[6:9]
   fmt.Println(Q2)     // ["April" "May" "June"]
   fmt.Println(summer) // ["June" "July" "August"]
   ```
   ![](../assets/two_overlapping_slices_of_an_array_of_months.png)
3. 内置的len和cap函数分别返回slice的长度和容量
4. 和数组不同的是，slice之间不能比较（会报编译错误）。标准库提供了高度优化的bytes.Equal函数来判断两个字节型slice是否相等
5. slice唯一合法的比较操作是和nil比较
6. 一个零值的slice等于nil。一个nil值的slice并没有底层数组，此时底层数组的长度和容量都是0。但是也有非nil值的slice的长度和容量也是0的
   ```go
   var s []int    // len(s) == 0, s == nil
   s = nil        // len(s) == 0, s == nil
   s = []int(nil) // len(s) == 0, s == nil
   s = []int{}    // len(s) == 0, s != nil
   
   if summer == nil { /* ... */ }
   ```
7. 内置的make函数创建一个指定元素类型、长度和容量的slice
   ```go
   make([]T, len)      // len == cap
   make([]T, len, cap) // same as make([]T, cap)[:len]
   ```
8. 随着append，slice的扩容规律是：2倍扩容
9. 一个slice可以用来模拟一个stack
   ```go
   stack = append(stack, v) // push v
   top := stack[len(stack)-1] // top of stack
   stack = stack[:len(stack)-1] // pop
   ```

## 4.3 Map
1. 在Go语言中，一个`map`就是一个哈希表的引用，map类型可以写为`map[K]V`，其中K和V分别对应key和value
2. map中所有的key都有相同的类型，所有的value也有相同的类型，其中K对应的key必须是支持==比较运算符的数据类型
3. 虽然浮点数类型也是支持相等运算符比较的，但是将浮点数用做key类型则是一个坏的想法（key可能是NaN/Inf），而V对应的数据类型则没有任何的限制
4. 内置的make函数可以创建一个map
   ```go
   ages := make(map[string]int) // mapping from strings to ints
   
   ages := map[string]int{
       "alice":   31,
       "charlie": 34,
   }
   
   ages := map[string]int{} // 创建一个空的map
   ```
5. Map中的元素通过key对应的下标语法访问，即使这些元素不在map中也没有关系；如果查找失败将返回value类型对应的零值
   ```go
   ages["alice"] = 32         // 新增
   
   fmt.Println(ages["alice"]) // 查询
   if age, ok := ages["bob"]; !ok { 
       /* ... */
   }
   
   delete(ages, "alice")      // 删除
   
   for name, age := range ages { // 遍历
       fmt.Printf("%s\t%d\n", name, age)
   }
   for name := range ages { 
       // 也可以只遍历key
   }
   ```
6. 在向map存数据前必须先创建map，向一个nil值的map存入元素将导致一个panic异常
7. 和slice一样，map之间也不能进行相等比较；唯一的例外是和nil进行比较


## 4.4 结构体
1. 结构体是一种聚合的数据类型，是由零个或多个任意类型的值聚合成的实体。每个值称为结构体的成员。
   ```go
   type Employee struct {
       ID        int
       Name      string
       Address   string
       DoB       time.Time
       Position  string
       Salary    int
       ManagerID int
   }
   
   var dilbert Employee
   
   dilbert.Salary -= 5000 // demoted, for writing too few lines of code
   
   position := &dilbert.Position
   *position = "Senior " + *position // promoted, for outsourcing to Elbonia
   
   var employeeOfTheMonth *Employee = &dilbert
   employeeOfTheMonth.Position += " (proactive team player)"
   (*employeeOfTheMonth).Position += " (proactive team player)"
   
   func EmployeeByID(id int) *Employee { /* ... */ }
   fmt.Println(EmployeeByID(dilbert.ManagerID).Position) // "Pointy-haired boss"
   id := dilbert.ID
   EmployeeByID(id).Salary = 0 // fired for... no real reason
   ```
2. 结构体变量的成员可以通过点操作符访问
3. 点操作符也可以和指向结构体的指针一起工作
4. 结构体成员的先后顺序不同，会定义不同的结构体类型
5. 如果结构体成员名字是以大写字母开头的，那么该成员就是导出的
6. 一个命名为S的结构体类型将不能再包含S类型的成员，但是可以包含`*S`指针类型的成员
7. 如果结构体没有任何成员的话就是空结构体，写作struct{}，它的大小为0，也不包含任何信息
8. 结构体值也可以用结构体字面值表示，结构体字面值可以指定每个成员的值。
   ```go
   type Point struct{ X, Y int }
   
   p := Point{1, 2}
   anim := gif.GIF{LoopCount: nframes}
   
   func Scale(p Point, factor int) Point {
   return Point{p.X * factor, p.Y * factor}
   }
   fmt.Println(Scale(Point{1, 2}, 5)) // "{5 10}"
   
   func Bonus(e *Employee, percent int) int {
       return e.Salary * percent / 100
   }
   
   // 以下2种方式等价
   pp := &Point{1, 2} // 第1种
   pp := new(Point)   // 第2种
   *pp = Point{1, 2}
   ```
9. 结构体可以作为函数的参数和返回值。如果考虑效率的话，较大的结构体通常会以指针的方式传入和返回
10. 如果结构体的全部成员都是可以比较的，那么结构体也是可以比较的（即可以作为map的key）
11. 匿名成员：只声明一个成员对应的数据类型而不指定成员的名字
   ```go
   type Point struct {
       X, Y int
   }
   
   type Circle struct {
       Center Point
       Radius int
   }
   
   type Wheel struct {
       Circle Circle
       Spokes int
   }
   
   // 访问每个成员变得很繁琐
   var w Wheel
   w.Circle.Center.X = 8
   w.Circle.Center.Y = 8
   w.Circle.Radius = 5
   w.Spokes = 20
   
   // 引入匿名成员后就很方便了
   type Circle struct {
       Point
       Radius int
   }
   
   type Wheel struct {
       Circle
       Spokes int
   }
   
   var w Wheel
   w.X = 8            // equivalent to w.Circle.Point.X = 8
   w.Y = 8            // equivalent to w.Circle.Point.Y = 8
   w.Radius = 5       // equivalent to w.Circle.Radius = 5
   w.Spokes = 20
   
   ```
12. 因为匿名成员也有一个隐式的名字，因此不能同时包含两个类型相同的匿名成员，这会导致名字冲突

## 4.5 JSON
1. JSON：JavaScript对象表示法
2. Go语言对JSON、XML、ASN.1这些标准格式的编码和解码都有良好的支持，由标准库中的encoding/json、encoding/xml、encoding/asn1等包提供支持
3. 将一个Go语言中的结构体、slice转为JSON的过程叫**编组**（marshaling）
4. 编码的逆操作是**解码**（unmarshaling），对应将JSON数据解码为Go语言的数据结构

## 4.6 文本和HTML模板
1. 文本和HTML模板由text/template和html/template等模板包提供
2. 一个模板是一个字符串或一个文件，里面包含了一个或多个由双花括号包含的{{action}}对象