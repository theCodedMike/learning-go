# The Go Programming Language

## [Chapter 1: Tutorial 入门](ch1_tutorial/README.md)
### 1.1 Hello, World
### 1.2 Command Line Arguments 命令行参数
### 1.3 Finding Duplicate Lines 查找重复的行
### 1.4 Animated GIGs GIF动画
### 1.5 Fetching a URL 获取URL
### 1.6 Fetching URLs Concurrently 并发获取多个URL
### 1.7 A Web Server Web服务
### 1.8 Loose Ends 本章要点

## [Chapter 2: Program Structure 程序结构](./ch2_program_structure/README.md)
### 2.1 Names 命名
### 2.2 Declarations 声明
### 2.3 Variables 变量
### 2.4 Assignments 赋值
### 2.5 Type Declarations 类型
### 2.6 Packages and Files 包和文件
### 2.7 Scope 作用域

## [Chapter 3: Basic Data Types 基础数据类型](ch3_basic_data_types/README.md)
### 3.1 Integers
### 3.2 Floating-Point Numbers
### 3.3 Complex Numbers
### 3.4 Booleans
### 3.5 Strings
### 3.6 Constants

## [Chapter 4: Composite Types 复合数据类型](./ch4_composite_types/README.md)
### 4.1 Arrays
### 4.2 Slices
### 4.3 Maps
### 4.4 Structs
### 4.5 JSON
### 4.6 Text and HTML Templates

## [Chapter 5: Functions 函数](./ch5_functions/README.md)
### 5.1 Function Declarations
### 5.2 Recursion
### 5.3 Multiple Return Values
### 5.4 Errors
### 5.5 Function Values
### 5.6 Anonymous Functions
### 5.7 Variadic Functions
### 5.8 Deferred Function Calls
### 5.9 Panic
### 5.10 Recover

## [Chapter 6: Methods 方法](./ch6_methods/README.md)
### 6.1 Method Declarations
### 6.2 Methods with a Pointer Receiver
### 6.3 Composing Types by Struct Embedding
### 6.4 Method Values and Expressions
### 6.5 Example: Bit Vector Type
### 6.6 Encapsulation

## [Chapter 7: Interfaces 接口](./ch7_interfaces/README.md)
### 7.1 Interfaces as Contracts
### 7.2 Interface Types
### 7.3 Interface Satisfaction
### 7.4 Parsing Flags with *flag.Value*
### 7.5 Interface Values
### 7.6 Sorting with *sort.Interface*
### 7.7 The *http.Handler* Interface
### 7.8 The *error* Interface
### 7.9 Example: Expression Evaluator
### 7.10 Type Assertions
### 7.11 Discriminating Errors with Type Assertions
### 7.12 Querying Behaviors with Interface Type Assertions
### 7.13 Type Switches
### 7.14 Example: Token-Based XML Decoding
### 7.15 A Few Words of Advice

## [Chapter 8: Goroutines and Channels 协程和通道](./ch8_goroutines_and_channels/README.md)
### 8.1 Goroutines
### 8.2 Example: Concurrent Clock Server
### 8.3 Example: Concurrent Echo Server
### 8.4 Channels
### 8.5 Looping in Parallel
### 8.6 Example: Concurrent Web Crawler
### 8.7 Multiplexing with *select*
### 8.8 Example: Concurrent Directory Traversal
### 8.9 Cancellation
### 8.10 Example: Chat Server

## [Chapter 9: Concurrency with Shared Variables 基于共享变量的并发](./ch9_concurrency_with_shared_variables/README.md)
### 9.1 Race Conditions
### 9.2 Mutual Exclusion: *sync.Mutex*
### 9.3 Read/Write Mutexes: *sync.RWMutex*
### 9.4 Memory Synchronization
### 9.5 Lazy Initialization: *sync.Once*
### 9.6 The Race Detector
### 9.7 Example: Concurrent Non-Blocking Cache
### 9.8 Goroutines and Threads

## [Chapter 10: Packages and the Go Tool 包和工具](./ch10_packages_and_the_go_tool/README.md)
### 10.1 Introduction
### 10.2 Import Paths
### 10.3 The Package Declaration
### 10.4 Import Declarations
### 10.5 Blank Imports
### 10.6 Packages and Naming
### 10.7 The Go Tool

## [Chapter 11: Testing 测试](./ch11_testing/README.md)
### 11.1 The *go test* Tool
### 11.2 *Test* Functions
### 11.3 Coverage
### 11.4 *Benchmark* Functions
### 11.5 Profiling
### 11.6 *Example* Functions

## [Chapter 12: Reflection 反射](./ch12_reflection/README.md)
### 12.1 Why Reflections?
### 12.2 *reflect.Type* and *reflect.Value*
### 12.3 *Display*, a Recursive Value Printer
### 12.4 Example: Encoding S-Expressions
### 12.5 Setting Variables with *reflect.Value*
### 12.6 Example: Decoding S-Expressions
### 12.7 Accessing Struct Field Tags
### 12.8 Displaying the Methods of a Type
### 12.9 A Word of Caution

## [Chapter 13: Low-Level Programming 底层编程](./ch13_low_level_programming/README.md)
### 13.1 *unsafe.Sizeof*, *Alignof*, and *Offsetof*
### 13.2 *unsafe.Pointer*
### 13.3 Example: Deep Equivalence
### 13.4 Calling C Code with *cgo*
### 13.5 Another Word of Caution