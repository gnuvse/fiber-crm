package main

import (
	"fmt"

	"github.com/gnuvse/fiber-crm/database"
	"github.com/gnuvse/fiber-crm/lead"
	"github.com/gofiber/fiber"
	"github.com/jinzhu/gorm"
)

func setupRouters(app *fiber.App) {
	app.Get("/api/v1.lead", lead.GetLeads)
	app.Get("/api/v1/lead/:id", lead.GetLead)
	app.Post("/api/v1/lead", lead.NewLead)
	app.Delete("api/v1/lead/:id", lead.DeleteLead)
}

func initDatabase() {
	var err error
	database.DBconn, err = gorm.Open("sqlite3", "leads.db")
	if err != nil {
		panic("failed to connect database")
	}
	fmt.Println("Connection opened to database")
	database.DBconn.AutoMigrate(&lead.Lead{})
	fmt.Println("Database migrated")

}

func main() {
	app := fiber.New()
	initDatabase()
	setupRouters(app)
	// listen port
	app.Listen(3000)
	defer database.DBconn.Close()

}
