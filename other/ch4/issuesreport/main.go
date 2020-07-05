package main

import (
	"GoStudy/ch4/github"
	"html/template"
	"log"
	"os"
	"time"
)

const templ = `{{.TotalCount}}	issues: 
{{range	.Items}}---------------------------------------
Number:	{{.Number}} 
User:	{{.User.Login}} 
Title:	{{.Title | printf "%.64s"}} 
Age:	{{.CreatedAt | daysAgo}} days 
{{end}}`

func daysAgo(t time.Time) int {
	return int(time.Since(t).Hours() / 24)
}

var report = template.Must(template.New("issuelist").
	Funcs(template.FuncMap{"daysAgo": daysAgo}).
	Parse(templ))

func main() {

	key := []string{"a"}
	result, err := github.SearchIssues(key)
	if err != nil {
		log.Fatal(err)
	}
	if err := report.Execute(os.Stdout, result); err != nil {
		log.Fatal(err)
	}

}
