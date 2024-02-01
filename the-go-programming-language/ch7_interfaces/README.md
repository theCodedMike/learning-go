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

## 7.4 flag.Value接口

## 7.5 接口值

## 7.6 sort.Interface接口

## 7.7 http.Handler接口

## 7.8 error接口

## 7.9 示例：表达式求值

## 7.10 类型断言

## 7.11 基于类型断言识别错误类型

## 7.12 通过类型断言查询接口

## 7.13 类型switch

## 7.14 示例：基于Token的XML解码

## 7.15 补充几点
