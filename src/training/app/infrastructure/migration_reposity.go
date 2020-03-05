package infrastructure

import (
	"log"
	"training/app/domain"
	"training/app/utils"

	_ "github.com/go-sql-driver/mysql"
)

type MigrationRepository struct {
}

func NewMigrationRepository() *MigrationRepository {
	return &MigrationRepository{}
}

//DB初期化
func (t *MigrationRepository) Init() {
	db, err := utils.OpenDataBase()
	if err != nil {
		log.Fatalf("cannot connect database: %+v", err)
	}

	db.AutoMigrate(&domain.Task{})
	defer db.Close()
}
