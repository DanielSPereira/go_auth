package controllers

import (
	"auth/app/models"
	"auth/pkg/utils"
	"auth/platform/database"
	"strconv"
	"time"

	"github.com/gofiber/fiber/v2"
)

func SignIn(ctx *fiber.Ctx) error {
	signIn := &models.SignIn{}

	if err := ctx.BodyParser(signIn); err != nil {
		ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": true,
			"msg":   err.Error(),
		})
	}

	// create db connection
	db, err := database.MysqlConnection()
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": true,
			"msg":   err.Error(),
		})
	}

	// verify if user exists
	user, err := db.GetUserByEmail(signIn.Email)
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": true,
			"msg":   err.Error(),
		})
	}

	// compare password
	match := utils.CheckPasswordHash(signIn.Password, user.Password)
	if !match {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": true,
			"msg":   "Invalid password!",
		})
	}

	// generate JWT
	tokens, err := utils.GenerateTokens(strconv.Itoa(user.ID))
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": true,
			"msg":   err.Error(),
		})
	}

	return ctx.Status(fiber.StatusCreated).JSON(fiber.Map{
		"error": false,
		"msg":   "",
		"tokens": fiber.Map{
			"access":  tokens.AccessToken,
			"refresh": tokens.RefreshToken,
		},
		"user": fiber.Map{
			"id":    1,
			"email": signIn.Email,
		},
	})
}

func SignUp(ctx *fiber.Ctx) error {
	var err error

	signUp := &models.SignUp{}

	if err := ctx.BodyParser(signUp); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": true,
			"msg":   err.Error(),
		})
	}

	db, err := database.MysqlConnection()
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": true,
			"msg":   err.Error(),
		})
	}

	// verify if user exists - if err == nil (user exists)
	_, err = db.GetUserByEmail(signUp.Email)
	if err == nil {
		return ctx.Status(fiber.StatusConflict).JSON(fiber.Map{
			"error": true,
			"msg":   "User already exists!",
		})
	}

	// save user
	user := &models.User{Email: signUp.Email, Password: signUp.Password}
	if err = db.CreateUser(user); err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": true,
			"msg":   err.Error(),
		})
	}

	return ctx.Status(fiber.StatusCreated).JSON(fiber.Map{
		"error": false,
		"msg":   "",
	})
}

func CurrentUser(ctx *fiber.Ctx) error {
	// Get now time.
	now := time.Now().Unix()

	// Get claims from JWT.
	claims, err := utils.ExtractTokenMetadata(ctx)
	if err != nil {
		// Return status 500 and JWT parse error.
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": true,
			"msg":   err.Error(),
		})
	}

	// Set expiration time from JWT data of current book.
	expires := claims.Expires

	// Checking, if now time greather than expiration from JWT.
	if now >= expires {
		// Return status 401 and unauthorized error message.
		return ctx.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error": true,
			"msg":   "unauthorized, your token has expired",
		})
	}

	// Connect to database
	db, err := database.MysqlConnection()
	if err != nil {
		// Return status 500 and database connection error.
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": true,
			"msg":   err.Error(),
		})
	}

	// Get user from database
	user, err := db.GetUserById(claims.UserID)
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": true,
			"msg":   err.Error(),
		})
	}

	// Return status 200 OK.
	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
		"error": false,
		"msg":   "",
		"user":  user,
	})
}
