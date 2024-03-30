package main

import (
	"database/sql"
	"fmt"
	"log"
)

var registersDb []registerDbDetails

func loadRegistrations(dbPtr *sql.DB) {
	rows, err := dbPtr.Query(`SELECT id, username, email, passHsh, salt, ticketId 
		FROM users;`)
	if err != nil {
		fmt.Println("Error Selecting from Database")
		log.Fatal(err)
	}
	defer rows.Close()

	currRegister := registerDbDetails{}
	it := 0
	for rows.Next() {
		rows.Scan(&currRegister.id, &currRegister.username, &currRegister.email,
			&currRegister.passHsh, &currRegister.salt, &currRegister.ticketId)
		registersDb = append(registersDb, currRegister)
		it++
	}

	fmt.Println("-----------Registrations Loaded Successfully-----------")
}

func registerUser(dbPtr *sql.DB, reg registerDbDetails) {
	query := fmt.Sprintf(`INSERT INTO users(username, email, passHsh, salt)
		VALUES('%s', '%s', '%s', '%s');`,
		reg.username, reg.email, reg.passHsh, reg.salt)

	_, err = dbPtr.Exec(query)
	if err != nil {
		fmt.Println("Error Registering New User")
		log.Fatal(err)
	}
}

func loadTicketId(dbPtr *sql.DB, id uint32, ticketId uint64) {
	query := fmt.Sprintf(`UPDATE users SET ticketId=%v WHERE id=%v;`,
		ticketId, id)
	_, err = dbPtr.Exec(query)
	if err != nil {
		fmt.Println("Error Updating Ticket Value")
	}
}
