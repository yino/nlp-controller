package infrastructure

import (
	"log"
	"os"
)

func GetEnv() string {
	pwdPath, _ := os.Getwd()
	env := os.Getenv("env")
	defaultEnv := ""
	filename := ".yaml"
	switch env {
	case "prod":
		filename = "./prod" + filename
		defaultEnv = env
	case "test":
		filename = pwdPath+"/../" + "test" + filename
		defaultEnv = env
	default:
		filename = "./dev" + filename
		defaultEnv = "dev"
	}

	log.Println("env", defaultEnv)
	return filename
}
