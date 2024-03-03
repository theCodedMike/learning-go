# 第13章 底层编程

## 13.1 unsafe.Sizeof, Alignof和Offsetof
1. `unsafe.Sizeof`函数返回操作数在内存中的字节大小，参数可以是任意类型的表达式。
2. **内存空洞**是编译器自动添加的没有被使用的内存空间，用于保证后面每个字段或元素的地址相对于结构或数组的开始地址能够合理地对齐。
   ```text
            类型                            大小
   bool                               1个字节
   intN, uintN, floatN, complexN      N/8个字节（例如float64是8个字节）
   int, uint, uintptr                 1个机器字
   *T                                 1个机器字
   string                             2个机器字（data、len）
   []T                                3个机器字（data、len、cap）
   map                                1个机器字
   func                               1个机器字
   chan                               1个机器字
   interface                          2个机器字（type、value）
   
   var x struct {
       a bool
       b int16
       c []int
   }
   ```
   ![](../assets/holes_in_a_struct.png)
3. Go语言的规范并没有要求一个字段的声明顺序和内存中的顺序是一致的，所以理论上一个编译器可以随意地重新排列每个字段的内存位置。
4. `unsafe.Alignof`函数返回对应参数的类型需要对齐的倍数。
5. `unsafe.Offsetof`函数的参数必须是一个字段`x.f`，然后返回`f`字段相对于`x`起始地址的偏移量，包括可能的空洞。

## 13.2 unsafe.Pointer
1. `unsafe.Pointer`是一种特别定义的指针类型，它可以包含任意类型变量的地址。
2. 一个普通的`*T`类型指针可以被转化为`unsafe.Pointer`类型指针，并且一个`unsafe.Pointer`类型指针也可以被转回普通的指针，被转回的普通指针类型并不需要和原始的`*T`类型相同。
3. 一个`unsafe.Pointer`指针也可以被转化为`uintptr`类型，然后保存到指针型数值变量中。这种转换虽然也是可逆的，但是将`uintptr`转为`unsafe.Pointer`指针可能会破坏类型系统，因为并不是所有的数字都有有效的内存地址。
4. reflect包的`DeepEqual`函数可以对两个值进行深度相等判断，它可以工作在任意的类型上，甚至对于一些不支持==操作运算符的类型也可以工作。

## 13.3 示例：深度相等判断
1. 尽管`DeepEqual`函数很方便，而且可以支持任意的数据类型，但是它也有不足之处。例如，它将一个`nil`值的`map`和非`nil`值但是空的`map`视作不相等，同样`nil`值的`slice`和非`nil`但是空的`slice`也视作不相等。
   ```go
   var a, b []string = nil, []string{}
   fmt.Println(reflect.DeepEqual(a, b)) // "false"
   
   var c, d map[string]int = nil, make(map[string]int)
   fmt.Println(reflect.DeepEqual(c, d)) // "false"
   ```

## 13.4 通过cgo调用C代码

## 13.5 几点忠告
1. 谨慎使用`reflect`包和`unsafe`包。