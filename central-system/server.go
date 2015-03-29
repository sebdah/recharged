package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/sebdah/recharged/central-system/routers"
)

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "5000"
	}

	fmt.Printf("Starting webserver on port %s\n", port)
	http.ListenAndServe(":"+port, routers.Router())
}
