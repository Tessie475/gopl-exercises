// Exercise 4.14: a web server that queries GitHub once and then lets you
// navigate the list of issues, milestones, and users.
//
// Usage: issuesweb <owner> <repo>
//
// It fetches the repository's issues at startup, indexes them in memory, and
// serves linked HTML pages: an index, one page per issue, and filtered lists
// per user and per milestone. Reads GITHUB_PERSONAL_ACCESS_TOKEN if set (for
// private repos and to avoid rate limits).
package main

import (
	"html/template"
	"log"
	"net/http"
	"os"
	"sort"
	"strconv"
)

// store holds the data fetched once at startup, with lookup indexes.
type store struct {
	repo        string
	issues      []*Issue
	byNumber    map[int]*Issue
	byUser      map[string][]*Issue
	byMilestone map[string][]*Issue
	users       []string // sorted, for the index
	milestones  []string // sorted, for the index
}

// newStore indexes the issues by number, user, and milestone. Pull requests
// (which GitHub mixes into the issues endpoint) are skipped.
func newStore(owner, repo string, issues []*Issue) *store {
	s := &store{
		repo:        owner + "/" + repo,
		byNumber:    make(map[int]*Issue),
		byUser:      make(map[string][]*Issue),
		byMilestone: make(map[string][]*Issue),
	}
	for _, iss := range issues {
		if iss.PullRequest != nil {
			continue
		}
		s.issues = append(s.issues, iss)
		s.byNumber[iss.Number] = iss
		if iss.User != nil {
			s.byUser[iss.User.Login] = append(s.byUser[iss.User.Login], iss)
		}
		if iss.Milestone != nil {
			s.byMilestone[iss.Milestone.Title] = append(s.byMilestone[iss.Milestone.Title], iss)
		}
	}
	for u := range s.byUser {
		s.users = append(s.users, u)
	}
	for m := range s.byMilestone {
		s.milestones = append(s.milestones, m)
	}
	sort.Strings(s.users)
	sort.Strings(s.milestones)
	return s
}

// Templates. template.Must parses them at startup and panics if any is
// malformed, so a broken template is caught immediately, not per request.
// html/template auto-escapes values by context, including URL query params.
var (
	indexTmpl = template.Must(template.New("index").Parse(`
<h1>{{.Repo}}</h1>
<p>{{len .Issues}} issues</p>
<h2>Issues</h2>
<table border="1">
<tr><th>#</th><th>state</th><th>user</th><th>milestone</th><th>title</th></tr>
{{range .Issues}}
<tr>
  <td><a href="/issue?number={{.Number}}">{{.Number}}</a></td>
  <td>{{.State}}</td>
  <td>{{with .User}}<a href="/user?login={{.Login}}">{{.Login}}</a>{{end}}</td>
  <td>{{with .Milestone}}<a href="/milestone?title={{.Title}}">{{.Title}}</a>{{end}}</td>
  <td>{{.Title}}</td>
</tr>
{{end}}
</table>
<h2>Users</h2>
<ul>{{range .Users}}<li><a href="/user?login={{.}}">{{.}}</a></li>{{end}}</ul>
<h2>Milestones</h2>
<ul>{{range .Milestones}}<li><a href="/milestone?title={{.}}">{{.}}</a></li>{{end}}</ul>
`))

	issueTmpl = template.Must(template.New("issue").Parse(`
<h1>#{{.Number}} {{.Title}}</h1>
<p>state: {{.State}}</p>
<p>user: {{with .User}}<a href="/user?login={{.Login}}">{{.Login}}</a>{{end}}</p>
<p>milestone: {{with .Milestone}}<a href="/milestone?title={{.Title}}">{{.Title}}</a>{{end}}</p>
<p>created: {{.CreatedAt}}</p>
<p><a href="{{.HTMLURL}}">view on GitHub</a></p>
<pre>{{.Body}}</pre>
<p><a href="/">&larr; back</a></p>
`))

	listTmpl = template.Must(template.New("list").Parse(`
<h1>{{.Title}}</h1>
<table border="1">
<tr><th>#</th><th>state</th><th>title</th></tr>
{{range .Issues}}
<tr>
  <td><a href="/issue?number={{.Number}}">{{.Number}}</a></td>
  <td>{{.State}}</td>
  <td>{{.Title}}</td>
</tr>
{{end}}
</table>
<p><a href="/">&larr; back</a></p>
`))
)

// indexView / listView are the data shapes passed to the templates.
type indexView struct {
	Repo       string
	Issues     []*Issue
	Users      []string
	Milestones []string
}

type listView struct {
	Title  string
	Issues []*Issue
}

func (s *store) indexHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" { // "/" is a catch-all, so reject anything else
		http.NotFound(w, r)
		return
	}
	render(w, indexTmpl, indexView{
		Repo: s.repo, Issues: s.issues, Users: s.users, Milestones: s.milestones,
	})
}

func (s *store) issueHandler(w http.ResponseWriter, r *http.Request) {
	number, err := strconv.Atoi(r.URL.Query().Get("number"))
	if err != nil {
		http.Error(w, "bad issue number", http.StatusBadRequest)
		return
	}
	issue, ok := s.byNumber[number]
	if !ok {
		http.NotFound(w, r)
		return
	}
	render(w, issueTmpl, issue)
}

func (s *store) userHandler(w http.ResponseWriter, r *http.Request) {
	login := r.URL.Query().Get("login")
	render(w, listTmpl, listView{
		Title:  "Issues by " + login,
		Issues: s.byUser[login],
	})
}

func (s *store) milestoneHandler(w http.ResponseWriter, r *http.Request) {
	title := r.URL.Query().Get("title")
	render(w, listTmpl, listView{
		Title:  "Milestone: " + title,
		Issues: s.byMilestone[title],
	})
}

// render executes a template and reports a 500 if it fails.
func render(w http.ResponseWriter, t *template.Template, data any) {
	if err := t.Execute(w, data); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func main() {
	if len(os.Args) < 3 {
		log.Fatal("usage: issuesweb <owner> <repo>")
	}
	owner, repo := os.Args[1], os.Args[2]

	log.Printf("fetching issues for %s/%s ...", owner, repo)
	issues, err := fetchIssues(owner, repo)
	if err != nil {
		log.Fatal(err)
	}
	s := newStore(owner, repo, issues)
	log.Printf("loaded %d issues, %d users, %d milestones",
		len(s.issues), len(s.users), len(s.milestones))

	http.HandleFunc("/", s.indexHandler)
	http.HandleFunc("/issue", s.issueHandler)
	http.HandleFunc("/user", s.userHandler)
	http.HandleFunc("/milestone", s.milestoneHandler)

	const addr = "localhost:8000"
	log.Printf("serving on http://%s", addr)
	log.Fatal(http.ListenAndServe(addr, nil))
}
