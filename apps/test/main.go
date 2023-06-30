package main

import (
	"log"

	_ "github.com/gestgo/gest/package/technique/version"
)

func main() {
	log.Println("hello")
	// go build -o main --ldflags="-X 'github.com/gestgo/gest/package/technique/version.Version=0.0.2' -X 'github.com/gestgo/gest/package/technique/version.Date=$(date)'"
}
