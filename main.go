package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"time"
)

func main() {

	// check PostgreSQL connectivity
	// wait for some time to let Postgres DB to start (for docker-compose)
	fmt.Println("Going to Sleep...")
	time.Sleep(5 * time.Second)
	fmt.Println("Waking up...")
}
