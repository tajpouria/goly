package repository

import (
	"fmt"

	"github.com/go-playground/validator/v10"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var Model Repository
var Validate *validator.Validate

type Goly struct {
	ID       uint64 `json:"id" gorm:"column:id;primaryKey"`
	Redirect string `json:"redirect" gorm:"column:redirect;not null" validate:"required,url"`
	Goly     string `json:"goly" gorm:"column:goly;unique;not null" validate:"url_encoded"`
	Random   bool  `json:"random" gorm:"column:random"`
	Clicked  bool  `json:"clicked" gorm:"column:clicked"`
}

func Setup() {
	dsn := "host=localhost user=postgres password=postgres dbname=postgres port=5432 sslmode=disable"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	err = db.AutoMigrate(&Goly{})
	if err != nil {
		fmt.Println(err)
	}

	Model = NewRepository(db)
	Validate = validator.New()
}
