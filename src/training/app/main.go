package main

import (
	"training/app/presentation"

	"training/app/infrastructure"
	"training/app/usecase"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.LoadHTMLGlob("templates/*.html")
	taskRepo := infrastructure.NewTaskRepository()
	migrationRepo := infrastructure.NewMigrationRepository()
	taskService := usecase.NewTaskService(taskRepo)
	taskHandler := presentation.NewTaskHandler(taskService)
	taskHandler.SetupRouter(r.Group("/"))
	migrationRepo.Init()

	if err := r.Run(":8080"); err != nil {
		println("cannot start server: %+v", err)
	}
}
