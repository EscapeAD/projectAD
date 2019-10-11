package config

import (
	"log"
	"os"
	"strconv"

	godotenv "github.com/joho/godotenv/autoload"
)

var (
	//PORT - api port
	PORT = 0
	// DBURL - DB URL
	DBURL = ""
)

// Load - Load server config
func Load() {
	var err error
	err = godotenv.Load()
	if err != nil {
		log.Fatal(err)
	}

	PORT, err = strconv.Atoi(os.Getenv("API_PORT"))
	if err != nil {
		log.Println(err)
		// Default 9000
		PORT = 9000
	}
}
