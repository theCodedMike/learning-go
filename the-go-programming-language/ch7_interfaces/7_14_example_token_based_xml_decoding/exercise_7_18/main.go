// 练习7.18:
// 使用基于标记的解码API，编写一个可以读取任意XML文档并构造这个文档所代表的通用节点树的程序。
// 节点有两种类型：CharData节点表示文本字符串，和 Element节点表示被命名的元素和它们的属性。每一个元素节点有一个子节点的切片。
//
// 你可能发现下面的定义会对你有帮助。
// ```go
// import "encoding/xml"
//
// type Node interface{} // CharData or *Element
//
// type CharData string
//
//	type Element struct {
//	    Type     xml.Name
//	    Attr     []xml.Attr
//	    Children []Node
//	}
//
// ```
package main

func main() {
	// todo!
}
