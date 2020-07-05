package github

import "time"
import "net/url"
import "net/http"
import "fmt"
import "encoding/json"
import "strings"

//IssuesURL
const IssuesURL = "https://api.github.com/search/issues"

//IssuesSearchResult
type IssuesSearchResult struct {
	TotalCount int
	Items      []*Issue
}

//Issue
type Issue struct {
	Number    int
	HTMLURL   string
	Title     string
	State     string
	User      *User
	CreatedAt time.Time
	Body      string
}

//User
type User struct {
	Login   string
	HTMLURL string
}

//SearchIssues
func SearchIssues(terms []string) (*IssuesSearchResult, error) {
	q := url.QueryEscape(strings.Join(terms, " "))
	resp, err := http.Get(IssuesURL + "?q=" + q)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode != http.StatusOK {
		resp.Body.Close()
		return nil, fmt.Errorf("search query failed: %s", resp.Status)
	}
	var result IssuesSearchResult
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		resp.Body.Close()
		return nil, err
	}
	resp.Body.Close()
	return &result, nil
}
