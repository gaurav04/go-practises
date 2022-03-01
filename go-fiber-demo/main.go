package main

import (
	"github.com/go-fiber-demo/databases"
	"github.com/go-fiber-demo/lead"
	"github.com/gofiber/fiber"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

func setupRoutes(app *fiber.App) {
	app.Get("/api/v1/lead", lead.GetLeads)
	app.Get("/api/v1/lead/:id", lead.GetLead)
	app.Post("/api/v1/lead", lead.NewLead)
	app.Delete("/api/v1/lead/:id", lead.DeleteLead)
}

func Connect() {
	var err error
	databases.DBConn, err = gorm.Open("mysql", "root:12345678@/leads?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		panic(err)
	}
	databases.DBConn.AutoMigrate(&lead.Lead{})
}

func main() {
	app := fiber.New()
	Connect()
	setupRoutes(app)
	app.Listen(4000)
	//defer databases.DBConn.Close()

}
