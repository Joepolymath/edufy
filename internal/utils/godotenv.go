package utils

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

func LoadEnv(env_path string) error {
	wd, err := GetProjectRootPath()
	if err != nil {
		log.Fatal(err)
	}
	return godotenv.Load(wd + env_path)
	// if err != nil {
	// 	log.Println(err)
	// 	log.Fatalf("Error loading specified environment file")
	// }
}

func GoDotEnvVariable(key string) string {
	return os.Getenv(key)
}
