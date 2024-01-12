package main

import (
	"github.com/kevinanielsen/opinionated/src/database"
	"github.com/kevinanielsen/opinionated/src/initializers"
)

func init() {
	database.ConnectToDB()
	initializers.LoadEnvVariables()
}

func main() {

}
