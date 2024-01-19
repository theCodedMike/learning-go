# 第1章 入门

1. Go语言只有for循环这一种循环语句
   ```text
   for initialization; condition; post {
       // zero or more statements
   }
   
   for condition {
       // a traditional "while" loop
   }
   
   for {
       // a traditional infinite loop
   }
   ```
2. 在Go语言中，只有i++/i--是合法的，++i/--i是非法的
3. 符号:=是短变量声明（short variable declaration）的一部分
4. 声明一个变量有好几种方式，下面这些都等价
   ```text
   s := ""（只能用在函数内部，而不能用于包变量）
   var s string （依赖于字符串的默认初始化零值机制，被初始化为""）
   var s = ""
   var s string = ""
   ```
5. fmt.Printf函数的格式化输出如下:
   ```text
   %d          十进制整数
   %x, %o, %b  十六进制，八进制，二进制整数。
   %f, %g, %e  浮点数： 3.141593 3.141592653589793 3.141593e+00
   %t          布尔：true或false
   %c          字符（rune） (Unicode码点)
   %s          字符串
   %q          带双引号的字符串"abc"或带单引号的字符'c'
   %v          变量的自然形式（natural format）
   %T          变量的类型
   %%          字面上的百分号标志（无操作数）
   ```
