// Issueshtml prints an HTML table of issues matching the search terms.
package main

import (
	"gopl.io/ch4_composite_types/4_5_json/github"
	"html/template"
	"log"
	"os"
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
//	go run ./ch4_composite_types/4_6_text_and_html_templates/issueshtml/main.go repo:golang/go commenter:gopherbot json encoder >issues.html
func main() {
	result, err := github.SearchIssues(os.Args[1:])
	if err != nil {
		log.Fatal(err)
	}

	if err := issueList.Execute(os.Stdout, result); err != nil {
		log.Fatal(err)
	}
}
