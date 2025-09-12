package main

import (
	"net/http"

	"github.com/MrBista/golang-todo-project/config"
	"github.com/MrBista/golang-todo-project/helper"
	"github.com/MrBista/golang-todo-project/src/app"
	"github.com/MrBista/golang-todo-project/src/controllers"
	"github.com/MrBista/golang-todo-project/src/repository"
	"github.com/MrBista/golang-todo-project/src/services"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

func main() {
	configViper := viper.New()
	configViper.SetConfigFile("config.yaml")
	configViper.AddConfigPath(".")
	logrus.SetLevel(logrus.TraceLevel)

	if err := configViper.ReadInConfig(); err != nil {
		logrus.Fatal(err)

	}

	db, errDb := config.NewDatabase(configViper)
	logrus.Info("Masuk sini 0")
	if errDb != nil {
		logrus.Fatal(errDb)
	}

	userRepo := repository.NewUserRepository()
	userService := services.NewUserService(userRepo, db)
	userController := controllers.NewUserController(userService)

	todoRepo := repository.NewTodo()
	todoServices := services.NewTodo(todoRepo, db)
	todoController := controllers.NewTodoController(todoServices)

	router := app.NewRouter(userController, todoController)

	port := configViper.GetString("app.port")
	server := http.Server{
		Addr:    "127.0.0.1:" + port,
		Handler: router,
	}

	helper.Logger().Info("App runing in port " + port)

	err := server.ListenAndServe()

	logrus.Fatal(err)

}
