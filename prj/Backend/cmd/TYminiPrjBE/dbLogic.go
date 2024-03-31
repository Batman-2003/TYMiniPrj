package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"
)

var registersDb []registerDbDetails

func loadRegistrations(dbPtr *sql.DB) {
	rows, err := dbPtr.Query(`SELECT id, username, email, passHsh, salt, ticketId 
		FROM users;`)
	if err != nil {
		log.Println("Error Selecting from Database")
		log.Fatal(err)
	}
	defer rows.Close()

	currRegister := registerDbDetails{}
	for rows.Next() {
		rows.Scan(&currRegister.id, &currRegister.username, &currRegister.email,
			&currRegister.passHsh, &currRegister.salt, &currRegister.ticketId)
		registersDb = append(registersDb, currRegister)
	}

	fmt.Println("-----------Registrations Loaded Successfully-----------")
}

func registerUser(dbPtr *sql.DB, reg registerDbDetails) {
	query := fmt.Sprintf(`INSERT INTO users(username, email, passHsh, salt)
		VALUES('%s', '%s', '%s', '%s');`,
		reg.username, reg.email, reg.passHsh, reg.salt)

	_, err = dbPtr.Exec(query)
	if err != nil {
		log.Println("Error Registering New User")
		log.Fatal(err)
	}
}

func loadTicketId(dbPtr *sql.DB, id uint32, ticketId uint64) {
	query := fmt.Sprintf(`UPDATE users SET ticketId=%v WHERE id=%v;`,
		ticketId, id)
	fmt.Println(query)
	_, err = dbPtr.Exec(query)
	if err != nil {
		log.Println("Error Updating Ticket Value")
	}
}

func updatePassword(email, passHsh string) {
	query := fmt.Sprintf(`UPDATE users SET passHsh='%s' WHERE email='%s';`,
		passHsh, email)
	_, err = dbPtr.Exec(query)
	if err != nil {
		log.Printf("Error While Updating password in DB.")
		log.Fatal(err)
	} else {
		log.Printf("Password Updated Successfully")
		registersDb = nil
		loadRegistrations(dbPtr)
	}
}

// Cleanup Old Data
func cleanup() {
	files, err := os.ReadDir("./../Frontend/resources/QRCodes/")
	if err != nil {
		log.Printf("Error While Looking For QRCodes Dir")
		log.Fatal(err)
	}

	for _, file := range files {
		deleteFile := true
		for _, regs := range registersDb {
			if file.Name() == fmt.Sprint(regs.ticketId)+".png" {
				deleteFile = false

			}
		}
		if deleteFile {
			log.Printf("File : %s is getting deleted\n", file.Name())
			path := fmt.Sprintf("./../Frontend/resources/QRCodes/%s", file.Name())
			err = os.Remove(path)
			if err != nil {
				log.Fatal("Error While Deleting a Redundant png file")
			}
		}
	}
}
