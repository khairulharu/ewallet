package component

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/khairulharu/ewallet/internal/config"
	_ "github.com/lib/pq"
)

func GetDatabaseConnection(cnf *config.Config) *sql.DB {
	dsn := fmt.Sprintf(""+
		"host=%s "+
		"port=%s "+
		"user=%s "+
		"password=%s "+
		"dbname=%s "+
		"sslmode=disable",
		cnf.Database.Host,
		cnf.Database.Port,
		cnf.Database.User,
		cnf.Database.Pass,
		cnf.Database.Name,
	)

	database, err := sql.Open("postgres", dsn)
	if err != nil {
		log.Fatalf("error when connect to database: %s", err.Error())
	}
	err = database.Ping()
	if err != nil {
		log.Fatalf("error when connect to database: %s", err.Error())
	}
	return database

}
