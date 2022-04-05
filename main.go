package main

import (
	"mvc-go/app"
	"mvc-go/db"
)

func main() {
	db.StartDbEngine()
	app.StartRoute()
}
