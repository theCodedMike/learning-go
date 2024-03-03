# 第12章 反射
Go语言提供了一种机制，能够在运行时更新变量和检查它们的值、调用它们的方法和它们支持的内在操作，而不需要在编译时就知道这些变量的具体类型。这种机制被称为**反射**。

## 12.1 为何需要反射？

## 12.2 reflect.Type和reflect.Value
1. **反射**是由`reflect`包提供的。它定义了两个重要的类型——`Type`和`Value`。一个`Type`表示一个Go类型。
2. 函数`reflect.TypeOf`接受任意的`interface{}`类型，并以`reflect.Type`形式返回其动态类型。
   ```go
   t := reflect.TypeOf(3)  // a reflect.Type
   fmt.Println(t.String()) // "int"
   fmt.Println(t)          // "int"
   
   fmt.Printf("%T\n", 3) // "int"
   
   v := reflect.ValueOf(3) // a reflect.Value
   fmt.Println(v)          // "3"
   fmt.Printf("%v\n", v)   // "3"
   fmt.Println(v.String()) // NOTE: "<int Value>"
   ```
3. 函数`reflect.ValueOf`接受任意的`interface{}`类型，并返回一个装载着其动态值的`reflect.Value`。
4. `reflect.Value`和`interface{}`都能装载任意的值。所不同的是，一个空的接口隐藏了值内部的表示方式和所有方法。

## 12.3 Display递归打印

## 12.4 示例：编码S表达式

## 12.5 通过reflect.Value修改值
   ```go
   x := 2                   // value   type    variable?
   a := reflect.ValueOf(2)  // 2       int     no
   b := reflect.ValueOf(x)  // 2       int     no
   c := reflect.ValueOf(&x) // &x      *int    no
   d := c.Elem()            // 2       int     yes (x)
   
   fmt.Println(a.CanAddr()) // "false"
   fmt.Println(b.CanAddr()) // "false"
   fmt.Println(c.CanAddr()) // "false"
   fmt.Println(d.CanAddr()) // "true"
   
   x := 2
   d := reflect.ValueOf(&x).Elem()   // d refers to the variable x
   px := d.Addr().Interface().(*int) // px := &x 先获取指针
   *px = 3                           // x = 3    再更新值
   fmt.Println(x)                    // "3"
   
   d.Set(reflect.ValueOf(4))         // 可以不获取指针就直接set
   fmt.Println(x)                    // "4"
   
   x := 1
   rx := reflect.ValueOf(&x).Elem()
   rx.SetInt(2)                     // OK, x = 2
   rx.Set(reflect.ValueOf(3))       // OK, x = 3
   rx.SetString("hello")            // panic: string is not assignable to int
   rx.Set(reflect.ValueOf("hello")) // panic: string is not assignable to int
   
   var y interface{}
   ry := reflect.ValueOf(&y).Elem()
   ry.SetInt(2)                     // panic: SetInt called on interface Value
   ry.Set(reflect.ValueOf(3))       // OK, y = int(3)
   ry.SetString("hello")            // panic: SetString called on interface Value
   ry.Set(reflect.ValueOf("hello")) // OK, y = "hello"
   
   stdout := reflect.ValueOf(os.Stdout).Elem() // *os.Stdout, an os.File var
   fmt.Println(stdout.Type())                  // "os.File"
   fd := stdout.FieldByName("fd")
   fmt.Println(fd.Int())                       // "1"
   fd.SetInt(2)                                // panic: unexported field
   fmt.Println(fd.CanAddr(), fd.CanSet())      // "true false"
   ```
1. 可以通过调用`reflect.ValueOf(&x).Elem()`来获取任意变量x对应的可取地址的Value。
2. 要从变量对应的可取地址的reflect.Value来访问变量需要三个步骤。1：调用Addr()方法，2：在Value上调用Interface()方法，3：将得到的interface{}类型的接口强制转为普通的类型指针。
3. 反射可以越过Go语言的导出规则的限制，从而读取结构体中未导出的成员，但并不能修改这些未导出的成员。

## 12.6 示例：解码S表达式

## 12.7 获取结构体字段标签

## 12.8 显示一个类型的方法集

## 12.9 几点忠告
1. 避免因反射而导致的脆弱性问题，其最好的解决方法是将所有的反射相关的使用控制在包的内部。
2. 基于反射的代码通常比正常的代码运行速度慢一到两个数量级。
3. 当反射能使程序更加清晰的时候可以考虑使用。测试是一个特别适合使用反射的场景。