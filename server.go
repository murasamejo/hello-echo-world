package main

import (
	"fmt"
	"net/http"
	"time"

	"github.com/labstack/echo/v4"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

// TODO: string or *string
// gorm.Model を宣言すると deleted_at が自動で注入される（されてしまう）
type FakeMaterial struct {
	ID                      uint `gorm:"primaryKey"`
	Name                    string
	Email                   string
	Credit_card_number      string
	Credit_card_expiry_date time.Time
	Credit_card_type        string
	City                    string
	Gender                  string
	MyNumber                uint
	Created_at              time.Time
	Updated_at              time.Time
}

func main() {
	db, err := gorm.Open(sqlite.Open("development.sqlite3"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	var fake_material FakeMaterial

	db.First(&fake_material)
	fmt.Println(fake_material)

	e := echo.New()

	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, Echo World!")
	})
	e.GET("/users/:id", getUser)

	e.Logger.Fatal(e.Start(":1323"))
}

func getUser(c echo.Context) error {
	db, err := gorm.Open(sqlite.Open("development.sqlite3"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	var fake_material FakeMaterial

	id := c.Param("id")
	db.First(&fake_material, id)
	return c.String(http.StatusOK, fake_material.Email)
}
