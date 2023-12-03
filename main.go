package main

import (
	"html/template"
	"net/http"
	"os"
)

type Context struct {
	Color   string
	Message string
}

func handle(w http.ResponseWriter, r *http.Request) {
	color, ok := os.LookupEnv("COLOR")
	if !ok {
		color = "black"
	}

	message, ok := os.LookupEnv("MESSAGE")
	if !ok {
		message = "Hello world"
	}

	context := Context{
		Color:   color,
		Message: message,
	}

	page := `
		<html>
			<head>
				<title>Hello world</title>
			</head>
			<style>
				h1 {
					color: {{.Color}};
					text-align: center;
				}
			</style>
			<body>
				<h1>{{.Message}}</h2>
			</body>
		</html>
	`
	t := template.New("hello-world")
	t, _ = t.Parse(page)
	t.Execute(w, context)
}

func main() {
	http.HandleFunc("/", handle)
	http.ListenAndServe(":8000", nil)
}
