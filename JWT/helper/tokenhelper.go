package helper

import (
	"errors"
	"os"
	"time"

	"github.com/Goutham/Gin/models"
	"github.com/golang-jwt/jwt/v5"
	"github.com/joho/godotenv"
	"gorm.io/gorm"
)

type SignedDetails struct {
	Email     string
	Name      string
	UserID    string
	User_Type string
	jwt.RegisteredClaims
}

var _ = godotenv.Load(".env")
var secret_key = os.Getenv("SECRET_KEY")

func VerifyToken(userToken string) (*SignedDetails, error) {
	token, err := jwt.ParseWithClaims(userToken, &SignedDetails{}, func(t *jwt.Token) (interface{}, error) {
		return []byte(secret_key), nil
	})
	if err != nil {
		return nil, err
	}
	claims, ok := token.Claims.(*SignedDetails)

	if !ok {
		return nil, errors.New("error validaitng token provided")
	}
	if claims.ExpiresAt.Time.Before(time.Now()) {
		return nil, errors.New("token expired")
	}
	return claims, nil
}

func UpdateToken(db *gorm.DB, token, refresh_token, PhoneNumber string) error {

	createdAt, _ := time.Parse(time.RFC1123, time.Now().Format(time.RFC1123))

	if err := db.Model(&models.User{}).Where("Phone=?", PhoneNumber).Updates(map[string]any{
		"token":         token,
		"refresh_token": refresh_token,
		"createdat":     createdAt,
	}).Error; err != nil {
		return err
	}
	return nil
}

func GenerateAllTokens(Email, Name, UserID, User_Type string) (string, string, error) {

	claims := &SignedDetails{
		Email:     Email,
		Name:      Name,
		UserID:    UserID,
		User_Type: User_Type,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(24 * time.Hour)),
		},
	}

	refreshClaims := &SignedDetails{
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(24 * time.Hour)),
		},
	}
	token, err := jwt.NewWithClaims(jwt.SigningMethodHS256, claims).SignedString([]byte(secret_key))
	if err != nil {

		return "", "", err
	}
	refresh_token, err := jwt.NewWithClaims(jwt.SigningMethodHS256, refreshClaims).SignedString([]byte(secret_key))
	if err != nil {

		return "", "", err
	}

	return token, refresh_token, nil
}
