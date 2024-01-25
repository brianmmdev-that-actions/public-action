package main

import (
	"log"
	"os"
)

func main() {
	args := os.Args
	if len(args) > 1 {
		log.Printf("Hello %v!!", args[1])
		return
	}
	log.Println("Hello world!!")
}
