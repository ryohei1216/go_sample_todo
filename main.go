package main

import (
	"fmt"

	"main/app/controllers"
	"main/app/models"
)

func TestConnection() {

}

func main() {
	fmt.Println(models.Db)
	go controllers.StartMainServer()

	for {

	}
}
