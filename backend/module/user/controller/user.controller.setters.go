package user_controller

import (
	"github.com/gofiber/fiber/v2"
	user_dto "github.com/root9464/Ton-students/module/user/dto"
)

func (c *userController) AddUserInfo(ctx *fiber.Ctx) error {
	userInfo := new(user_dto.UserInfoType)
	if err := ctx.BodyParser(userInfo); err != nil {
		return &fiber.Error{
			Code:    400,
			Message: err.Error(),
		}
	}

	if err := c.userService.AddUserInfo(ctx.Context(), userInfo); err != nil {
		return &fiber.Error{
			Code:    500,
			Message: err.Error(),
		}
	}

	return ctx.Status(200).JSON(&fiber.Map{
		"status":  "success",
		"message": "User info added successfully",
	})
}

func (c *userController) SelectVisibleName(ctx *fiber.Ctx) error {
	visibleName := new(user_dto.SelectVisibleNameType)
	if err := ctx.BodyParser(visibleName); err != nil {
		return &fiber.Error{
			Code:    400,
			Message: err.Error(),
		}
	}

	if err := c.userService.SelectVisibleName(ctx.Context(), visibleName); err != nil {
		return &fiber.Error{
			Code:    500,
			Message: err.Error(),
		}
	}

	return ctx.Status(200).JSON(&fiber.Map{
		"status":  "success",
		"message": "User visible name selected successfully",
	})
}

func (c *userController) SetUserNickname(ctx *fiber.Ctx) error {
	nickname := new(user_dto.SetUserNicknameType)
	if err := ctx.BodyParser(nickname); err != nil {
		return &fiber.Error{
			Code:    400,
			Message: err.Error(),
		}
	}

	if err := c.userService.SetUserNickname(ctx.Context(), nickname); err != nil {
		return &fiber.Error{
			Code:    500,
			Message: err.Error(),
		}
	}

	return ctx.Status(200).JSON(&fiber.Map{
		"status":  "success",
		"message": "User nickname set successfully",
	})
}

func (c *userController) UpdateUserInfo(ctx *fiber.Ctx) error {
	userInfo := new(user_dto.UpdateUserInfoType)
	if err := ctx.BodyParser(userInfo); err != nil {
		return &fiber.Error{
			Code:    400,
			Message: err.Error(),
		}
	}

	if err := c.userService.UpdateUserInfo(ctx.Context(), userInfo); err != nil {
		return &fiber.Error{
			Code:    500,
			Message: err.Error(),
		}
	}

	return ctx.Status(200).JSON(&fiber.Map{
		"status":  "success",
		"message": "User info updated successfully",
	})
}
