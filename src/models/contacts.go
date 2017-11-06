package models

import (
	"database/sql"
	"log"

	// mysql driver
	_ "github.com/go-sql-driver/mysql"
)

// Contact is a struct containing Contact data
type Contact struct {
	ID       int    `json:"ID"`
	LName    string `json:"LName"`
	FName    string `json:"FName"`
	Email    string `json:"Email"`
	Portable string `json:"Portable"`
	Phone    string `json:"Phone"`
}

// ContactCollection is a collection of Contacts
type ContactCollection struct {
	Contacts []Contact `json:"items"`
}

// Migrate creates the table if it haven't already got created
func Migrate(db *sql.DB) {
	sql := `
	CREATE TABLE IF NOT EXISTS contacts(
		id INT(11) NOT NULL AUTO_INCREMENT,
		lName VARCHAR(255) NULL DEFAULT NULL,
		fName VARCHAR(255) NULL DEFAULT NULL,
		email VARCHAR(255) NULL DEFAULT NULL,
		phone VARCHAR(18) NULL DEFAULT NULL,
		cPhone VARCHAR(18),

		CONSTRAINT contact_pk PRIMARY KEY (id)
	);`

	// we execute the sql query
	// we don't expect a stored result for this query so we use _
	_, err := db.Exec(sql)
	// exit if something goes wrong
	if err != nil {
		log.Panic("Migration:", err)
	}
}

// FindAll simply select all contaacts from the database, shove them into a new collection
func FindAll(db *sql.DB) ContactCollection {
	sql := `SELECT * FROM contacts`
	rows, err := db.Query(sql)
	// Exit if the SQL doesn't work for some reason
	if err != nil {
		log.Panic("[models/contacts.go/FindAll():1]", err)
	}
	// make sure to cleanup when the program exits
	defer rows.Close()

	result := ContactCollection{}
	for rows.Next() {
		contact := Contact{}
		err = rows.Scan(&contact.ID, &contact.LName, &contact.FName, &contact.Email, &contact.Portable, &contact.Phone)
		if err != nil {
			log.Panic("[models/contacts.go/FindAll():2]", err)
		}
		// Weird but it's how it is
		result.Contacts = append(result.Contacts, contact)
	}
	return result
}

// AddContact inserts a new contaact into database and returns the new id on success and panics on failure
func AddContact(db *sql.DB, lName string, fName string, email string, portable string, phone string) (int64, error) {
	sql := `INSERT INTO contacts(lName, fName, email, cPhone, phone) 
	VALUES(?, ?, ?, ?, ?)`

	// Create a prepared SQL statement
	stmt, err := db.Prepare(sql)
	// Exit if we get an error
	if err != nil {
		log.Panic("[models/contacts.go/AddContact():1]", err)
	}
	// Make sure to cleanup when the func exits
	defer stmt.Close()

	// Replace the '?' with the values
	result, err2 := stmt.Exec(lName, fName, email, portable, phone)
	// Exit if we get an error
	if err2 != nil {
		log.Panic("[models/contacts.go/AddContact():2]", err)
	}

	return result.LastInsertId()
}

// ModContact modifies a contact
func ModContact(db *sql.DB, c Contact) (int64, error) {
	// c := *pt
	sql := `UPDATE contacts
	SET lName=?, fName=?, email=?, cPhone=?, phone=?
	WHERE id=?`

	// Create a prepared SQL statement
	stmt, err := db.Prepare(sql)
	// Exit if we get an error
	if err != nil {
		log.Panic("[models/contacts.go/ModContact():1]", err)
	}

	// Replace
	result, err2 := stmt.Exec(c.LName, c.FName, c.Email, c.Portable, c.Portable, c.ID)
	if err2 != nil {
		log.Panic(c.LName, c.FName, c.Email, c.Portable, c.Portable, c.ID)
		log.Panic("[models/contacts.go/ModContact():2]", err)
	}

	return result.RowsAffected()
}

// DelContact deletes a contact
func DelContact(db *sql.DB, id int) (int64, error) {
	sql := `DELETE FROM contacts
	WHERE id = ?`

	// Create a prepared SQL statement
	stmt, err := db.Prepare(sql)
	// Exit if we get an error
	if err != nil {
		log.Panic("[models/contacts.go/DelContact():1]", err)
	}

	// Replace
	result, err2 := stmt.Exec(id)
	// Exit if we get an error
	if err2 != nil {
		log.Panic("[models/contacts.go/DelContact():2]", err)
	}

	return result.RowsAffected()
}
