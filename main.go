package main

import (
	"es/router"
	"fmt"
	"log"
	"net/http"
	"os"
)

func main() {
	logFile, _ := os.OpenFile("log.txt", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	defer logFile.Close()
	log.SetOutput(logFile)
	fmt.Println("Hello")
	r := router.Router()

	log.Println("Server started at port :8080")
	if err := http.ListenAndServe(":8080", r); err != nil {
		log.Fatal(err)
	}
}
