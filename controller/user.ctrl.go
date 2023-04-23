package controller

import (
	"fmt"
	"gotest/models"
	"gotest/utils"

	"github.com/gofiber/fiber/v2"
)

type UserController struct {
	UserModel *models.UserModel
}

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
	var user models.User
	if err := c.BodyParser(&user); err != nil {
		return fiber.NewError(500, "check name")
	}
	return c.Status(fiber.StatusOK).JSON(map[string]interface{}{
		"result": "success",
	})
}

func (u *UserController) Insert(c *fiber.Ctx) error {
	user := new(models.User)
	if err := c.BodyParser(user); err != nil {
		c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "failed to parse JSON",
		})
		return err
	}
	userid := user.Id
	name := user.Name
	nick := user.Nick
	pass := user.Password

	err := u.UserModel.Insert(user, userid, name, nick, pass)
	if err != nil {
		c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "failed to insert",
		})
		return err
	}

	return c.JSON(fiber.Map{
		"result": 1,
	})
}

func (u *UserController) Select(c *fiber.Ctx) error {
	user := new(models.User)
	err := c.BodyParser(&user)
	fmt.Println(user)
	if err != nil {
		c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "cannot parse json",
		})
		return err
	}
	users, err := u.UserModel.Select(user.Name)
	if err != nil {
		c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "cannot parse json",
		})
		return err
	}
	return c.JSON(users)
}

func (u *UserController) Delete(c *fiber.Ctx) error {
	user := new(models.User)
	err2 := c.BodyParser(user)
	fmt.Println(err2)
	if err2 != nil {

		c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "cannot parse json",
		})
		return nil
	}
	if user.No <= 0 {
		c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "invalid user no",
		})
		return nil
	}
	err := u.UserModel.Delete(user.No)
	fmt.Println(err)
	if err != nil {
		c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "cannot parse json",
		})
		return nil
	}
	return c.JSON(map[string]interface{}{
		"result": 1,
	})
}

func (u *UserController) Update(c *fiber.Ctx) error {
	// 분해해서 아이디값
	id, err := utils.VarifiyToken(c)
	fmt.Println(id)
	if err != nil {
		c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error": err.Error(),
		})
		return err
	}

	user := new(models.User)
	if err := c.BodyParser(user); err != nil {
		c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "cannot parse json",
		})
		return err
	}

	// 분해해서 얻은 아이디로 업데이트
	user.Id = id
	if err := u.UserModel.Update(user); err != nil {
		c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
		return err
	}

	return c.JSON(map[string]interface{}{
		"result": 1,
	})
}

func (u *UserController) Login(c *fiber.Ctx) error {
	user := new(models.User)
	err := c.BodyParser(&user)
	if err != nil {
		c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "cannot parse json",
		})
		return err
	}
	token, refreshtoken, err := u.UserModel.Login(user.Id, user.Password)
	if err != nil {
		c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "cannot parse json",
		})
		return err
	}
	return c.Status(fiber.StatusAccepted).JSON(fiber.Map{
		"token":        token,
		"refreshtoken": refreshtoken,
	})
}
