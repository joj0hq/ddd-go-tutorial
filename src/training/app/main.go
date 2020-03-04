package main

import (
	"training/app/presentation"

	"training/app/infrastructure"

	"github.com/gin-gonic/gin"
	_ "github.com/mattn/go-sqlite3"
)

func main() {
	r := gin.Default()
	r.LoadHTMLGlob("templates/*.html")
	infrastructure.DbInit()
	presentation.TaskHandler(r)
	r.Run()
}
