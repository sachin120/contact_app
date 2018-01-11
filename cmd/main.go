package main

import (
	"contact_v4/models"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"strconv"
)

func main() {
	models.InitDB("contact.db")

	http.Handle("/jresource/", http.StripPrefix("/jresource", http.FileServer(http.Dir("/home/sachin/go_work/src/contact_v4/web/js"))))
	http.Handle("/cresource/", http.StripPrefix("/cresource", http.FileServer(http.Dir("/home/sachin/go_work/src/contact_v4/web/css"))))
	http.HandleFunc("/", allContacts)
	http.HandleFunc("/insert", insertContact)
	http.HandleFunc("/delete", delete)
	http.ListenAndServe(":3000", nil)
}

func allContacts(w http.ResponseWriter, r *http.Request) {
	conts, err := models.AllContacts()
	if err != nil {
		http.Error(w, http.StatusText(500), 500)
		return
	}
	tpl, err := template.ParseFiles("../web/html/index.html")
	if err != nil {
		log.Fatalln(err)
		http.Error(w, http.StatusText(500), 500)
		return
	}
	err = tpl.Execute(w, conts)
	if err != nil {
		log.Fatalln("Execute erorr")
		http.Error(w, http.StatusText(500), 500)
		return
	}
}

func insertContact(w http.ResponseWriter, r *http.Request) {
	name := r.FormValue("name")
	contactNo := r.FormValue("contact")
	address := r.FormValue("address")
	emailid := r.FormValue("email_id")

	err := models.InsertContact(name, contactNo, address, emailid)
	if err != nil {
		fmt.Fprintf(w, "Error inserting element\n %s", err)
	}
	fmt.Fprintf(w, "done.")
}

func delete(w http.ResponseWriter, r *http.Request) {
	value := r.URL.Query()["delete"]
	responses := make([]int, 0)
	for _, v := range value {
		i, _ := strconv.Atoi(v)
		responses = append(responses, models.DeleteContact(i))
	}
	fmt.Fprintf(w, "%v", responses)
}
