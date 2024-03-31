package myEmail

import (
	"fmt"
	"log"
	"net/smtp"
)

func SendMail(body []byte, senderId, appPass, port string, receiverIds []string) {
	auth := smtp.PlainAuth("", senderId, appPass, "smtp.gmail.com")
	err := smtp.SendMail("smtp.gmail.com:"+port, auth, senderId, receiverIds, body)
	if err != nil {
		fmt.Println("Error Sending Mail")
		log.Fatal(err)
	} else {
		fmt.Println("Email Was Sent Successfully. To: " + receiverIds[0])
	}
}
