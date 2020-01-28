package services

import (
	"github.com/darwin1224/go-echo-auth-crud/models"
	"github.com/jinzhu/gorm"
)

type (
	UserService struct {
		Repo *gorm.DB
	}
)

func (s *UserService) GetAllUser() (*[]models.User, error) {
	users := new([]models.User)
	if err := s.Repo.Find(&users).Error; err != nil {
		return nil, err
	}
	return users, nil
}

func (s *UserService) GetUserById(id int) (*models.User, error) {
	user := new(models.User)
	if err := s.Repo.Where("ID = ?", id).First(&user).Error; err != nil {
		return nil, err
	}
	return user, nil
}

func (s *UserService) InsertUser(user *models.User) error {
	if err := s.Repo.Save(&user).Error; err != nil {
		return err
	}
	return nil
}

func (s *UserService) UpdateUser(user *models.User) error {
	if err := s.Repo.Save(&user).Error; err != nil {
		return err
	}
	return nil
}

func (s *UserService) DeleteUser(user *models.User) error {
	if err := s.Repo.Delete(&user).Error; err != nil {
		return err
	}
	return nil
}

func (s *UserService) GetUserByUsername(userDto *models.User) (*models.User, error) {
	user := new(models.User)
	if err := s.Repo.Where("username = ?", userDto.Username).First(&user).Error; err != nil {
		return nil, err
	}
	return user, nil
}
