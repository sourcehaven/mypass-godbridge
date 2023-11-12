package controllers

import (
	"github.com/sourcehaven/mypass-godbridge/pkg/models"
	"gorm.io/gorm"
)

type VaultController struct {
	*gorm.DB
}

func (c *VaultController) Create(model models.Vault) (err error) {
	if err != nil {
		return
	}
	tx := c.Begin()
	if err = tx.Create(&model).Error; err != nil {
		tx.Rollback()
		return
	}
	if err = tx.Commit().Error; err != nil {
		return
	}
	return
}
