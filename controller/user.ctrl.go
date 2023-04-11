package controller

import (
	"github.com/gofiber/fiber/v2"
)

type UserController struct{}

func (u *UserController) Hello(c *fiber.Ctx) error {
	return c.SendString("hello world!")
}

func (u *UserController) Error(c *fiber.Ctx) error {
	return fiber.NewError(500, "Custom Error Message")
}

func (u *UserController) Params(c *fiber.Ctx) error {
	return c.JSON(c.Params("name"))
}
func (u *UserController) Header(c *fiber.Ctx) error {
	return c.JSON(c.GetReqHeaders())
}

func (u *UserController) Post(c *fiber.Ctx) error {
	var user, nick models.User
	if err := c.BodyParser(&user); err != nil {
		return fiber.NewError(500, "check name")
	}
	if err := c.BodyParser(&nick); err != nil {
		return fiber.NerError(500, "check nick")
	}
	return models.Post(user)
}

func (u *UserContorller) Insert(c *fiber.Ctx) error {
	user := new(models.User)
	if err := c.BodyParser(user); err != nil {
		return err
	}
	return models.Insert(user)
}

func (u *UserController) Select(c *fiber.Ctx) error {
	user := new(models.User)
	if err := c.BodyParser(user); err != nil {
		return err
	}
	return models.Select(users)
}

func (u *UserController) Delete(c *fiber.Ctx) error {
	user := new(models.User)
	if err := c.BodyParser(user); err != nil {
		return err
	}
	return models.Delete(users.No)
}

func (u *UserController) Update(c *fiber.Ctx) error {
	user := new(models.User)
	if err := c.BodyParser(user); err != nil {
		return err
	}
	return models.Update(user.Nick, user.No)
}
