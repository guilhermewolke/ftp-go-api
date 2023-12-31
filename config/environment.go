package config

import (
	"fmt"
	"os"

	project_root_directory "github.com/golang-infrastructure/go-project-root-directory"
	"github.com/joho/godotenv"
)

func getEnvValue(key string) string {
	dir, err := project_root_directory.GetRootDirectory()

	if err != nil {
		panic(err)
	}

	err = godotenv.Load(fmt.Sprintf("%s/.env", dir))

	if err != nil {
		panic(err)
	}

	return os.Getenv(key)
}
