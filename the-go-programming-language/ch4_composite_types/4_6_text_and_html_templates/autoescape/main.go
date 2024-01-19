package main

import (
	"html/template"
	"log"
	"os"
)

// 在终端执行：
//
//	go run ./ch4_composite_types/4_6_text_and_html_templates/autoescape/main.go > autoescape.html
func main() {
	const templ = `<p>A: {{.A}}</p><p>B: {{.B}}</p>`
	t := template.Must(template.New("escape").Parse(templ))

	var data struct {
		A string
		B template.HTML
	}

	data.A = "<b>Hello!</b>"
	data.B = "<b>Hello!</b>"

	if err := t.Execute(os.Stdout, data); err != nil {
		log.Fatal(err)
	}
}
