package main

import (
	"github.com/kevinanielsen/opinionated/src/database"
	"github.com/kevinanielsen/opinionated/src/initializers"
	"github.com/kevinanielsen/opinionated/src/router"
)

func init() {
	initializers.LoadEnvVariables()
	database.ConnectToDB()
}

func main() {
	router.Router()
}
