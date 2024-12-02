package me

import (
	"github.com/vanthang24803/api-ecommerce/internal/database"
	"github.com/vanthang24803/api-ecommerce/internal/dto"
	"github.com/vanthang24803/api-ecommerce/internal/models"
	"github.com/vanthang24803/api-ecommerce/internal/util"
	"gorm.io/gorm"
)

type Repository interface {
	UpdateProfileHandler(user *dto.Payload, json *dto.UpdateProfile) (interface{}, error)
	GetProfileHandler(payload *dto.Payload) (interface{}, error)
	UploadAvatarHandler(payload *dto.Payload, fileName string) (interface{}, error)
}

func (r *repos) GetProfileHandler(payload *dto.Payload) (interface{}, error) {
	var existingUser models.User

	if err := r.DB.Preload("Roles").Where("id = ?", payload.Id).First(&existingUser).Error; err != nil {
		return nil, util.BadRequestException("User not found!")
	}

	return existingUser, nil
}

func (r *repos) UpdateProfileHandler(payload *dto.Payload, json *dto.UpdateProfile) (interface{}, error) {
	var existingUser models.User

	if err := r.DB.Where("id = ?", payload.Id).First(&existingUser).Error; err != nil {
		return nil, util.BadRequestException("User not found!")
	}

	if err := r.DB.Model(existingUser).Updates(models.User{
		FirstName: json.FirstName,
		LastName:  json.LastName,
		Email:     json.Email,
	}).Error; err != nil {
		return nil, util.BadRequestException("Failed to update user profile")
	}

	return existingUser, nil
}

func (r *repos) UploadAvatarHandler(payload *dto.Payload, fileName string) (interface{}, error) {
	var existingUser models.User

	if err := r.DB.Where("id = ?", payload.Id).First(&existingUser).Error; err != nil {
		return nil, util.BadRequestException("User not found!")
	}

	if err := r.DB.Model(existingUser).Updates(models.User{
		Avatar: fileName,
	}).Error; err != nil {
		return nil, util.BadRequestException("Failed to update user profile")
	}

	return "Upload avatar successfully!", nil
}

type repos struct {
	DB *gorm.DB
}

func MeRepository() Repository {
	return &repos{
		DB: database.GetDb(),
	}
}
