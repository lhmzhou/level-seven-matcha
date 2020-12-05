package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	_ "github.com/lib/pq"
)

const (
	host     = "localhost"
	port     = 0137
	user     = "gouser"
	password = "postgres123"
	dbname   = "accounts"
)

type User struct {
	UserId        string
	UserType      string
	AccountNumber string
	InsrtId       string
	InsrtTs       time.Time
	UpdtId        string
	UpdtTs        time.Time
}

type Users struct {
	UserId        string `gorm:"primary_key"`
	UserType      string
	AccountNumber string
	InsrtId       string
	InsrtTs       time.Time
	UpdtId        string
	UpdtTs        time.Time
}

// TableName tells gorm which table name to use
func (c Users) TableName() string {
	return "accounts.users"
}

func getPsqlInfo() string {
	var psqlInfo string

	if os.Getenv("PG_DB_PORT") == "" {
		psqlInfo = fmt.Sprintf("host=%s port=%d user=%s "+
			"password=%s dbname=%s sslmode=disable",
			host, port, user, password, dbname)
	} else {
		psqlInfo = fmt.Sprintf("host=%s port=%s user=%s dbname=%s sslmode=disable password=%s",
			os.Getenv("PG_DB_HOST"),
			os.Getenv("PG_DB_PORT"),
			os.Getenv("PG_DB_USER"),
			os.Getenv("PG_DB_NAME"),
			os.Getenv("PG_DB_PASSWORD"))
	}

	return psqlInfo
}

func checkDBConnectivity() bool {

	//psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
	//	"password=%s dbname=%s sslmode=disable",
	//	host, port, user, password, dbname)

	psqlInfo := getPsqlInfo()

	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		fmt.Println("Oops...there's an error opening connection...")
		log.Fatal(err)
		//panic(err)
		return false
	}
	defer db.Close()

	err = db.Ping()
	if err != nil {
		fmt.Println("Oops...there's an error pinging to DB server...")
		log.Fatal(err)
		//panic(err)
		return false
	}

	fmt.Println("Score! We've successfully connected to the server!")

	return true
}

func main() {

	// check PostgreSQL connectivity
	// wait for some time to let Postgres DB to start (for docker-compose)
	fmt.Println("Going to Sleep...")
	time.Sleep(5 * time.Second)
	fmt.Println("Waking up...")

	var retryCount int8
	for retryCount = 0; retryCount < 3; retryCount++ {
		if checkDBConnectivity() {
			break
		}
		fmt.Println("Going to Sleep...")
		time.Sleep(5 * time.Second)
		fmt.Println("Waking up...")
	}

}
