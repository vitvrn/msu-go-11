package main

import (
	"html/template"
	"io/ioutil"
	"net/http"
)

func main() {
	tmpl := template.New("main")
	tmpl, _ = tmpl.Parse(
		`<div style="display: inline-block; border: 1px solid #aaa; border-radius: 3px;">
			<pre>{{.}}</pre>
		</div>`)

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		path := r.URL.Path
		c := http.Client{}
		resp, _ := c.Get("http://artii.herokuapp.com/make?text=" + path)
		defer resp.Body.Close()
		body, _ := ioutil.ReadAll(resp.Body)
		tmpl.Execute(w, string(body))
	})

	http.ListenAndServe(":8081", nil)
}
