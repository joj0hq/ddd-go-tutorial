package repository

import "training/app/domain"

type TaskRepository interface {
	DbInit()
	Create(text string, status string)
	Update(id int, text string, status string)
	Delete(id int)
	GetAll() []domain.Task
	GetById(id int) domain.Task
}
