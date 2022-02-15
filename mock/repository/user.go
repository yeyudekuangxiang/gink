package repository

import (
	"fmt"
	"github.com/yeyudekuangxiang/gink/model/entity"
	"gorm.io/gorm"
	"strconv"
)

func NewUserMockRepository() UserMockRepository {
	return UserMockRepository{}
}

type UserMockRepository struct {
	db *gorm.DB
}

func (u UserMockRepository) GetUserById(id int) (*entity.User, error) {
	return &entity.User{
		ID:       id,
		Guid:     strconv.Itoa(id),
		Nickname: fmt.Sprintf("mock%d", id),
		Avatar:   "http://avatar.png",
	}, nil
}

func (u UserMockRepository) GetUserByGuid(guid string) (*entity.User, error) {
	return &entity.User{
		ID:       1,
		Guid:     guid,
		Nickname: fmt.Sprintf("mock%s", guid),
		Avatar:   "http://avatar.png",
	}, nil
}
