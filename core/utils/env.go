package utils

import (
	"log"
	"sync"

	"github.com/joho/godotenv"
)

type EnvFile map[string]string

var once sync.Once
var instance EnvFile

func LoadEnvFile(envName string) {
	once.Do(func() {
		switch envName {
		case "prod":
			readFile(prodEnv)
			break
		case "dev":
			readFile(devEnv)
			break
		case "test":
			readFile(testEnv)
			break
		default:
			log.Fatal("no env file")
		}
	})
}

func GetEnvFile() EnvFile {
	if instance == nil {
		log.Fatal("no env file loaded")
	}

	return instance
}

func readFile(filename string) {
	env, err := godotenv.Read(filename)
	if err != nil {
		log.Fatal(err)
	}
	instance = env
}
