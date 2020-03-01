package main

import (
	"training/app/presentation"

	"training/app/infrastructure"

	"github.com/gin-gonic/gin"
	_ "github.com/mattn/go-sqlite3"
)

func main() {
	r := gin.Default()
	const TemplateDir string = "templates/"
	r.LoadHTMLGlob(TemplateDir + "*.html")
	infrastructure.DbInit()
	presentation.TaskHandler(r)
	r.Run()
}
