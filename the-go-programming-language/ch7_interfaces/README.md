# 第7章 接口

## 7.1 接口是合约
1. 接口类型是一种抽象的类型。

## 7.2 接口类型
1. 接口类型具体描述了一系列方法的集合，一个实现了这些方法的具体类型是这个接口类型的实例。
2. 新的接口类型可以通过组合已有的接口来定义。
   ```go
   package io
   type Reader interface {
       Read(p []byte) (n int, err error)
   }
   type Writer interface {
       Write(p []byte) (n int, err error)
   }
   type Closer interface {
       Close() error
   }
   
   // 接口内嵌 方式一
   type ReadWriter interface {
       Reader
       Writer
   }
   type ReadWriteCloser interface {
       Reader
       Writer
       Closer
   }
   
   // 接口内嵌 方式二
   type ReadWriter interface {
       Read(p []byte) (n int, err error)
       Write(p []byte) (n int, err error)
   }
   // 接口内嵌 方式三
   type ReadWriter interface {
       Read(p []byte) (n int, err error)
       Writer
   }
   ```

## 7.3 实现接口的条件
1. 一个类型如果拥有一个接口需要的所有方法，那么这个类型就实现了这个接口。
2. 一个类型属于某个接口只要这个类型实现了这个接口即可。
3. `interface{}`被称为空接口类型，对实现它的类型没有要求，所以可以将任意一个值赋给空接口类型。
   ```go
   //type any = interface{}
   var any interface{}
   any = true
   any = 12.34
   any = "hello"
   any = map[string]int{"one": 1}
   any = new(bytes.Buffer)
   ```

## 7.4 flag.Value接口
```go
package flag

type Value interface {
    String() string
    Set(string) error
}
```

## 7.5 接口值
1. **接口值**由两个部分组成，一个具体的类型和那个类型的值，它们被称为接口的`动态类型`和`动态值`。
2. 一个接口值是零值意味着它的类型和值的部分都是nil。
![](../assets/a_nil_interface_value.png)
3. 接口上的调用必须使用**动态分配**，因为不是直接调用，所以编译器必须把代码生成在类型描述符的方法上，然后间接调用那个地址，这个调用的接收者是一个接口动态值的拷贝。
![](../assets/an_interface_value.png)
4. 一个接口值可以持有任意大的动态值。
5. 接口值是可比较的，故它们可以用作map的键或者作为switch语句的操作数。
6. 一个不包含任何值的nil接口值和一个刚好包含nil指针的接口值是不同的。
![](../assets/a_non_nil_interface_containing_a_nil_ptr.png)

## 7.6 sort.Interface接口
```go
package sort

type Interface interface {
   Len() int
   Less(i, j int) bool
   Swap(i, j int)
}
```

## 7.7 http.Handler接口
```go
package http

type Header map[string][]string

type Handler interface {
   ServeHTTP(ResponseWriter, *Request)
}

type ResponseWriter interface {
   Header() Header
   Write([]byte) (int, error)
   WriteHeader(statusCode int)
}
```

## 7.8 error接口
```go
type error interface {
	Error() string
}
```
1. 创建一个error最简单的方法是调用errors.New函数（不推荐）。
2. 调用fmt.Errorf函数。

## 7.9 示例：表达式求值

## 7.10 类型断言

## 7.11 基于类型断言识别错误类型

## 7.12 通过类型断言查询接口

## 7.13 类型switch

## 7.14 示例：基于Token的XML解码

## 7.15 补充几点
