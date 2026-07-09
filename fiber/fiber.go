package fiber

import (
	"regexp"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v3"
)

var validate = validator.New()

type User struct {
	Email    string `json:"email" validate:"required,email"`
	Fullname string `json:"fullname" validate:"required,fullname"`
	Age      int    `json:"age" validate:"required,numeric,min=1"`
}

// Setup builds the Fiber app with routes and validation registered.
func Setup() *fiber.App {
	app := fiber.New()

	validate.RegisterValidation("fullname", validateFullname)

	app.Post("/users", func(c fiber.Ctx) error {
		user := new(User)

		if err := c.Bind().Body(&user); err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Can't parse JSON"})
		}

		if err := validate.Struct(user); err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
		}

		return c.Status(fiber.StatusOK).JSON(user)
	})

	return app
}

func validateFullname(fl validator.FieldLevel) bool {
	return regexp.MustCompile(`^[a-zA-Z\s]+$`).MatchString(fl.Field().String())
}
