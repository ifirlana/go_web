package main


import (
	"html/template"
	"net/http"
	"log"
)

type Gopher struct {
	Name string
}

func hello(w http.ResponseWriter, r *http.Request)  {
	var gophername string
	gophername = r.URL.Query().Get("name");

	if (gophername == "") {
		gophername = "GOPHER"
	}
	gopher := Gopher{Name: gophername}
	renderTemplate(w, "./template/index.html", gopher)
}

func renderTemplate(w http.ResponseWriter, templateFile string, templateData interface{})  {
	t, err := template.ParseFiles(templateFile)
	if err != nil {
		log.Fatal("Error Encountered while parse the template: ", err)
	}
	t.Execute(w, templateData)
}

func main() {
	http.HandleFunc("/", hello);
	http.ListenAndServe(":8080", nil);
}
