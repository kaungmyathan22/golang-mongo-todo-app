package authentication

import (
	"errors"
	"fmt"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/kaungmyathan22/golang-mongo-todo-app/src/features/authentication/schemas"
)

var validate *validator.Validate = validator.New()

func LoginHandler(c *fiber.Ctx) error {
	var payload schemas.LoginSchema
	if err := c.BodyParser(&payload); err != nil {
		if errors.Is(err, fiber.ErrUnprocessableEntity) {
			// Fields are invalid, return Unprocessable Entity
			return c.Status(fiber.StatusUnprocessableEntity).JSON(fiber.Map{
				"error": "Invalid fields in the request payload",
			})
		}
		return err
	}
	// err := validate.Struct(payload)
	// validationErrors := err.(validator.ValidationErrors)
	if errs := validate.Struct(payload); errs != nil {
		type ErrorResponse struct {
			Error       bool
			FailedField string
			Tag         string
			Value       interface{}
		}
		validationErrors := []ErrorResponse{}
		for _, err := range errs.(validator.ValidationErrors) {
			// In this case data object is actually holding the User struct
			var elem ErrorResponse
			elem.FailedField = err.Field() // Export struct field name
			elem.Tag = err.Tag()           // Export struct tag
			elem.Value = err.Value()       // Export field value
			elem.Error = true
			fmt.Println(err.Error())

			validationErrors = append(validationErrors, elem)
		}
		return c.JSON(validationErrors)
	}

	return c.JSON(map[string]string{"message": "User logging in"})
}

func RegisterHandler(c *fiber.Ctx) error {
	return c.JSON(map[string]string{"message": "Register endpoint"})
}

func LogoutHandler(c *fiber.Ctx) error {
	return c.JSON(map[string]string{"message": "Logout endpoint"})
}

func ForgotPasswordHandler(c *fiber.Ctx) error {
	return c.JSON(map[string]string{"message": "ForgotPassword endpoint"})
}

func ResetPasswordHandler(c *fiber.Ctx) error {
	return c.JSON(map[string]string{"message": "ResetPassword endpoint"})
}
