package usecase

import (
	"training/app/domain"
	"training/app/domain/repository"

	_ "github.com/go-sql-driver/mysql"
)

type TaskService interface {
	Create(text string, status string)
	Update(id int, text string, status string)
	Delete(id int)
	GetAll() []domain.Task
	GetById(id int) domain.Task
}

type taskService struct {
	TaskRepos repository.TaskRepository
}

func NewTaskService(
	taskRepos repository.TaskRepository) *taskService {
	return &taskService{
		TaskRepos: taskRepos,
	}
}

//DB初期化
func (ts *taskService) Init() {
	ts.TaskRepos.Init()
}

//DB追加
func (ts *taskService) Create(text string, status string) {
	ts.TaskRepos.Create(text, status)
}

//DB更新
func (ts *taskService) Update(id int, text string, status string) {
	ts.TaskRepos.Update(id, text, status)
}

//DB削除
func (ts *taskService) Delete(id int) {
	ts.TaskRepos.Delete(id)
}

//DB全取得
func (ts *taskService) GetAll() []domain.Task {
	return ts.TaskRepos.GetAll()
}

//DB一つ取得
func (ts *taskService) GetById(id int) domain.Task {
	return ts.TaskRepos.GetById(id)
}
