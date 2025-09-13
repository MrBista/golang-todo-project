package app

import (
	"github.com/MrBista/golang-todo-project/src/controllers"
	"github.com/MrBista/golang-todo-project/src/exception"
	"github.com/MrBista/golang-todo-project/src/middleware"
	"github.com/julienschmidt/httprouter"
)

func authRouter(router *httprouter.Router, userController controllers.UserController) {
	router.POST("/api/v1/auth/register", userController.UserRegister)
	router.POST("/api/v1/auth/login", userController.LoginUser)

}
func todoRouter(router *httprouter.Router, todoController controllers.TodoController) {
	router.POST("/api/v1/todos", middleware.AutthMiddlware(todoController.CreateTodo))
	router.PUT("/api/v1/todos/:todoId", middleware.AutthMiddlware(todoController.UpdateTodo))
	router.GET("/api/v1/todos/:todoId", middleware.AutthMiddlware(todoController.FindByIdTodo))
	router.GET("/api/v1/todos", middleware.AutthMiddlware(todoController.FindAllTodo))
	router.DELETE("/api/v1/todos/:todoId", middleware.AutthMiddlware(todoController.DeleteByIdTodo))
}

func NewRouter(userController controllers.UserController, todoController controllers.TodoController) *httprouter.Router {
	router := httprouter.New()

	// route todo
	authRouter(router, userController)
	todoRouter(router, todoController)

	router.PanicHandler = exception.ErrorHandler
	return router
}
