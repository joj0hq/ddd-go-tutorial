package presentation

import (
	"strconv"
	"training/app/infrastructure"

	"github.com/gin-gonic/gin"
)

func TaskHandler(router *gin.Engine) {
	//Index
	router.GET("/", func(ctx *gin.Context) {
		tasks := infrastructure.GetAll()
		ctx.HTML(200, "index.html", gin.H{
			"tasks": tasks,
		})
	})
	//Create
	router.POST("/new", func(ctx *gin.Context) {
		text := ctx.PostForm("text")
		status := ctx.PostForm("status")
		infrastructure.Create(text, status)
		ctx.Redirect(302, "/")
	})
	//Detail
	router.GET("/detail/:id", func(ctx *gin.Context) {
		n := ctx.Param("id")
		id, err := strconv.Atoi(n)
		if err != nil {
			panic(err)
		}
		task := infrastructure.GetById(id)
		ctx.HTML(200, "detail.html", gin.H{"task": task})
	})
	//Update
	router.POST("/update/:id", func(ctx *gin.Context) {
		n := ctx.Param("id")
		id, err := strconv.Atoi(n)
		if err != nil {
			panic("ERROR")
		}
		text := ctx.PostForm("text")
		status := ctx.PostForm("status")
		infrastructure.Update(id, text, status)
		ctx.Redirect(302, "/")
	})
	//削除確認
	router.GET("/delete_check/:id", func(ctx *gin.Context) {
		n := ctx.Param("id")
		id, err := strconv.Atoi(n)
		if err != nil {
			panic("ERROR")
		}
		task := infrastructure.GetById(id)
		ctx.HTML(200, "delete.html", gin.H{"task": task})
	})
	//Delete
	router.POST("/delete/:id", func(ctx *gin.Context) {
		n := ctx.Param("id")
		id, err := strconv.Atoi(n)
		if err != nil {
			panic("ERROR")
		}
		infrastructure.Delete(id)
		ctx.Redirect(302, "/")
	})
}
