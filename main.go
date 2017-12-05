package main

import (
	//"html/template"
	//"log"
	//"os"
	//"io/ioutil"
	"net/http"
	//"fmt"
	"time"
	"encoding/json"
)

//type User struct {
//	FirstName string
//	LastName string
//	Email string
//	Age int
//}

//func (u User) IsOld() bool  {
//	return u.Age > 30
//}
type User struct {
	FirstName string `json:"first_name"`
	LastName string `json:"last_name"`
	Email string `json:"email"`
	CreatedAt time.Time
}

type HomePageHandler struct{}

func (h *HomePageHandler) ServeHTTP(w http.ResponseWriter, r *http.Request)  {
	//name := r.URL.Query().Get("name")
	//if name == "" {
	//	name = "World";
	//}
	//
	//fmt.Fprintf(w, "Hello %s!", name)
	user := new(User)
	json.NewDecoder(r.Body).Decode(user)
	user.CreatedAt = time.Now()

	w.Header().Set("Content-Type","application/json")
	w.WriteHeader(http.StatusCreated)

	data, _:= json.Marshal(user)
	w.Write(data)
}

func main()  {
	//u := User{"Mark", "Zuckerberg", "mark@example.com", 37}
	////body, _ := ioutil.ReadFile("templates/file.tmpl")
	////tmpl, err := template.New("Some Name").Parse(string(body))
	//
	//tmpl, _ := template.ParseFiles("templates/file2.html", "templates/file3.html")
	//err := tmpl.ExecuteTemplate(os.Stdout, "file2.html", u)
	//
	//if err != nil {
	//	log.Panic(err)
	//}
	//tmpl.Execute(os.Stdout, u)

	//http.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
	//	fmt.Fprintln(writer, "Hello World!")
	//})
	//
	//barHandle := func(w http.ResponseWriter, r *http.Request) {
	//	fmt.Fprintln(w, "Hello Bar!")
	//}
	//
	//http.HandleFunc("/bar", barHandle)

	http.Handle("/", &HomePageHandler{})

	http.ListenAndServe(":3000",nil)

}