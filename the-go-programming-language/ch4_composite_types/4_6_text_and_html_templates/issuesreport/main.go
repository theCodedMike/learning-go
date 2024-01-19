// Issuesreport prints a report of issues matching the search terms.
package main

import (
	"gopl.io/ch4_composite_types/4_5_json/github"
	"log"
	"os"
	"text/template"
	"time"
)

const templ = `{{.TotalCount}} issues:
{{range .Items}}----------------------------------------
Number: {{.Number}}
User:   {{.User.Login}}
Title:  {{.Title | printf "%.64s"}}
Age:    {{.CreatedAt | daysAgo}} days
{{end}}`

func daysAgo(t time.Time) int {
	return int(time.Since(t).Hours() / 24)
}

var report = template.Must(
	template.New("issuelist").Funcs(template.FuncMap{"daysAgo": daysAgo}).Parse(templ),
)

// 在终端执行：
//
//	go run ./ch4_composite_types/4_6_text_and_html_templates/issuesreport/main.go repo:golang/go is:open json decoder
func main() {
	result, err := github.SearchIssues(os.Args[1:])
	if err != nil {
		log.Fatal(err)
	}

	if err := report.Execute(os.Stdout, result); err != nil {
		log.Fatal(err)
	}
}

func noMust() {
	report, err := template.New("report").Funcs(template.FuncMap{"daysAgo": daysAgo}).Parse(templ)
	if err != nil {
		log.Fatal(err)
	}

	result, err := github.SearchIssues(os.Args[1:])
	if err != nil {
		log.Fatal(err)
	}

	if err := report.Execute(os.Stdout, result); err != nil {
		log.Fatal(err)
	}
}

// $ go run ./ch4/4_6_text_and_html_templates/issuesreport/main.go repo:golang/go is:open json decoder
// 86 issues:
// ----------------------------------------
// Number: 48298
// User:   dsnet
// Title:  encoding/json: add Decoder.DisallowDuplicateFields
// Age:    859 days
// ----------------------------------------
// Number: 61627
// User:   nabice
// Title:  x/tools/gopls: feature: CLI syntax for renaming by identifier, n
// Age:    172 days
// ----------------------------------------
// Number: 11046
// User:   kurin
// Title:  encoding/json: Decoder internally buffers full input
// Age:    3149 days
// ...
