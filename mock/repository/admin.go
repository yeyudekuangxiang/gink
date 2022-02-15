package repository

import (
	"fmt"
	"github.com/yeyudekuangxiang/gink/model/entity"
	"gorm.io/gorm"
)

func NewAdminMockRepository() AdminMockRepository {
	return AdminMockRepository{}
}

type AdminMockRepository struct {
	db *gorm.DB
}

func (a AdminMockRepository) GetAdminById(id int) (*entity.Admin, error) {
	return &entity.Admin{
		ID:     id,
		Name:   fmt.Sprintf("mock%d", id),
		Avatar: "http://avatar.png",
	}, nil
}
