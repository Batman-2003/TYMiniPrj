package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"
	"text/template"

	"github.com/joho/godotenv"
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
	cleanup()

	err = godotenv.Load()
	if err != nil {
		fmt.Println("Error Loading Environment Variables")
		log.Fatal(err)
	}
	email = os.Getenv("email")
	apass = os.Getenv("apass")
	port = os.Getenv("port")

	tmpl, err = template.ParseGlob("./../Frontend/html/*.html")
	if err != nil {
		log.Fatal(err)
	}

	http.HandleFunc("/", homeHandler)
	http.HandleFunc("/login", loginHandler)
	http.HandleFunc("/login/forgotPass", forgotPassHandler)
	http.HandleFunc("/login/forgotPass/changePass", changePassHandler)
	http.HandleFunc("/login/forgotPass/updatePass", updatePassHandler)

	http.HandleFunc("/register", registerHandler)
	http.HandleFunc("/index", indexHandler)
	http.HandleFunc("/ticket-booking", bookingFormHandler)
	http.HandleFunc("/receipt", receiptHandler)
	http.HandleFunc("/contactUs", contactUsHandler)
	http.HandleFunc("/logout", logoutHandler)

	http.Handle("/Resources/", http.StripPrefix("/Resources", http.FileServer(http.Dir("./../Frontend/Resources/"))))

	log.Fatal(http.ListenAndServe(":42069", nil))
}
