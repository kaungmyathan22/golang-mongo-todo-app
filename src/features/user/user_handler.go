package user

import "github.com/gofiber/fiber/v2"

func CreateUser(c *fiber.Ctx) error {
	return c.JSON(map[string]string{"message": "CreateUser"})
}

func UpdateUser(c *fiber.Ctx) error {
	return c.JSON(map[string]string{"message": "UpdateUser"})
}

func GetUsers(c *fiber.Ctx) error {
	return c.JSON(map[string]string{"message": "GetUsers"})
}

func GetUser(c *fiber.Ctx) error {
	return c.JSON(map[string]string{"message": "GetUser"})
}

func RemoveUserById(c *fiber.Ctx) error {
	return c.JSON(map[string]string{"message": "RemoveUserById"})
}
