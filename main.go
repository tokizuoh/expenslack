package main

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		fmt.Println("error")
	}
	accessToken := os.Getenv("ACCESS_TOKEN")
	fmt.Println(accessToken)
}
