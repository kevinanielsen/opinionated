package main

import (
	"github.com/kevinanielsen/opinionated/src/database"
	"github.com/kevinanielsen/opinionated/src/initializers"
)

func init() {
	initializers.LoadEnvVariables()
	database.ConnectToDB()
}

func main() {

}
