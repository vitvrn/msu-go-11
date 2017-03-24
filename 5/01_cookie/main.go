package main

import (
	"fmt"
	"net/http"
)

var commentFormTmpl = `
<html>
	<body>
	<form action="/post_comment" method="post">
		Login: <input type="text" name="login"><br>
		<textarea name="comment"></textarea><br>
		<input type="submit">
	</form>
	</body>
</html>
`

func main() {

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {

		// w.Write([]byte(body))
	})

	http.HandleFunc("/get_cookie", func(w http.ResponseWriter, r *http.Request) {

	})

	http.ListenAndServe(":8081", nil)
}

//PanicOnErr panics on error
func PanicOnErr(err error) {
	if err != nil {
		panic(err)
	}
}
