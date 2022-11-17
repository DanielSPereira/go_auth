package utils

import (
	"os"
	"strconv"
	"strings"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v4"
)

type RequestTokenMetadata struct {
	UserID  int
	Expires int64
}

func ExtractTokenMetadata(ctx *fiber.Ctx) (*RequestTokenMetadata, error) {
	// get parsedBearerToken from context headers
	parsedBearerToken := getParsedToken(ctx)

	// decode token
	token, err := jwt.Parse(parsedBearerToken, jwtParseFunction)
	if err != nil {
		ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": true,
			"msg":   err.Error(),
		})
	}

	// type assertion of token claims
	claims, ok := token.Claims.(jwt.MapClaims)
	if ok && token.Valid {
		userID, _ := claims["id"].(string)

		expires := int64(claims["expires"].(float64))

		userIDInt, _ := strconv.Atoi(userID)

		return &RequestTokenMetadata{
			UserID:  userIDInt,
			Expires: expires,
		}, nil
	}

	return nil, err
}

func jwtParseFunction(token *jwt.Token) (interface{}, error) {
	return []byte(os.Getenv("ACCESS_TOKEN_SECRET")), nil
}

func getParsedToken(ctx *fiber.Ctx) string {
	bearerToken := ctx.Get("Authorization")

	parsedBearerToken := strings.Split(bearerToken, " ")

	return parsedBearerToken[1]
}
