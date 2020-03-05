package main

import (
	"log"
	"training/app/infrastructure"
	"training/app/presentation"
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
		log.Fatalf("cannot start server: %+v", err)
	}
}
