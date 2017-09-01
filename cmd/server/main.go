package main

import (
	"log"
	"os"

	"github.com/aashishthite/message_parser/pkg/handlers"
)

func init() {
	log.SetOutput(os.Stdout)
	log.SetPrefix("")
	log.SetFlags(log.Lshortfile | log.Ldate | log.Ltime)
}

func main() {
	log.Println("Hello World")

	h := handlers.New()
	log.Fatalf(h.ListenAndServe().Error())
}
