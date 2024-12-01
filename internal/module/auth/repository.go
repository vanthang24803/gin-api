package auth

import (
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/vanthang24803/api-ecommerce/internal/database"
	"github.com/vanthang24803/api-ecommerce/internal/models"
	"github.com/vanthang24803/api-ecommerce/internal/util"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type Repository interface {
	RegisterHandler(json *models.Register) (interface{}, error)
	LoginHandler(json *models.Login) (interface{}, error)
}

type repos struct {
	DB *gorm.DB
}

func (r *repos) RegisterHandler(json *models.Register) (interface{}, error) {
	var existingUser models.User

	if err := r.DB.Where("email = ?", json.Email).First(&existingUser).Error; err != nil {
		if err != gorm.ErrRecordNotFound {
			return nil, util.BadRequestException(err.Error())
		}
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(json.Password), bcrypt.DefaultCost)
	if err != nil {
		return nil, util.BadRequestException("Hash password wrong!")
	}

	newAccount := models.User{
		FirstName: json.FirstName,
		LastName:  json.LastName,
		Password:  string(hashedPassword),
		Email:     json.Email,
		Avatar:    "",
	}

	if err := r.DB.Create(&newAccount).Error; err != nil {
		return nil, util.BadRequestException(err.Error())
	}

	return newAccount, nil
}

func (r *repos) LoginHandler(json *models.Login) (interface{}, error) {
	var user models.User

	if err := r.DB.Where("email = ?", json.Email).First(&user).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, util.BadRequestException("Email not found")
		}
		return nil, util.BadRequestException(err.Error())
	}

	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(json.Password))
	if err != nil {
		return nil, util.BadRequestException("Invalid password")
	}

	token, err := generateJWTToken(user)
	if err != nil {
		return nil, util.BadRequestException("Error generating token")
	}

	return map[string]interface{}{
		"token": token,
	}, nil

}

func generateJWTToken(user models.User) (string, error) {
	secretKey := []byte(util.GetEnv("JWT_SECRET"))

	claims := jwt.MapClaims{
		"sub":   user.ID.String(),
		"email": user.Email,
		"exp":   time.Now().Add(time.Hour * 24).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(secretKey)
}

func AuthRepository() Repository {
	return &repos{
		DB: database.GetDb(),
	}
}
