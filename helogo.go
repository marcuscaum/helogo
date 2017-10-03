package main

import (
	"fmt"
	"io/ioutil"
	"encoding/json"
	"net/http"
	"html/template"
)

var templates *template.Template
var homeTemplate *template.Template


type Page struct {
	Head struct {
		Title string
		Content string
	}
	Body struct {
		Title string
		Content string
	}
	Footer struct {
		Title string
		Content string
	}
}

func main() {
	PopulateTemplate()

	http.HandleFunc("/", HomeFunc)
	fmt.Println("Server running on 8081")
	http.ListenAndServe("0.0.0.0:8081", nil)
}

func PopulateTemplate() {
	templates, err := template.ParseGlob("./templates/*.html")
	check(err)

	homeTemplate = templates.Lookup("home.html")
}

func HomeFunc(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		p := TransformJson()
		homeTemplate.Execute(w, p)
	}
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func JsonReader() []byte {
	j, err := ioutil.ReadFile("page.json")
	check(err)

	return j
}

func TransformJson() Page {
	var p Page
	j := JsonReader()

	err := json.Unmarshal(j, &p)
	check(err)
	
	return p
}


