package cli

import (
	"encoding/json"
	"fmt"

	log "github.com/sirupsen/logrus"
)

func PrintResult(label string, v interface{}) {
	out, err := json.MarshalIndent(v, "", "   ")
	if err != nil {
		log.Panic(err)
	}
	fmt.Printf("%s: %s\n", label, string(out))
}
