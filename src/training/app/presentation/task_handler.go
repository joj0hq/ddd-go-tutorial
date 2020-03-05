package presentation

import (
	"log"
	"strconv"
	"training/app/usecase"

	"github.com/gin-gonic/gin"
)

type TaskHandler struct {
	taskService usecase.TaskService
}

func NewTaskHandler(
	taskService usecase.TaskService,
) *TaskHandler {
	return &TaskHandler{
		taskService: taskService,
	}
}

func (t *TaskHandler) SetupRouter(r *gin.RouterGroup) {
	r.GET("", t.index)
	r.POST("/new", t.get)
	r.GET("/detail/:id", t.show)
	// TODO: ベタのHTMLなのでこのような実装になっているが本来は良くない (jojo on 2020/03/05)
	r.POST("/update/:id", t.update)
	r.GET("/delete_check/:id", t.deleteCheck)
	r.POST("/delete/:id", t.delete)
}

func (th *TaskHandler) index(ctx *gin.Context) {
	tasks := th.taskService.GetAll()
	ctx.HTML(200, "index.html", gin.H{
		"tasks": tasks,
	})
}

func (th *TaskHandler) get(ctx *gin.Context) {
	text := ctx.PostForm("text")
	status := ctx.PostForm("status")
	th.taskService.Create(text, status)
	ctx.Redirect(302, "/")
}

func (th *TaskHandler) show(ctx *gin.Context) {
	n := ctx.Param("id")
	id, err := strconv.Atoi(n)
	if err != nil {
		log.Fatalf("cannot show task: %+v", err)
	}
	task := th.taskService.GetById(id)
	ctx.HTML(200, "detail.html", gin.H{"task": task})
}

func (th *TaskHandler) update(ctx *gin.Context) {
	n := ctx.Param("id")
	id, err := strconv.Atoi(n)
	if err != nil {
		log.Fatalf("cannot update task: %+v", err)
	}
	text := ctx.PostForm("text")
	status := ctx.PostForm("status")
	th.taskService.Update(id, text, status)
	ctx.Redirect(302, "/")
}

func (th *TaskHandler) deleteCheck(ctx *gin.Context) {
	n := ctx.Param("id")
	id, err := strconv.Atoi(n)
	if err != nil {
		log.Fatalf("cannot delete check: %+v", err)
	}
	task := th.taskService.GetById(id)
	ctx.HTML(200, "delete.html", gin.H{"task": task})
}

func (th *TaskHandler) delete(ctx *gin.Context) {
	n := ctx.Param("id")
	id, err := strconv.Atoi(n)
	if err != nil {
		log.Fatalf("cannot delete task: %+v", err)
	}
	th.taskService.Delete(id)
	ctx.Redirect(302, "/")
}
