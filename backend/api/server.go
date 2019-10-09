package api

import (
	"escapead/projectAD/backend/api/router"
	"fmt"
	"log"
	"net/http"
)

// Run start server
func Run() {
	port := ":3000"
	fmt.Println("Online and Listening on port", port)
	r := router.NEW()
	log.Fatal(http.ListenAndServe(port, r))
}
