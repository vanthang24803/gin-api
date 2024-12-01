package auth

import (
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/vanthang24803/api-ecommerce/internal/database"
	"github.com/vanthang24803/api-ecommerce/internal/models"
	"github.com/vanthang24803/api-ecommerce/internal/util"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type Repository interface {
	RegisterHandler(json *RegisterRequest) (interface{}, error)
	LoginHandler(json *LoginRequest) (*TokenResponse, error)
}

type repos struct {
	DB *gorm.DB
}

func (r *repos) RegisterHandler(json *RegisterRequest) (interface{}, error) {
	var existingUser models.User
	if err := r.DB.Where("email = ?", json.Email).First(&existingUser).Error; err != nil && err != gorm.ErrRecordNotFound {
		return nil, util.BadRequestException(err.Error())
	}

	var customerRole models.Role
	if err := r.DB.Where("name = ?", "Customer").First(&customerRole).Error; err != nil {
		return nil, util.NotFoundException("Customer role not found!")
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(json.Password), bcrypt.DefaultCost)
	if err != nil {
		return nil, util.BadRequestException("Hash password failed!")
	}

	newAccount := models.User{
		FirstName: json.FirstName,
		LastName:  json.LastName,
		Password:  string(hashedPassword),
		Email:     json.Email,
		Avatar:    "",
	}

	if err := r.DB.Create(&newAccount).Error; err != nil {
		return nil, util.BadRequestException("Failed to create user account: " + err.Error())
	}

	newAccountRole := models.UserRole{
		UserID: newAccount.ID,
		RoleID: customerRole.ID,
	}
	if err := r.DB.Create(&newAccountRole).Error; err != nil {
		return nil, util.BadRequestException("Failed to assign role to user: " + err.Error())
	}

	return newAccount, nil
}

func (r *repos) LoginHandler(json *LoginRequest) (*TokenResponse, error) {
	var user models.User
	var tokenAccount models.Token

	if err := r.DB.Where("email = ?", json.Email).First(&user).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, util.UnauthorizedException()
		}
		return nil, util.BadRequestException(err.Error())
	}

	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(json.Password))
	if err != nil {
		return nil, util.UnauthorizedException()
	}

	accessToken, refreshToken, err := generateJWTToken(&Payload{
		Email:    user.Email,
		FullName: user.FirstName + " " + user.LastName,
	})
	if err != nil {
		return nil, util.BadRequestException("Error generating token")
	}

	if err := r.DB.Where("name = ? AND user_id = ?", "RefreshToken", user.ID).First(&tokenAccount).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			newToken := models.Token{
				Name:      "RefreshToken",
				Value:     refreshToken,
				UserID:    user.ID,
				CreatedAt: time.Now(),
				UpdatedAt: time.Now(),
			}
			if err := r.DB.Create(&newToken).Error; err != nil {
				return nil, util.BadRequestException("Error creating token")
			}
		} else {
			return nil, util.BadRequestException(err.Error())
		}
	} else {
		tokenAccount.Value = refreshToken
		tokenAccount.UpdatedAt = time.Now()
		if err := r.DB.Save(&tokenAccount).Error; err != nil {
			return nil, util.BadRequestException("Error updating token")
		}
	}

	return &TokenResponse{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	}, nil
}

func generateJWTToken(payload *Payload) (string, string, error) {
	accessTokenClaims := jwt.MapClaims{
		"sub":     payload.Id,
		"iat":     time.Now().Unix(),
		"exp":     time.Now().Add(24 * time.Hour).Unix(),
		"payload": &payload,
	}

	accessToken := jwt.NewWithClaims(jwt.SigningMethodHS256, accessTokenClaims)

	accessTokenString, err := accessToken.SignedString([]byte(os.Getenv("JWT_SECRET")))
	if err != nil {
		return "", "", util.BadRequestException(err.Error())
	}

	refreshTokenClaims := jwt.MapClaims{
		"sub": payload.Id,
		"iat": time.Now().Unix(),
		"exp": time.Now().Add(7 * 24 * time.Hour).Unix(),
	}

	refreshToken := jwt.NewWithClaims(jwt.SigningMethodHS256, refreshTokenClaims)

	refreshTokenString, err := refreshToken.SignedString([]byte(os.Getenv("JWT_REFRESH")))
	if err != nil {
		return "", "", util.BadRequestException(err.Error())
	}

	return accessTokenString, refreshTokenString, nil
}

func AuthRepository() Repository {
	return &repos{
		DB: database.GetDb(),
	}
}
