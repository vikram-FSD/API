package main

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

func main() {
	const (
		host     = "localhost"
		port     = 5432
		user     = "postgres"
		password = "password"
		dbname   = "first_db"
	)
	//construct connection string
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)
	//open database connection
	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		panic(err)
	}
	defer db.Close()
	//check if the connection is successful
	err = db.Ping()
	if err != nil {
		panic(err)
	}
	fmt.Println("Successfully connected to PostgreSQL Database")
	id := 1
	name := "Vikram"
	age := 20
	country := "India"
	//SQL statements with placeHolders
	sqlStatement := `INSERT INTO persons(id, name, age, country )values($1, $2, $3, $4) RETURNING id`
	//Execute the SQL statments
	var returnId sql.NullInt64 // Use sql.NullInt64 to handle NULL values for id
	err = db.QueryRow(sqlStatement, id, name, age, country).Scan(&returnId)
	if err != nil {
		panic(err)
	}

	if returnId.Valid {
		fmt.Printf("New record ID is %d\n", id)
	} else {
		fmt.Println("ID was Null")
	}

}
