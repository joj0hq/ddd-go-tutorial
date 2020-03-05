package infrastructure

import (
	"log"
	"training/app/domain"
	"training/app/utils"

	_ "github.com/go-sql-driver/mysql"
)

type TaskRepository struct {
}

func NewTaskRepository() *TaskRepository {
	return &TaskRepository{}
}

func (t *TaskRepository) Create(text string, status string) {
	db, err := utils.OpenDataBase()
	if err != nil {
		log.Fatalf("cannot connect database: %+v", err)
	}
	db.Create(&domain.Task{Text: text, Status: status})
	defer db.Close()
}

func (t *TaskRepository) Update(id int, text string, status string) {
	db, err := utils.OpenDataBase()
	if err != nil {
		log.Fatalf("cannot connect database: %+v", err)
	}
	var todo domain.Task
	db.First(&todo, id)
	todo.Text = text
	todo.Status = status
	db.Save(&todo)
	db.Close()
}

func (t *TaskRepository) Delete(id int) {
	db, err := utils.OpenDataBase()
	if err != nil {
		log.Fatalf("cannot connect database: %+v", err)
	}
	var todo domain.Task
	db.First(&todo, id)
	db.Delete(&todo)
	db.Close()
}

func (t *TaskRepository) GetAll() []domain.Task {
	db, err := utils.OpenDataBase()
	if err != nil {
		log.Fatalf("cannot connect database: %+v", err)
	}
	var tasks []domain.Task
	db.Order("created_at desc").Find(&tasks)
	db.Close()
	return tasks
}

func (t *TaskRepository) GetById(id int) domain.Task {
	db, err := utils.OpenDataBase()
	if err != nil {
		log.Fatalf("cannot connect database: %+v", err)
	}
	var todo domain.Task
	db.First(&todo, id)
	db.Close()
	return todo
}
