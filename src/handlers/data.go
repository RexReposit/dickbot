package handlers

import (
	"bot/src/controllers"
	"bot/src/models"
	"fmt"

	tele "gopkg.in/telebot.v4"
)

type Data struct {
	Db *controllers.DataController
}

func (d *Data) InitDB() {
	d.Db = controllers.NewDB()
	go d.Db.CronRun()
}

func (d *Data) Dick(c tele.Context) error {
	var user models.User
	var retMessage string

	id := c.Sender().ID
	lName := c.Sender().LastName
	fName := c.Sender().FirstName

	d.Db.DB.First(&user, id)

	newRange := controllers.RandRange()
	newSize := user.DickSize + newRange

	if user.IsBlocked {
		retMessage = fmt.Sprintf("%s %s, ты уже отращивал свой писюн, следющая попытка будет после 22:00 по МСК.", fName, lName)
	} else {
		if user.ID != id {
			d.Db.DB.Create(&models.User{ID: id, FirstName: fName, LastName: lName, DickSize: newSize, IsBlocked: false, IsAdmin: false})
		} else {
			d.Db.DB.Model(&user).Update("dick_size", newSize)
		}

		if newRange == 0 {
			retMessage = fmt.Sprintf("%s %s, твой писюн сегодня не подрос. Не переживай, просто он, видимо, решил взять выходной — пусть лучше отдохнет перед завтрашними подвигами!", fName, lName)
		} else if newRange > 0 {
			retMessage = fmt.Sprintf("%s %s, твой писюн сегодня подрос на %d см. Сейчас он равен %d см.", fName, lName, newRange, newSize)
		} else if newRange < 0 {
			retMessage = fmt.Sprintf("%s %s, твой писюн сегодня сократился на %d см. Сейчас он равен %d см.", fName, lName, -newRange, newSize)
		}
		d.Db.DB.Model(&user).Update("is_blocked", true)
	}

	return c.Send(retMessage)
}

func (d *Data) TopDick(c tele.Context) error {
	var users []models.User
	var retMessage string

	d.Db.DB.Order("dick_size desc").Limit(10).Find(&users)

	for iter, user := range users {
		retMessage += fmt.Sprintf("%d. %s %s, %d см\n", iter+1, user.FirstName, user.LastName, user.DickSize)
	}

	return c.Send(retMessage)
}

func (d *Data) ClearStatistics(c tele.Context) error {
	var user models.User
	d.Db.DB.First(&user, c.Sender().ID)

	if user.IsAdmin {
		d.Db.DB.Model(&models.User{}).Update("dick_size", 0)
		return c.Send("Статистика очищена!")
	}

	return c.Send("Ты не админ!")
}
