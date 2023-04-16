package main

import "github.com/oaraujocesar/go-api/configs"

func main() {
	_, err := configs.LoadConfig(".")
	if err != nil {
		panic(err)
	}
}
