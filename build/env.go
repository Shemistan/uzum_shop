package build

import "github.com/joho/godotenv"

const DEBUG = true

func SetConfig() error {
	err := godotenv.Load("./build/.env")
	if err != nil {
		return err
	}

	return nil
}
