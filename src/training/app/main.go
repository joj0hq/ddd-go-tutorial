package main

import (
	"training/app/presentation"

	"training/app/infrastructure"
	"training/app/usecase"

	"github.com/gin-gonic/gin"
	_ "github.com/mattn/go-sqlite3"
)

func main() {
	r := gin.Default()
	r.LoadHTMLGlob("templates/*.html")
	taskRepo := infrastructure.NewTaskRepository()
	taskService := usecase.NewTaskService(taskRepo)
	taskHandler := presentation.NewTaskHandler(taskService)
	taskHandler.SetupRouter(r.Group("/"))
	taskRepo.DbInit()

	if err := r.Run(":8080"); err != nil {
		println("cannot start server: %+v", err)
	}
}
