package main

import (
	"github.com/katsu105/hexagonal-todo/infrastructure/persistence"
	"github.com/katsu105/hexagonal-todo/usecase"
	"github.com/katsu105/hexagonal-todo/webserver/http/handlers"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	// データベース接続
	// dbConfig := &persistence.DatabaseConfig{
	// 	Host:     "localhost",
	// 	Port:     "3306",
	// 	User:     "user",
	// 	Password: "password",
	// 	DbName:   "todoapp",
	// }

	// db, err := persistence.ConnectToDatabase(dbConfig)
	// if err != nil {
	// 	panic(err)

	// }
	taskRepo := persistence.NewTaskRepository(nil)
	createTask := usecase.NewCreateTask(taskRepo)
	taskHandler := handlers.NewTaskHandler(createTask)

	e := echo.New()
	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// Routes
	e.POST("/", taskHandler.CreateTask)

	// Start server
	e.Logger.Fatal(e.Start(":1323"))
}
