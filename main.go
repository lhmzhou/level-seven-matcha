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

func crudWithDBSql() {
	var db *sql.DB
	var res sql.Result
	var err error

	// psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
	//	"password=%s dbname=%s sslmode=disable",
	//	host, port, user, password, dbname)

	psqlInfo := getPsqlInfo()

	db, err = sql.Open("postgres", psqlInfo)
	if err != nil {
		log.Fatal(err)
		panic(err)
	}
	defer db.Close()

	var userId, userType string
	var count int64

	// INSERT
	sqlStatementInsert := `
INSERT INTO accounts. users (USER_ID, USER_TYPE, ACCOUNT_NUMBER, INSRT_ID, INSRT_TS, UPDT_ID, UPDT_TS)
VALUES ($1, $2, $3, $4, CURRENT_TIMESTAMP, $5, CURRENT_TIMESTAMP)
RETURNING USER_ID;`

	err = db.QueryRow(sqlStatementInsert, "4567890123", "TYPE4", "456789012345678", "GO", "GO").Scan(&userId)
	if err != nil {
		log.Fatal(err)
		panic(err)
	}
	fmt.Println("New inserted record ID is:", userId)

	// SELECT single row
	sqlStatementSelectSingleRow := `SELECT * FROM accounts. users WHERE USER_ID=$1;`
	var user User
	row := db.QueryRow(sqlStatementSelectSingleRow, "4567890123")
	err = row.Scan(&user.UserId, &user.UserType, &user.AccountNumber,
		&user.InsrtId, &user.InsrtTs, &user.UpdtId, &user.UpdtTs)
	switch err {
	case sql.ErrNoRows:
		fmt.Println("Sorry, no rows were returned...")
		return
	case nil:
		fmt.Println(user)
	default:
		log.Fatal(err)
		panic(err)
	}

	// UPDATE with Exec
	sqlStatementUpdateExec := `
UPDATE accounts. users
SET USER_TYPE = $2
WHERE USER_ID = $1;`
	res, err = db.Exec(sqlStatementUpdateExec, "4567890123", "TYPE4.1")
	if err != nil {
		log.Fatal(err)
		panic(err)
	}
	count, err = res.RowsAffected()
	if err != nil {
		log.Fatal(err)
		panic(err)
	}
	fmt.Println("Number of rows updated: ", count)

	// UPDATE with QueryRow
	sqlStatementUpdateQueryRow := `
UPDATE accounts. users
SET USER_TYPE = $2
WHERE USER_ID = $1
RETURNING USER_ID, USER_TYPE;`

	err = db.QueryRow(sqlStatementUpdateQueryRow, "4567890123", "TYPE4.2").Scan(&userId, &userType)
	if err != nil {
		log.Fatal(err)
		panic(err)
	}
	fmt.Println("Updated USER_ID:", userId, "USER_TYPE to", userType)

	// SELECT Multiple Rows (if exists)
	rows, err := db.Query("SELECT * FROM accounts. users")
	if err != nil {
		// handle this error better than this
		log.Fatal(err)
		panic(err)
	}
	defer rows.Close()
	for rows.Next() {
		err = rows.Scan(&user.UserId, &user.UserType, &user.AccountNumber,
			&user.InsrtId, &user.InsrtTs, &user.UpdtId, &user.UpdtTs)
		if err != nil {
			// handle this error
			log.Fatal(err)
			panic(err)
		}
		fmt.Println(user)
	}
	// get any error encountered during iteration
	err = rows.Err()
	if err != nil {
		log.Fatal(err)
		panic(err)
	}

	// DELETE with Exec
	sqlStatementDeleteExec := `
DELETE FROM accounts. users
WHERE USER_ID = $1;`
	res, err = db.Exec(sqlStatementDeleteExec, "4567890123")
	if err != nil {
		log.Fatal(err)
		panic(err)
	}
	count, err = res.RowsAffected()
	if err != nil {
		log.Fatal(err)
		panic(err)
	}
	fmt.Println("Deleted", count, "row")
}

