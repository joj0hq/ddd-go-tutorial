package infrastructure

import (
	"github.com/jinzhu/gorm"
	"training/app/domain"
)

//DB初期化
func DbInit() {
	db, err := gorm.Open("sqlite3", "test.sqlite3")
	if err != nil {
		panic("データベース開けず！（dbInit）")
	}
	db.AutoMigrate(&domain.Task{})
	defer db.Close()
}

//DB追加
func DbInsert(text string, status string) {
	db, err := gorm.Open("sqlite3", "test.sqlite3")
	if err != nil {
		panic("データベース開けず！（dbInsert)")
	}
	db.Create(&domain.Task{Text: text, Status: status})
	defer db.Close()
}

//DB更新
func DbUpdate(id int, text string, status string) {
	db, err := gorm.Open("sqlite3", "test.sqlite3")
	if err != nil {
		panic("データベース開けず！（dbUpdate)")
	}
	var todo domain.Task
	db.First(&todo, id)
	todo.Text = text
	todo.Status = status
	db.Save(&todo)
	db.Close()
}

//DB削除
func DbDelete(id int) {
	db, err := gorm.Open("sqlite3", "test.sqlite3")
	if err != nil {
		panic("データベース開けず！（dbDelete)")
	}
	var todo domain.Task
	db.First(&todo, id)
	db.Delete(&todo)
	db.Close()
}

//DB全取得
func DbGetAll() []domain.Task {
	db, err := gorm.Open("sqlite3", "test.sqlite3")
	if err != nil {
		panic("データベース開けず！(dbGetAll())")
	}
	var todos []domain.Task
	db.Order("created_at desc").Find(&todos)
	db.Close()
	return todos
}

//DB一つ取得
func DbGetOne(id int) domain.Task {
	db, err := gorm.Open("sqlite3", "test.sqlite3")
	if err != nil {
		panic("データベース開けず！(dbGetOne())")
	}
	var todo domain.Task
	db.First(&todo, id)
	db.Close()
	return todo
}
