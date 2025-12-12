package hades

import (
	"fmt"
	"os"
)

func GetDomain(debug bool) string {
	if debug {
		port := os.Getenv("PORT")
		return fmt.Sprintf("http://localhost:%s", port)
	} else {
		domain := os.Getenv("DOMAIN")
		return fmt.Sprintf("https://%s", domain)
	}
}
