package main

import (
	"fmt"
	"log"
	"os"
)

var ConnStr = FetchEnv("DATABSE_URL")

func main() {
	fmt.Println(ConnStr)
}

func FetchEnv(key string) string {
	value, exists := os.LookupEnv(key)

	if !exists {
		log.Fatalf("FATAL: Environment variable %s is not set!", key)
	}

	return value
}
