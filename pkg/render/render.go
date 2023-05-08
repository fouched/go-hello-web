package render

import (
	"fmt"
	"html/template"
	"net/http"
)

func RenderTemplage(w http.ResponseWriter, tmpl string) {
	parsedTemplate, _ := template.ParseFiles("./templates/"+tmpl, "./templates/base.layout.gohtml")
	err := parsedTemplate.Execute(w, nil)
	if err != nil {
		fmt.Println("Error parsing template", err)
		return
	}
}
