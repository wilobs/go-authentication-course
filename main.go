package main

import (
	"fmt"
	"net/http"
	"net/url"
)

func main() {
	http.HandleFunc("/", index)
	http.HandleFunc("/register", register)
	http.ListenAndServe(":8081", nil)
}

func index(w http.ResponseWriter, r *http.Request) {
	errMsg := r.FormValue("errormsg")
	if errMsg != "" {
	}
	html := `
	<!DOCTYPE html>
	<html lang="en">
	<head>
		<meta charset="UTF-8">
		<meta http-equiv="X-UA-Compatible" content="IE=edge">
		<meta name="viewport" content="width=device-width, initial-scale=1.0">
		<title>Document</title>
	</head>
	<body>
		<h1>%s At all</h1>
		<form action="/register" method="post">
			<input type="email" name="e">
			<input type="password" name="p">
			<input type="submit">
		</form>
	</body>
	</html>`

	fmt.Fprintf(w, html, errMsg)
}

func register(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		errorMsg := url.QueryEscape("your method was not POST")
		http.Redirect(w, r, "/?errormsg="+errorMsg, http.StatusSeeOther)
		return
	}
}
