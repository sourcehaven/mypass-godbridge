package controllers

import (
	"errors"
	"github.com/sourcehaven/mypass-godbridge/pkg/models"
	"github.com/sourcehaven/mypass-godbridge/pkg/security/crypto"
	"gorm.io/gorm"
)

var ErrAuthPw = errors.New("authentication error :: paswords do not match")

type UserController struct {
	*gorm.DB
}

func (c *UserController) Create(model models.User) (err error) {
	if err != nil {
		return
	}
	tx := c.Begin()
	if model.Password, err = crypto.PasswordHash(model.Password); err != nil {
		return
	}
	if err = tx.Create(&model).Error; err != nil {
		tx.Rollback()
		return
	}
	if err = tx.Commit().Error; err != nil {
		return
	}
	return
}

func (c *UserController) ActivateByUsername(username string, oldPassword string, newPassword string) (err error) {
	tx := c.Begin()
	var user *models.User
	if err = tx.Where(&models.User{Username: username}).First(&user).Error; err != nil {
		return
	}
	valid, err := crypto.ValidatePassword(oldPassword, user.Password)
	if err != nil {
		return
	}
	if !valid {
		return ErrAuthPw
	}
	newHash, err := crypto.PasswordHash(newPassword)
	if err != nil {
		return
	}
	if err = tx.Model(&user).Updates(models.User{Active: true, Password: newHash}).Error; err != nil {
		tx.Rollback()
		return
	}
	if err = tx.Commit().Error; err != nil {
		return
	}
	return
}

func (c *UserController) DeleteByUsername(username string) (err error) {
	tx := c.Begin()
	var user *models.User
	if err = tx.Where(&models.User{Username: username}).First(&user).Error; err != nil {
		return
	}
	// Delete record permanently
	if err = tx.Unscoped().Delete(&user).Error; err != nil {
		tx.Rollback()
		return
	}
	if err = tx.Commit().Error; err != nil {
		return
	}
	return
}

func (c *UserController) Authenticate(username string, password string) (err error) {
	user := &models.User{}
	if err = c.Where(&models.User{Username: username}).First(user).Error; err != nil {
		return
	}
	valid, err := crypto.ValidatePassword(password, user.Password)
	if err != nil {
		return
	}
	if !valid {
		return ErrAuthPw
	}
	return
}
