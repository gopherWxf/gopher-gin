package goft

import (
	"log"
	"os"
)

func LoadConfigFile() []byte {
	dir, _ := os.Getwd()
	file := dir + "/application.yaml"
	b, err := os.ReadFile(file)
	if err != nil {
		log.Println(err)
		return nil
	}
	return b
}
