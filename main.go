package main

import (
	"html/template"
	"log"
	"net/http"
)

type PageVariables struct {
	PageTitle string
	GPA       string
}

func main() {
	http.HandleFunc("/", homeRoute)
	http.HandleFunc("/selected", showGPA)
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func homeRoute(w http.ResponseWriter, r *http.Request) {
	// Display some radio buttons to the user

	Title := "GPA Calculator"
	MyPageVariables := PageVariables{
		PageTitle: Title,
	}

	t, err := template.ParseFiles("home.html") //parse the html file home.html
	if err != nil {                            // if there is an error
		log.Print("template parsing error: ", err) // log it
	}

	err = t.Execute(w, MyPageVariables) //execute the template and pass it the HomePageVars struct to fill in the gaps
	if err != nil {                     // if there is an error
		log.Print("template executing error: ", err) //log it
	}

}

func showGPA(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	// r.Form is now either
	// map[animalselect:[cats]] OR
	// map[animalselect:[dogs]]
	// so get the animal which has been selected
	x := r.Form.Get("course")

	Title := "Your preferred animal"
	MyPageVariables := PageVariables{
		PageTitle: Title,
		GPA:       x,
	}

	// generate page by passing page variables into template
	t, err := template.ParseFiles("home.html") //parse the html file homepage.html
	if err != nil {                            // if there is an error
		log.Print("template parsing error: ", err) // log it
	}

	err = t.Execute(w, MyPageVariables) //execute the template and pass it the HomePageVars struct to fill in the gaps
	if err != nil {                     // if there is an error
		log.Print("template executing error: ", err) //log it
	}
}
