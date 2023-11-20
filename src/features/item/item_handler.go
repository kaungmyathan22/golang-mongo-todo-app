package item

import "github.com/gofiber/fiber/v2"

func CreateItem(c *fiber.Ctx) error {
	return c.JSON(map[string]string{"message": "CreateItem endpoint"})
}

func UpdateItem(c *fiber.Ctx) error {
	return c.JSON(map[string]string{"message": "UpdateItem endpoint"})
}

func GetItems(c *fiber.Ctx) error {
	return c.JSON(map[string]string{"message": "GetItems endpoint"})
}

func GetItem(c *fiber.Ctx) error {
	return c.JSON(map[string]string{"message": "GetItem endpoint"})
}

func RemoveItemById(c *fiber.Ctx) error {
	return c.JSON(map[string]string{"message": "RemoveItemById endpoint"})
}
