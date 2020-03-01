package presentation

import (
	"github.com/gin-gonic/gin"
	"strconv"
	"training/app/infrastructure"
)

func TaskHandler(router *gin.Engine) {
	//Index
	router.GET("/", func(ctx *gin.Context) {
		todoList := infrastructure.DbGetAll()
		ctx.HTML(200, "index.html", gin.H{
			"todoList": todoList,
		})
	})
	//Create
	router.POST("/new", func(ctx *gin.Context) {
		text := ctx.PostForm("text")
		status := ctx.PostForm("status")
		infrastructure.DbInsert(text, status)
		ctx.Redirect(302, "/")
	})
	//Detail
	router.GET("/detail/:id", func(ctx *gin.Context) {
		n := ctx.Param("id")
		id, err := strconv.Atoi(n)
		if err != nil {
			panic(err)
		}
		todo := infrastructure.DbGetOne(id)
		ctx.HTML(200, "detail.html", gin.H{"todo": todo})
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
		infrastructure.DbUpdate(id, text, status)
		ctx.Redirect(302, "/")
	})
	//削除確認
	router.GET("/delete_check/:id", func(ctx *gin.Context) {
		n := ctx.Param("id")
		id, err := strconv.Atoi(n)
		if err != nil {
			panic("ERROR")
		}
		todo := infrastructure.DbGetOne(id)
		ctx.HTML(200, "delete.html", gin.H{"todo": todo})
	})
	//Delete
	router.POST("/delete/:id", func(ctx *gin.Context) {
		n := ctx.Param("id")
		id, err := strconv.Atoi(n)
		if err != nil {
			panic("ERROR")
		}
		infrastructure.DbDelete(id)
		ctx.Redirect(302, "/")
	})
}
