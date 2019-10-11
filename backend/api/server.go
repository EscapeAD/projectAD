package api

import (
	"escapead/projectAD/backend/api/config"
	"escapead/projectAD/backend/api/router"
	"escapead/projectAD/backend/auto"
	"fmt"
	"log"
	"net/http"
)

// Run start server
func Run() {
	config.Load()
	auto.Load()
	fmt.Println("Online and Listening on port", config.PORT)
	Listen(config.PORT)
}

// Listen - Load Server
func Listen(port int) {
	r := router.NEW()
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", port), r))
}
