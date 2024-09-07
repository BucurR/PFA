package db

import (
	"context"
	"database/sql"
	"fmt"
	"log"

	"github.com/denisenkom/go-mssqldb/azuread"
)

type AzureSQL struct {
	DB         *sql.DB
	server     string
	port       int
	database   string
	connString string
	passwrod   string
	user       string
}

func NewAzureSQLDb(server string, port int, database string, user string, password string) *AzureSQL {
	connString := fmt.Sprintf("server=%s;user id=%s;password=%s;port=%d;database=%s;",
		server, user, password, port, database)
	return &AzureSQL{
		server:     server,
		port:       port,
		database:   database,
		connString: connString,
		passwrod:   password,
		user:       user,
	}
}
func (d *AzureSQL) ConnectDB() error {
	var err error
	d.DB, err = sql.Open(azuread.DriverName, d.connString)
	if err != nil {
		log.Fatal("Error creating connection pool: ", err.Error())
		return err
	}
	ctx := context.Background()
	err = d.DB.PingContext(ctx)
	if err != nil {
		log.Fatal(err.Error())
		return err
	}
	fmt.Printf("Connected!\n")
	return nil
}
