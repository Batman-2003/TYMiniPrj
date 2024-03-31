package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/Batman-2003/TYminiPrj/Prj/Backend/internal/myQRLib"
	"golang.org/x/crypto/bcrypt"
)

//-------------------------------Func Defs-------------------------------

func loginHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		tmpl.ExecuteTemplate(w, "login.html", nil)
		return
	}

	loglet := loginDetails{
		username: r.FormValue("user"),
		password: r.FormValue("pass"),
	}

	for it, log := range registersDb {
		if loglet.username == log.username {
			currSalt := registersDb[it].salt
			err = bcrypt.CompareHashAndPassword([]byte(registersDb[it].passHsh),
				[]byte(loglet.password+currSalt))
			if err != nil {
				fmt.Println("Wrong Password")
				tmpl.ExecuteTemplate(w, "login.html", "Passwords don't match")
				return
			} else {
				user.Username = log.username
				user.Id = log.id
				user.TicketId = log.ticketId
				user.UserQR = fmt.Sprintf(`<img src="./../resources/QRCodes/%v.png" width="128px">`,
					log.ticketId)
				http.Redirect(w, r, "/index", http.StatusSeeOther)
			}
		}
	}
}

func registerHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		tmpl.ExecuteTemplate(w, "register.html", nil)
		return
	}

	emptyRegistret := registerDetails{}

	registret := registerDetails{
		email:    r.FormValue("email"),
		username: r.FormValue("user"),
		password: r.FormValue("pass"),
	}

	if registret == emptyRegistret {
		tmpl.ExecuteTemplate(w, "register.html", "Empty Req")
		return
	}

	newRegistret := true
	msgString := ""

	for _, reg := range registersDb {
		if reg.username == registret.username {
			newRegistret = false
			msgString = "Username Already Registered, Try Logging in"
		} else if reg.email == registret.email {
			newRegistret = false
			msgString = "Email Already Registered, Try Logging in"
		}
	}

	if newRegistret {
		registeretDb := registerDbDetails{
			username: registret.username,
			email:    registret.email,
			salt:     "",
		}

		passHshBytes, err := bcrypt.GenerateFromPassword([]byte(registret.password+registeretDb.salt), bcrypt.MinCost)
		if err != nil {
			fmt.Println("Error During Encryption")
			log.Fatal(err)
		}
		registeretDb.passHsh = string(passHshBytes)

		registerUser(dbPtr, registeretDb)
		registersDb = nil
		loadRegistrations(dbPtr)

		http.Redirect(w, r, "/index", http.StatusSeeOther)
	} else {
		tmpl.ExecuteTemplate(w, "register.html", msgString)
	}

}

func bookingFormHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		tmpl.ExecuteTemplate(w, "bookingForm.html", nil)
		return
	}

	currBooking := bookingFormIp{}
	premium, _ := strconv.Atoi(r.FormValue("premiumTicks"))
	base, _ := strconv.Atoi(r.FormValue("baseTicks"))
	minimum, _ := strconv.Atoi(r.FormValue("minimumTicks"))

	currBooking.tier1 = uint32(premium)
	currBooking.tier2 = uint32(base)
	currBooking.tier3 = uint32(minimum)

	currFeedback := bookingTicketFeedback{
		AddedToCart: true,
		MsgString: fmt.Sprintf("The Total is : %v", (currBooking.tier1*t1Cost)+
			(currBooking.tier2*t2Cost)+(currBooking.tier3*t3Cost)),
		Premium:      currBooking.tier1,
		Base:         currBooking.tier2,
		Minimum:      currBooking.tier3,
		PremiumCost:  t1Cost,
		BaseCost:     t2Cost,
		MinimumCost:  t3Cost,
		PremiumTotal: currBooking.tier1 * t1Cost,
		BaseTotal:    currBooking.tier2 * t2Cost,
		MinimumTotal: currBooking.tier3 * t3Cost,
	}

	if (currFeedback.Premium + currFeedback.Base + currFeedback.Minimum) > 0 {
		user.TicketId = uint64((100 * 100 * 100 * (user.Id + 10)) +
			(currFeedback.Premium * 100 * 100) + (currFeedback.Base * 100) +
			(currFeedback.Minimum))
		user.UserQR = fmt.Sprintf(`<img src="./../resources/QRCodes/%v.png" width="128px">`,
			user.TicketId)
		loadTicketId(dbPtr, user.Id, user.TicketId)

		data := "Username: " + user.Username + "\r\n" +
			"TicketId: " + fmt.Sprint(user.TicketId)
		myQRLib.CreateQRCode(data, uint32(user.TicketId))
	}

	tmpl.ExecuteTemplate(w, "bookingForm.html", currFeedback)
}

func receiptHandler(w http.ResponseWriter, r *http.Request) {
	tmpl.ExecuteTemplate(w, "receipt.html", user)
}

func contactUsHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		tmpl.ExecuteTemplate(w, "contactUs.html", nil)
	} else {
		tmpl.ExecuteTemplate(w, "contactUs.html", nil)
	}
}

func homeHandler(w http.ResponseWriter, r *http.Request) {
	http.Redirect(w, r, "/login", http.StatusSeeOther)
}

func logoutHandler(w http.ResponseWriter, r *http.Request) {
	user.Username = ""
	user.TicketId = 0
	user.Id = 0
	user.UserQR = ""
	http.Redirect(w, r, "/index", http.StatusSeeOther)
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	tmpl.ExecuteTemplate(w, "index.html", user)
}
