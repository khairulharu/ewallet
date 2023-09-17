package component

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/khairulharu/ewallet/internal/config"
	_ "github.com/lib/pq"
)

func NewDatabase(con *config.Config) *sql.DB {
	dsn := fmt.Sprintf(
		"host=%s "+
			"port=%s "+
			"user=%s "+
			"password=%s "+
			"dbname=%s "+
			"sslmode=disable",
		con.DB.Host,
		con.DB.Port,
		con.DB.User,
		con.DB.Pass,
		con.DB.Name,
	)

	database, err := sql.Open("postgres", dsn)
	if err != nil {
		log.Fatalf("error when connect to database: %s", err.Error())
	}
	log.Fatal("database connect")
	return database

}
