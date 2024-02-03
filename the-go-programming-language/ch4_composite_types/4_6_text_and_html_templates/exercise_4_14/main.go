// 练习4.14：
// 创建一个web服务器，查询一次GitHub，然后生成Bug报告、里程碑和对应的用户信息。
package main

import (
	"gopl.io/ch4_composite_types/4_5_json/github"
	"html/template"
	"log"
	"net/http"
)

var issueList = template.Must(template.New("issuelist").Parse(`
<h1>{{.TotalCount}} issues</h1>
<table>
<tr style='text-align: left'>
  <th>#</th>
  <th>State</th>
  <th>User</th>
  <th>Title</th>
</tr>
{{range .Items}}
<tr>
  <td><a href='{{.HTMLURL}}'>{{.Number}}</a></td>
  <td>{{.State}}</td>
  <td><a href='{{.User.HTMLURL}}'>{{.User.Login}}</a></td>
  <td><a href='{{.HTMLURL}}'>{{.Title}}</a></td>
</tr>
{{end}}
</table>
`))

// 在终端执行：
//
//  1. 在终端输入以启动服务：go run ./ch4_composite_types/4_6_text_and_html_templates/exercise_4_14/main.go
//  2. 在浏览器地址栏输入：http://localhost:8080/issues
func main() {
	http.HandleFunc("/issues", handler)
	log.Fatal(http.ListenAndServe("localhost:8080", nil))
}

func handler(w http.ResponseWriter, _ *http.Request) {
	var terms = []string{"repo:golang/go", "commenter:gopherbot", "json", "encoder"}
	issues, err := github.SearchIssues(terms)
	if err != nil {
		w.Write([]byte(err.Error()))
	}

	if err = issueList.Execute(w, issues); err != nil {
		w.Write([]byte(err.Error()))
	}
}
