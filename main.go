package main

import (
	"log"
	"os"
)

func main() {
	filename := "log.txt"
	logargs(filename)
	log.Flags()

	os.Exit(1)
}

func logargs(filename string) {
	file, err := os.OpenFile(filename, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	log.SetOutput(file)
	log.Println(os.Args)
}
