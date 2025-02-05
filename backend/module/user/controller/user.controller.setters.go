package user_controller

import (
	"github.com/gofiber/fiber/v2"
	user_dto "github.com/root9464/Ton-students/module/user/dto"
	"github.com/root9464/Ton-students/shared/utils"
)

func (c *userController) AddUserInfo(ctx *fiber.Ctx) error {
	userInfo := new(user_dto.UserCreateInfoType)
	if err := ctx.BodyParser(userInfo); err != nil {
		return &fiber.Error{
			Code:    400,
			Message: err.Error(),
		}
	}

	if err := c.userService.AddUserInfo(ctx.Context(), userInfo); err != nil {
		if errorResponse, code := utils.HandlerError(err); errorResponse != nil {
			return ctx.Status(code).JSON(errorResponse)
		}
	}

	return ctx.Status(200).JSON(&fiber.Map{
		"status":  "success",
		"message": "User info added successfully",
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
		if errorResponse, code := utils.HandlerError(err); errorResponse != nil {
			return ctx.Status(code).JSON(errorResponse)
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
		if errorResponse, code := utils.HandlerError(err); errorResponse != nil {
			return ctx.Status(code).JSON(errorResponse)
		}
	}

	return ctx.Status(200).JSON(&fiber.Map{
		"status":  "success",
		"message": "User info updated successfully",
	})
}

func (c *userController) DeleteUserInfo(ctx *fiber.Ctx) error {
	userInfo := ctx.Query("id")
	if userInfo == "" {
		return &fiber.Error{
			Code:    400,
			Message: "id is required",
		}
	}

	if err := c.userService.DeleteUserInfo(ctx.Context(), userInfo); err != nil {
		if errorResponse, code := utils.HandlerError(err); errorResponse != nil {
			return ctx.Status(code).JSON(errorResponse)
		}
	}

	return ctx.Status(200).JSON(&fiber.Map{
		"status":  "success",
		"message": "User info deleted successfully",
	})
}

func (c *userController) AddManyUserInfo(ctx *fiber.Ctx) error {
	userInfos := new(user_dto.ManyUserInfoType)
	if err := ctx.BodyParser(userInfos); err != nil {
		return &fiber.Error{
			Code:    400,
			Message: err.Error(),
		}
	}

	if err := c.userService.AddManyUserInfo(ctx.Context(), userInfos); err != nil {
		if errorResponse, code := utils.HandlerError(err); errorResponse != nil {
			return ctx.Status(code).JSON(errorResponse)
		}
	}

	return ctx.Status(200).JSON(&fiber.Map{
		"status":  "success",
		"message": "User info added successfully",
	})
}
