package basicKnowledge

import (
	"html/template"
	"os"
	"time"
)

type User struct {
	UserName,Password string
	RegTime time.Time
}

func ShowTime(t time.Time, format string)string  {
	return t.Format(format)
}

func TestTemplate(){
	u :=User{"yourname","passd",time.Now()}
	t,err:= template.New("text").Funcs(template.FuncMap{"showtime":ShowTime}).Parse(`<p>{{.UserName}}|{{.Password}}
| {{.RegTime.Format "2006-01-02 15:04:05"}}</p><p>{{.UserName}} | {{.Password}} |{{showtime .RegTime "2006-01-02 15:04:05"}}</p>`)
	if err != nil{
		panic(err)
	}
	t.Execute(os.Stdout, u)
}
