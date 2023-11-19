package handler

import (
	"fmt"
	"html/template"
	"net/http"
	"os"
)

func Handler(w http.ResponseWriter, r *http.Request) {
	fmt.Println(os.Getwd())
	var tmplFile = "templates/index.gohtml"
	tmpl := template.Must(template.ParseFiles(tmplFile))

	fmt.Fprint(w, tmpl)
}
