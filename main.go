package main

import (
	"html/template"
	"log"
	"os"
	//"io/ioutil"
)

type User struct {
	FirstName string
	LastName string
	Email string
	Age int
}

func (u User) IsOld() bool  {
	return u.Age > 30
}

func main()  {
	u := User{"Mark", "Zuckerberg", "mark@example.com", 37}
	//body, _ := ioutil.ReadFile("templates/file.tmpl")
	//tmpl, err := template.New("Some Name").Parse(string(body))

	tmpl, _ := template.ParseFiles("templates/file2.html", "templates/file3.html")
	err := tmpl.ExecuteTemplate(os.Stdout, "file2.html", u)

	if err != nil {
		log.Panic(err)
	}
	tmpl.Execute(os.Stdout, u)
}