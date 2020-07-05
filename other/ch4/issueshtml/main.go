package main

import (
	"GoStudy/ch4/github"
	"html/template"
	"log"
	"os"
)

var report = template.Must(template.New("issuelist").Parse(`
	<h1>{{.TotalCount}}	issues</h1> 
	<table> 
		<tr	style='text-align:	left'>		
			<th>#</th>		
			<th>State</th>		
			<th>User</th>		
			<th>Title</th> 
		</tr> 
		{{range	.Items}} 
		<tr>		
			<td><a	href='{{.HTMLURL}}'>{{.Number}}</a></td>		
			<td>{{.State}}</td>		
			<td><a	href='{{.User.HTMLURL}}'>{{.User.Login}}</a></td>	
			<td><a	href='{{.HTMLURL}}'>{{.Title}}</a></td> 
		</tr> 
		{{end}} 
		</table>
		 `))

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
