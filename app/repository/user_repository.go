package repository

import (
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/kooroshh/fiber-boostrap/app/models"
	"github.com/kooroshh/fiber-boostrap/pkg/database"
)

func InsertNewUser(ctx *fiber.Ctx, user *models.User) error {
	return database.DB.Create(user).Error
}

func InsertNewUserSession(ctx *fiber.Ctx, session *models.UserSession) error {
	return database.DB.Create(session).Error
}

func GetUserSessionByToken(ctx *fiber.Ctx, token string) (models.UserSession, error) {
	var (
		resp models.UserSession
		err  error
	)

	err = database.DB.Where("token = ?", token).Last(&resp).Error
	return resp, err
}

func DeleteUserSessionByToken(ctx *fiber.Ctx, token string) error {
	return database.DB.Exec("DELETE FROM user_sessions WHERE token = ?", token).Error
}

func GetUser(ctx *fiber.Ctx, username string) (models.User, error) {
	var (
		resp models.User
		err  error
	)

	err = database.DB.Where("username = ?", username).Last(&resp).Error
	return resp, err
}

func UpdateUserSessionToken(ctx *fiber.Ctx, token string, tokenExpired time.Time, refreshToken string) error {
	return database.DB.Exec("UPDATE user_sessions SET token = ?, token_expired = ? WHERE refresh_token = ?", token, tokenExpired, refreshToken).Error
}
