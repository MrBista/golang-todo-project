package main

import (
	"fmt"
	"log"

	"github.com/MrBista/golang-todo-project/config"
	"github.com/MrBista/golang-todo-project/src/repository"
	"github.com/MrBista/golang-todo-project/src/services"
	"github.com/spf13/viper"
)

func main() {
	fmt.Println("Hello World!")
	configViper := viper.New()
	configViper.SetConfigFile("config.yaml")
	configViper.AddConfigPath(".")
	db, err := config.NewDatabase(configViper)

	// ctx := context.Background()

	if err != nil {
		log.Fatal(err)
	}
	defer db.DB.Close()

	userRepo := repository.NewUserRepository(db)
	_ = services.NewUserService(userRepo)

	// http.ListenAndServe()

}
