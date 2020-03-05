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

// 新規タスクを作成する
func (ts *taskService) Create(text string, status string) {
	ts.TaskRepos.Create(text, status)
}

// タスクを更新する
func (ts *taskService) Update(id int, text string, status string) {
	ts.TaskRepos.Update(id, text, status)
}

// タスクを更新する
func (ts *taskService) Delete(id int) {
	ts.TaskRepos.Delete(id)
}

// タスクを全取得する
func (ts *taskService) GetAll() []domain.Task {
	return ts.TaskRepos.GetAll()
}

// タスクを1件取得する
func (ts *taskService) GetById(id int) domain.Task {
	return ts.TaskRepos.GetById(id)
}