func crudWithGORM() {

	// psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
	//	"password=%s dbname=%s sslmode=disable",
	//	host, port, user, password, dbname)

	psqlInfo := getPsqlInfo()

	db, err := gorm.Open("postgres", psqlInfo)
	// db, err := gorm.Open("postgres", "jdbc:postgresql://localhost:0137/accounts?currentSchema=accounts,sslmode=disable")
	if err != nil {
		log.Fatal(err)
		panic("Error: failed to connect database via GORM")
	}
	defer db.Close()

	log.Println("Success! Connection established")

	// set the search path to accounts
	// db.Exec("SET SEARCH_PATH to accounts")

	fmt.Println("Users Table exists ?", db.Debug().HasTable(&Users{}))

	// Drops table if already exists
	//  db.Debug().DropTableIfExists(&Users{})
	// After db connection is created.
	//  db.Debug().CreateTable(&Users{})

	// Auto create table based on Model
	// db.Debug().AutoMigrate(&Users{})

	// Select all records from a model and delete all
	db.Debug().Model(&Users{}).Delete(&Users{})

	var user *Users

	user = &Users{UserId: "5678901234", UserType: "Type5", AccountNumber: "567890123456789", InsrtId: "GO", InsrtTs: time.Now(), UpdtId: "GO", UpdtTs: time.Now()}
	// Create
	db.Debug().Create(user)

	// Read
	db.Debug().First(&user, "user_id = ?", "5678901234") // find user with user_id 5678901234
	fmt.Println(user)

	// You can insert multiple records too
	var users []Users = []Users{
		Users{UserId: "1234567890", UserType: "ACTIVE", AccountNumber: "123456789012345", InsrtId: "GO", InsrtTs: time.Now(), UpdtId: "GO", UpdtTs: time.Now()},
		Users{UserId: "2345678901", UserType: "Type2", AccountNumber: "234567890123456", InsrtId: "GO", InsrtTs: time.Now(), UpdtId: "GO", UpdtTs: time.Now()},
		Users{UserId: "3456789012", UserType: "Type3", AccountNumber: "345678901234567", InsrtId: "GO", InsrtTs: time.Now(), UpdtId: "GO", UpdtTs: time.Now()},
	}

	for _, user := range users {
		db.Debug().Create(&user)
	}

	user = nil
	// Get all records
	db.Debug().Find(&users)
	for _, user := range users {
		fmt.Println(user)
	}

	// Update
	user = &Users{UserId: "5678901234", UserType: "Type5"}
	// Select, edit, and save
	db.Debug().Find(&user)
	user.UserType = "Type5.1"
	db.Debug().Save(&user)

	// Select records and delete it
	db.Debug().Table("accounts.users").Where("user_id = ?", "5678901234").Delete(&Users{})

	// Find the record and delete it
	db.Debug().Where("user_id = ?", "1234567890").Delete(&Users{})

	// Select all records from a model and delete all
	db.Debug().Model(&Users{}).Delete(&Users{})

	// You can insert multiple records too
	users = []Users{
		Users{UserId: "4567890123", UserType: "ACTIVE", AccountNumber: "123456789012345", InsrtId: "GO", InsrtTs: time.Now(), UpdtId: "GO", UpdtTs: time.Now()},
		Users{UserId: "5678901234", UserType: "Type2", AccountNumber: "234567890123456", InsrtId: "GO", InsrtTs: time.Now(), UpdtId: "GO", UpdtTs: time.Now()},
		Users{UserId: "6789012345", UserType: "Type3", AccountNumber: "345678901234567", InsrtId: "GO", InsrtTs: time.Now(), UpdtId: "GO", UpdtTs: time.Now()},
	}

	for _, user := range users {
		db.Debug().Create(&user)
	}

	user = nil
	// Get all records
	db.Debug().Find(&users)
	for _, user := range users {
		fmt.Println(user)
	}

	//// Transactions
	// txt  := db.Begin()
	// err = tx.Create(&user).Error
	// if err != nil {
	//	tx.Rollback()
	// }
	// tx.Commit()

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

	if retryCount >= 3 {
		fmt.Println("Oops...failed connecting to DB after retries: ", retryCount)
		os.Exit(1)
	}

	// make CRUD operations with database/sql
	crudWithDBSql()

	// make CRUD operations with gorm
	crudWithGORM()
}
