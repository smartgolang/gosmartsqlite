package gosmartsqlite

// # Copyright 2022 Mobile Data Books, LLC. All rights reserved.
// # Use of this source code is governed by a BSD-style
// # license that can be found in the LICENSE file.

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

//!+Person
type Person struct {
	ID        int
	FirstName string
	LastName  string
	Email     string
}

//!-Person

//!+UpsertPerson
func UpsertPerson(db *sql.DB, ourPerson Person) {
	_, found := SearchForPerson(db, ourPerson.FirstName)
	if found {
		UpdatePerson(db, ourPerson)
	} else {
		AddPerson(db, ourPerson)
	}
}

func UpdatePerson(db *sql.DB, ourPerson Person) int64 {
	stmt, err := db.Prepare("UPDATE people set firstname = ?, lastname = ?, email = ? where id = ?")
	CheckErr(err)
	defer stmt.Close()

	res, err := stmt.Exec(ourPerson.FirstName, ourPerson.LastName, ourPerson.Email, ourPerson.ID)
	CheckErr(err)

	affected, err := res.RowsAffected()
	CheckErr(err)

	return affected
}

func AddPerson(db *sql.DB, newPerson Person) {
	stmt, _ := db.Prepare("INSERT INTO people (firstname, lastname, email) VALUES (?, ?, ?)")
	stmt.Exec(newPerson.FirstName, newPerson.LastName, newPerson.Email)
	defer stmt.Close()
}

//!-UpsertPerson

func SearchForPerson(db *sql.DB, searchString string) ([]Person, bool) {

	rows, _ := db.Query("SELECT id, firstname, lastname, email FROM people WHERE firstname like '%" + searchString + "%' OR lastname like '%" + searchString + "%' OR email like '%" + searchString + "%'")

	defer rows.Close()

	err := rows.Err()
	if err != nil {
		log.Fatal(err)
	}

	people := make([]Person, 0)
	for rows.Next() {
		ourPerson := Person{}
		err = rows.Scan(&ourPerson.ID, &ourPerson.FirstName, &ourPerson.LastName, &ourPerson.Email)
		if err != nil {
			log.Fatal(err)
		}
		people = append(people, ourPerson)
	}

	err = rows.Err()
	if err != nil {
		log.Fatal(err)
	}

	return people, len(people) > 0
}

//!+List
func List(db *sql.DB) []Person {

	rows, _ := db.Query("SELECT * FROM people")

	defer rows.Close()

	err := rows.Err()
	if err != nil {
		log.Fatal(err)
	}

	people := make([]Person, 0)

	for rows.Next() {
		ourPerson := Person{}
		err = rows.Scan(&ourPerson.ID, &ourPerson.FirstName, &ourPerson.LastName, &ourPerson.Email)
		if err != nil {
			log.Fatal(err)
		}

		people = append(people, ourPerson)
	}

	err = rows.Err()
	if err != nil {
		log.Fatal(err)
	}

	return people

}

//!-List

//!+GetPersonById
func GetPersonById(db *sql.DB, ourID string) Person {

	fmt.Printf("Value is %v", ourID)

	rows, _ := db.Query("SELECT id, firstname, lastname, email FROM people WHERE id = '" + ourID + "'")
	defer rows.Close()

	ourPerson := Person{}

	for rows.Next() {
		rows.Scan(&ourPerson.ID, &ourPerson.FirstName, &ourPerson.LastName, &ourPerson.Email)
	}

	return ourPerson
}

//!-GetPersonById

//!+DeletePerson
func DeletePerson(db *sql.DB, idToDelete int) int64 {

	stmt, err := db.Prepare("DELETE FROM people where id = ?")
	CheckErr(err)
	defer stmt.Close()

	res, err := stmt.Exec(idToDelete)
	CheckErr(err)

	affected, err := res.RowsAffected()
	CheckErr(err)

	return affected

}

//!-DeletePerson

func CheckErr(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
