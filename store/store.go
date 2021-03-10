package store

import (
	"fmt"

	"github.com/3dw1nM0535/galva/utils"
	"github.com/jinzhu/gorm"
	"github.com/joho/godotenv"

	// postgres dialect
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

var dbHost, dbPort, dbName, dbUser, dbPass, sslMode string
var logMode bool

func init() {
	godotenv.Load()
	dbHost = utils.MustGetEnv("DBHOST")
	dbPort = utils.MustGetEnv("DBPORT")
	dbName = utils.MustGetEnv("DBNAME")
	dbUser = utils.MustGetEnv("DBUSER")
	dbPass = utils.MustGetEnv("DBPASS")
	sslMode = utils.MustGetEnv("SSL_ENABLED")
	logMode = utils.MustGetBool("LOGGING_STATUS")
}

type ORM struct {
	Store *gorm.DB
}

func Factory() (*ORM, error) {
	dbURI := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=%s", dbHost, dbPort, dbUser, dbPass, dbName, sslMode)
	// pass URI and dialect to gorm.Open
	dbm, err := gorm.Open("postgres", dbURI)
	if err != nil {
		fmt.Printf("Error connecting to data store: " + err.Error())
	}
	// should we log every sql query to stdout?
	dbm.LogMode(logMode)
	return &ORM{Store: dbm}, nil
}
