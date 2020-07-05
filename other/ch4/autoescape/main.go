package main

import "html/template"
import "os"
import "log"

func main() {
	const templ = `<p>A: {{.A}}</p><p>B:	{{.B}}</p>`
	t := template.Must(template.New("excape").Parse(templ))
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
