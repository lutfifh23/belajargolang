package main

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

type Person struct {
	Id      int
	Name    string
	Address string
}

// type MyDriver struct{}

// func (md *MyDriver) Open(name string) (Conn, error)

// func (md *MyDriver) Open(name string) (Conn, error)

// func (md *MyDriver) Open(name string) (Conn, error)

func main() {
	// sql.Register("ini_driver_gw", &MyDriver{})
	fmt.Println("hello world")

	connectionString := "host=localhost port=5432 user=postgres password=#Cancer23 dbname=sesi7 sslmode=disable"
	db, err := sql.Open("postgres", connectionString)

	if err != nil {
		panic(err)
	}
	// getAll(db)
	insert(db, "Oniel", "Tangerang")
}

func insert(db *sql.DB, name, address string) {
	query := "insert into person(name, address) values($1, $2) returning *"

	row := db.QueryRow(query, name, address)

	var p Person

	err := row.Scan(&p.Id, &p.Name, &p.Address)

	if err != nil {
		panic(err)
	}

	fmt.Println(p)
}

func getAll(db *sql.DB) {
	query := "select * from person"

	rows, err := db.Query(query)

	if err != nil {
		panic(err)
	}

	var persons = []Person{}

	for rows.Next() {
		var p = Person{}

		err = rows.Scan(&p.Id, &p.Name, &p.Address)

		if err != nil {
			continue
		}

		persons = append(persons, p)
	}

	fmt.Println(persons)
}
