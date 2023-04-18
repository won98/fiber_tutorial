package utils

import (
	"fmt"
	"os"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v4"
)

func CreateToken(id string, no int) (string, error) {

	claims := jwt.MapClaims{}
	claims["authorized"] = true
	claims["id"] = id
	claims["no"] = no
	claims["exp"] = time.Now().Add(time.Hour * 24 * 365).Unix()
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(os.Getenv("ACCESS_KEY")))
}
func CreateRefreshToken(id string) (string, error) {

	claims := jwt.MapClaims{}
	claims["authorized"] = true
	claims["id"] = id
	claims["exp"] = time.Now().Add(time.Hour * 24 * 365).Unix()
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(os.Getenv("ACCESS_KEY")))
}

func VarifiyToken(c *fiber.Ctx) (string, error) {
	tokenData := c.Get("xauth")
	token, err := jwt.Parse(tokenData, func(token *jwt.Token) (interface{}, error) {
		if _, compare := token.Method.(*jwt.SigningMethodHMAC); !compare {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(os.Getenv("ACCESS_KEY")), nil
	})
	if err != nil {
		return "", fmt.Errorf("failed to parse token: %v", err)
	}
	claims, compare := token.Claims.(jwt.MapClaims)
	if compare && token.Valid {
		id := fmt.Sprintf("%v", claims["id"])
		return id, nil
	}
	return "", fmt.Errorf("invalid token")
}

func VarifiyRefreshToken(c *fiber.Ctx) (string, error) {
	tokenData := c.Get("xauth")
	token, err := jwt.Parse(tokenData, func(token *jwt.Token) (interface{}, error) {
		if _, compare := token.Method.(*jwt.SigningMethodHMAC); !compare {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(os.Getenv("ACCESS_KEY")), nil
	})
	if err != nil {
		return "", fmt.Errorf("failed to parse token: %v", err)
	}
	claims, compare := token.Claims.(jwt.MapClaims)
	if compare && token.Valid {
		id := fmt.Sprintf("%v", claims["id"])
		return id, nil
	}
	return "", fmt.Errorf("invalid token")
}

// func VarifiyToken(tokenString string) (string, error) {
// 	// 토큰 파싱
// 	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
// 		if _, compare := token.Method.(*jwt.SigningMethodHMAC); !compare {
// 			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
// 		}
// 		return []byte(os.Getenv("ACCESS_KEY")), nil
// 	})
// 	if err != nil {
// 		return "", fmt.Errorf("failed to parse token: %v", err)
// 	}

// 	// 토큰 유효성 검사 및 클레임 추출
// 	claims, compare := token.Claims.(jwt.MapClaims)
// 	if compare && token.Valid {
// 		id := fmt.Sprintf("%v", claims["id"])
// 		return id, nil
// 	}

// 	return "", fmt.Errorf("invalid token")
// }
