# 第10章 包和工具

## 10.1 简介

## 10.2 导入路径
   ```go
   import (
       "fmt"
       "math/rand"
       "encoding/json"
       "golang.org/x/net/html"
       "github.com/go-sql-driver/mysql"
   )
   ```
1. 每个包是由一个全局唯一的字符串所标识的导入路径定位。
2. 如果你计划分享或发布包，那么导入路径最好是全球唯一的。

## 10.3 包声明
1. 每个Go语言源文件的开头都必须有包声明语句。
2. 通常来说，默认的包名是包导入路径名的最后一段。（例如，math/rand包和crypto/rand包的包名都是rand）
3. 关于默认包名一般采用导入路径名的最后一段的约定也有三种例外情况。
   ```text
   1. 包对应一个可执行程序，也就是main包。
   2. 包所在的目录中可能有一些文件名是以_test.go为后缀的Go源文件，并且这些源文件声明的包名也是以_test为后缀名的。
   3. 一些依赖版本号的管理工具会在导入路径后追加版本号信息。
   ```

## 10.4 导入声明
   ```go
   import "fmt"
   import "os"
   
   import (
       "fmt"
       "html/template"
       "os"
       
       "golang.org/x/net/html"
   	   "golang.org/x/net/ipv4"
   	
   	   "crypto/rand"
   	   mrand "math/rand" // alternative name mrand avoids conflict
   )
   ```
1. 可以在一个Go语言源文件包声明语句之后，其它非导入声明语句之前，包含零到多个导入包声明语句。
2. 导入的包之间可以通过添加空行来分组；通常将来自不同组织的包独自分组。
3. `导入包的重命名`：如果想同时导入两个有着相同名字的包，那么导入声明必须至少为一个同名包指定一个新的包名以避免冲突。导入包的重命名只影响当前的源文件。

## 10.5 包的匿名导入
   ```go
   import _ "image/png" // register PNG decoder
   
   import (
       "database/sql"
       _ "github.com/lib/pq"              // enable support for Postgres
       _ "github.com/go-sql-driver/mysql" // enable support for MySQL
   )
   
   db, err = sql.Open("postgres", dbname) // OK
   db, err = sql.Open("mysql", dbname)    // OK
   db, err = sql.Open("sqlite3", dbname)  // returns error: unknown driver "sqlite3"
   ```
1. 如果只是导入一个包而并不使用导入的包将会导致一个编译错误。
2. `包的匿名导入`：用下划线`_`来重命名导入的包，可以抑制`unused import`编译错误。它通常被用来实现一个编译时机制。

## 10.6 包和命名
1. 当创建一个包时，一般要用短小的包名，但也不能太短导致难以理解。标准库中最常用的包有`bufio`、`bytes`、`flag`、`fmt`、`http`、`io`、`json`、`os`、`sort`、`sync`和`time`等包。
2. 尽可能让命名有描述性且无歧义。
3. 包名一般采用单数的形式。
4. 要避免包名有其它的含义。
5. 当设计一个包的时候，需要考虑包名和成员名两个部分如何很好地配合。

## 10.7 工具
1. go命令：`go <command> [arguments]`
   ```text
   The commands are:
           bug         start a bug report
           build       compile packages and dependencies
           clean       remove object files and cached files
           doc         show documentation for package or symbol
           env         print Go environment information
           fix         update packages to use new APIs
           fmt         gofmt (reformat) package sources
           generate    generate Go files by processing source
           get         add dependencies to current module and install them
           install     compile and install packages and dependencies
           list        list packages or modules
           mod         module maintenance
           work        workspace maintenance
           run         compile and run Go program
           test        test packages
           tool        run specified go tool
           version     print Go version
           vet         report likely mistakes in packages
   
   Use "go help <command>" for more information about a command.
   ```
2. GOPATH对应的工作区目录有三个子目录。
   ```text
   GOPATH/
       src/
           gopl.io/
               .git/
               ch1/
                   helloworld/
                       main.go
                   dup/
                       main.go
                   ...
           golang.org/x/net/
               .git/
               html/
                   parse.go
                   node.go
                   ...
       bin/
           helloworld
           dup
       pkg/
           darwin_amd64/
           ...
   ```
3. `go get`命令可以下载一个单一的包或者用`...`下载整个子目录里面的每个包。
4. `go build`命令可以编译命令行参数指定的每个包。
5. `go run`命令实际上是结合了构建和运行的两个步骤。
6. Go语言中的文档注释一般是完整的句子，第一行通常是摘要说明，以被注释者的名字开头。
7. `go list`命令可以查询可用包的信息。