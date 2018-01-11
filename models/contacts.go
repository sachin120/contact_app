package models

import (
	"errors"
	"fmt"
	"log"
)

// Contact has all table colms.
type Contact struct {
	Contid    int
	Uname     string
	Prinumber string
	Address   string
	Emailid   string
}

// AllContacts all Contacts in database
func AllContacts() ([]*Contact, error) {
	rows, err := db.Query("SELECT cont_id, u_name, pri_number, address, email_id FROM contact")
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	defer rows.Close()

	cnts := make([]*Contact, 0)
	for rows.Next() {
		ct := new(Contact)
		err := rows.Scan(&ct.Contid, &ct.Uname, &ct.Prinumber, &ct.Address, &ct.Emailid)
		if err != nil {
			return nil, err
		}
		cnts = append(cnts, ct)
	}

	if err := rows.Err(); err != nil {
		fmt.Println(err)
		return nil, err
	}
	return cnts, nil
}

// InsertContact inserts one contact
func InsertContact(arguments ...string) error {
	if len(arguments) < 4 {
		err := errors.New("require atleast 4 parameters")
		return err
	}
	statement, _ := db.Prepare("INSERT INTO contact(u_name, pri_number, address, email_id) VALUES(?, ?, ?, ?)")
	_, err := statement.Exec(arguments[0], arguments[1], arguments[2], arguments[3])
	return err
}

// DeleteContact requres contact id to delete contact
func DeleteContact(contactID int) int {
	statement, _ := db.Prepare("DELETE FROM contact WHERE cont_id=?")
	res, err := statement.Exec(contactID)
	var affectedRow int // if recorde not deleted then affectedRow is 0, else the contactID

	count, err := res.RowsAffected()
	if err != nil {
		log.Print(err)
	}
	if count > 0 {
		affectedRow = contactID
	} else {
		affectedRow = 0
	}
	return affectedRow
}
