package model

import (
	"log"
	"time"

	"github.com/atomedgesoft/scheduler/config"
)

type User struct {
	Id            string    `json:id`
	FirstName     string    `json:firstname`
	LastName      string    `json:lastname`
	EmailAddress  string    `json:emailaddress`
	Signinthrough string    `json:signinthrough`
	CreatedAt     time.Time `json:createdate`
	TimeZone      string    `json:timezone`
	IsActive      bool      `json:isactive`
	Country       string    `json:country`
}

func InsertUser(user User) string {
	db, _ := config.ConnectDB()
	defer db.Close()

	sqlStmt := `insert into users(id, firstname, lastname, emailaddress, signinthrough, createdat, timezone, country)values($1, $2, $3, $4, $5, $6, $7, $8) RETURNING id`

	err := db.QueryRow(sqlStmt, user.Id, user.FirstName, user.LastName, user.EmailAddress, user.Signinthrough, user.CreatedAt, user.TimeZone, user.Country).Scan(&user.Id)
	if err != nil {
		log.Fatal(err)
	}
	return user.Id
}
