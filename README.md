# go-smart-sqlite-cmd

A library for SQLite using Golang

using driver:
github.com/mattn/go-sqlite3

## Installation
```
git clone https://github.com/smartgolang/gosmartsqlite.git

cd gosmartsqlite

```

## Usage
```
import (
	"database/sql"
	"github.com/smartgolang/gosmartsqlite"
)
```
### Description

Create SQLite database:

```
	db, _ := sql.Open("sqlite3", "./my.db")
	statement, _ := db.Prepare("CREATE TABLE IF NOT EXISTS people (id INTEGER PRIMARY KEY, firstname TEXT, lastname TEXT, email TEXT)")
	statement.Exec()
```

Create db struct:
```
type Person struct {
	ID        int
	FirstName string
	LastName  string
	Email     string
}
```

Upsert:
```
	var newPerson Person
	newPerson.FirstName = "Constantine"
	newPerson.LastName = "Vassil"
	newPerson.Email = "Constantine@Vassil.com"
	UpsertPerson(db, newPerson)
          
```

Search:
```
	FirstName := "Constantine"
	people, _ := SearchForPerson(db, FirstName)
	for _, ourPerson := range people {
		fmt.Printf("\n----\nID: %d\nFirst Name: %s\nLast Name: %s\nEmail: %s", ourPerson.ID, ourPerson.FirstName, ourPerson.LastName, ourPerson.Email)
	}
```

Delete:
```
	FirstName := "Constantine"
	people, _ := SearchForPerson(db, FirstName)
	for _, ourPerson := range people {
		DeletePerson(db, ourPerson.ID)
	}
```

List:
```
	people := List(db)
	for _, ourPerson := range people {
		fmt.Printf("\n----\nID: %d\nFirst Name: %s\nLast Name: %s\nEmail: %s", ourPerson.ID, ourPerson.FirstName, ourPerson.LastName, ourPerson.Email)
	}
```

