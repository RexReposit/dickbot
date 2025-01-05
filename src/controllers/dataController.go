package controllers

import (
	"bot/src/models"
	"math/rand"
	"time"

	"github.com/glebarez/sqlite"
	"github.com/robfig/cron"
	"gorm.io/gorm"
)

type DataController struct {
	DB *gorm.DB
}

func NewDB() *DataController {
	db, err := gorm.Open(sqlite.Open("users.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	db.AutoMigrate(models.User{})
	return &DataController{DB: db}
}

func RandRange() int {
	rand.Seed(time.Now().UnixNano())
	return rand.Intn(21) - 5
}

func (d *DataController) ClearAllBlocked() {
	d.DB.Model(&models.User{}).Where("is_blocked = ?", true).Update("is_blocked", false)
}

func (d *DataController) CronRun() {
	c := cron.New()

	c.AddFunc("0 20 * * *", func() {
		d.ClearAllBlocked()
	})

	c.Start()
}
