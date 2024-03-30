package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"text/template"

	_ "github.com/mattn/go-sqlite3"
)

// -------------------------------Variables------------------------------
var err error
var tmpl *template.Template
var dbPtr *sql.DB

// -------------------------------Main Func------------------------------
func main() {
	fmt.Println("----------The Web Server Started Successfully----------")

	dbPtr, err = sql.Open("sqlite3", "./tickzy.db")
	if err != nil {
		fmt.Println("Error Connecting to Database")
		log.Fatal(err)
	}

	loadRegistrations(dbPtr)

	tmpl, err = template.ParseGlob("./../Frontend/html/*.html")
	if err != nil {
		log.Fatal(err)
	}

	http.HandleFunc("/", homeHandler)
	http.HandleFunc("/login", loginHandler)
	http.HandleFunc("/register", registerHandler)
	http.HandleFunc("/index", indexHandler)
	http.HandleFunc("/ticket-booking", bookingFormHandler)
	http.HandleFunc("/receipt", receiptHandler)
	http.HandleFunc("/contactUs", contactUsHandler)
	http.HandleFunc("/logout", logoutHandler)

	http.Handle("/resources/", http.StripPrefix("/resources", http.FileServer(http.Dir("./../Frontend/resources/"))))

	log.Fatal(http.ListenAndServe(":8420", nil))
}
