package lead

import (
	"github.com/go-fiber-demo/databases"
	"github.com/gofiber/fiber"
	"github.com/jinzhu/gorm"
)

type Lead struct {
	gorm.Model
	Name    string `json:"name,omitempty"`
	Phone   int64  `json:"phone,omitempty"`
	Company string `json:"company,omitempty"`
	Email   string `json:"email,omitempty"`
}

func GetLeads(c *fiber.Ctx) {
	db := databases.DBConn
	var leads []Lead
	db.Find(&leads)
	c.JSON(leads)
}

func GetLead(c *fiber.Ctx) {
	id := c.Params("id")
	db := databases.DBConn
	var lead Lead
	db.Find(&lead, id)
	c.JSON(lead)
}

func NewLead(c *fiber.Ctx) {
	lead := new(Lead)
	db := databases.DBConn
	err := c.BodyParser(lead)
	if err != nil {
		c.Status(503).Send(err)
		return
	}
	db.Create(&lead)
	c.JSON(lead)
}

func DeleteLead(c *fiber.Ctx) {
	id := c.Params("id")
	db := databases.DBConn
	var lead Lead
	db.First(&lead, id)
	if lead.Name == "" {
		c.Status(500).Send("No lead foudn with ID")
		return
	}
	db.Unscoped().Delete(&lead)
	c.Send("Deleted Successfully")
}
