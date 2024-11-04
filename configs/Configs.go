package configs

import (
	"log"

	"github.com/joho/godotenv"
)

func LoadEnv(){
	if err := godotenv.Overload(".Env"); err != nil{
		log.Fatal("unable to load environment variables")
	}
}