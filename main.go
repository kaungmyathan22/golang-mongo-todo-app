package main

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/kaungmyathan22/golang-mongo-todo-app/src/features/authentication"
	"github.com/kaungmyathan22/golang-mongo-todo-app/src/features/item"
	"github.com/kaungmyathan22/golang-mongo-todo-app/src/features/user"
)

func main() {
	app := fiber.New(fiber.Config{
		AppName: "Golang todo app",
	})

	//#region------------------- group routes
	publicRoutes := app.Group("/api/v1/authentication")
	itemRoutes := app.Group("/api/v1/item")
	userRoutes := app.Group("/api/v1/user")

	//#endregion
	//authentication related routes
	publicRoutes.Post("/login", authentication.LoginHandler)
	publicRoutes.Post("/register", authentication.RegisterHandler)
	publicRoutes.Post("/logout", authentication.LogoutHandler)
	publicRoutes.Post("/forgot-password", authentication.ForgotPasswordHandler)
	publicRoutes.Post("/reset-password", authentication.ResetPasswordHandler)

	//#region------------------- item routes
	itemRoutes.Get("/", item.GetItems)
	itemRoutes.Post("/", item.CreateItem)

	itemRoutes.Get("/:id", item.GetItem)
	itemRoutes.Patch("/:id", item.UpdateItem)
	itemRoutes.Delete("/:id", item.RemoveItemById)
	//#endregion

	//#region------------------- user routes
	userRoutes.Get("/", user.GetUsers)
	userRoutes.Post("/", user.CreateUser)

	userRoutes.Get("/:id", user.GetUser)
	userRoutes.Patch("/:id", user.UpdateUser)
	userRoutes.Delete("/:id", user.RemoveUserById)
	//#endregion
	port := "3000"
	err := app.Listen(":3000")
	if err != nil {
		fmt.Printf("something went wrong %s", err)
	}
	fmt.Printf("server is running at %s", port)
}
