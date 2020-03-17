package repository

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	"gitlab.com/pragmaticreviews/golang-gin-poc/entity"
)

type VideoRepository interface {
	Save(video entity.Video) error
	Update(video entity.Video) error
	Delete(video entity.Video) error
	FindAll() []entity.Video
	CloseDB()
}

type database struct {
	connection *gorm.DB
}

func NewVideoRepository() VideoRepository {
	db, err := gorm.Open("sqlite3", "test.db")
	if err != nil {
		panic("failed to connect database")
	}
	db.AutoMigrate(&entity.Video{}, &entity.Person{})
	return &database{
		connection: db,
	}
}

func (db *database) CloseDB() {
	err := db.connection.Close()
	if err != nil {
		panic("failed to close database")
	}
}

func (db *database) Save(video entity.Video) error {
	db.connection.Create(&video)
	db.connection.NewRecord(video)
	return nil
}

func (db *database) Update(video entity.Video) error {
	db.connection.Save(&video)
	return nil
}

func (db *database) Delete(video entity.Video) error {
	db.connection.Delete(&video)
	return nil
}

func (db *database) FindAll() []entity.Video {
	var videos []entity.Video
	db.connection.Set("gorm:auto_preload", true).Find(&videos)
	return videos
}
