package config

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

func Config(key string) string {
	fmt.Println("this is cofig")
	err := godotenv.Load(".env")
	if err != nil {
		fmt.Print("Error Loadinfg .env file")
	}
	return os.Getenv(key)
}
