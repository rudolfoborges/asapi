package main

import (
	"github.com/joho/godotenv"
	"github.com/rudolfoborges/asapi/internal/app"
)

func init() {
	godotenv.Load()
}

func main() {
	app.Run()
}
