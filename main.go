package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/MrBista/golang-todo-project/config"
	"github.com/MrBista/golang-todo-project/src/controllers"
	"github.com/MrBista/golang-todo-project/src/repository"
	"github.com/MrBista/golang-todo-project/src/services"
	"github.com/julienschmidt/httprouter"
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
	userService := services.NewUserService(userRepo)
	userController := controllers.NewUserController(userService)

	router := httprouter.New()

	router.POST("/api/v1/auth/register", userController.GetUserByEmail)
	router.GET("/users/:email", userController.GetUserByEmail)

	server := http.Server{
		Addr:    "127.0.0.1:8000",
		Handler: router,
	}

	err = server.ListenAndServe()

	log.Fatal(err)

}
