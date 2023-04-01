package main

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	e := echo.New()
	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// Routes
	e.GET("/", hello)

	// Start server
	e.Logger.Fatal(e.Start(":1323"))
}

// ハンドラーを定義
func hello(c echo.Context) error {
	return c.String(http.StatusOK, "Hello, World!")
}

/*

todo-app/
  ├─ cmd/
  │   └─ main.go
  ├─ domain/
  │   ├─ entities/
  │   │   └─ task.go
  │   └─ repository/
  │       └─ task_repository.go
  ├─ usecases/
  │   ├─ create_task.go
  │   ├─ list_tasks.go
  │   └─ update_task.go
  ├─ infrastructure/
  │   ├─ persistence/
  │   │   ├─ in_memory.go
  │   │   └─ database.go
  │   └─ webserver/
  │       ├─ http/
  │       │   ├─ server.go
  │       │   └─ task_handler.go
  │       └─ grpc/
  │           ├─ server.go
  │           └─ task_handler.go
  └─ config/
      └─ config.go


*/
