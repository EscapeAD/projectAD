package config

import (
	"log"
	"os"
	"strconv"
)

var (
	//PORT - api port
	PORT = 0
)

// Load - Load server config
func Load() {
	var err error
	PORT, err = strconv.Atoi(os.Getenv("PORT"))
	if err != nil {
		log.Println(err)
		// Default 9000
		PORT = 9000
	}
}
