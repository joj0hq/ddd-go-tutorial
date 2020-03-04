package infrastructure

import (
	"training/app/domain"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

type MigrationRepository struct {
}

func NewMigrationRepository() *MigrationRepository {
	return &MigrationRepository{}
}

var (
	DBMS     = "mysql"
	USER     = "root"
	PASS     = "password"
	PROTOCOL = "tcp(ddd-training-mysql:3306)"
	DBNAME   = "gopher"
	CONNECT  = USER + ":" + PASS + "@" + PROTOCOL + "/" + DBNAME + "?parseTime=true"
)

//DB初期化
func (t *MigrationRepository) Init() {
	db, err := gorm.Open(DBMS, CONNECT)
	if err != nil {
		panic("データベース開けず！（dbInit）")
	}
	db.AutoMigrate(&domain.Task{})
	defer db.Close()
}
