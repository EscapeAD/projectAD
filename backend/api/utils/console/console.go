package console

import (
	"encoding/json"
	"fmt"
	"log"
)

// Pretty - make pretty
func Pretty(data interface{}) {
	pretty, err := json.MarshalIndent(data, "", "")
	if err != nil {
		log.Print(err)
		return
	}
	fmt.Println(string(pretty))
}
