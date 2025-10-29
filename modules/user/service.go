package user

import (
	"gorm.io/gorm"
)

type Service struct {
	DB *gorm.DB
}

func NewService(db *gorm.DB) *Service {
	return &Service{DB: db}
}

func (s *Service) CreateUser(u *User) error {
	return s.DB.Create(u).Error
}

func (s *Service) GetUsers() ([]User, error) {
	var users []User
	err := s.DB.Find(&users).Error
	return users, err
}

func (s *Service) GetUser(id string) (*User, error) {
	var user User
	err := s.DB.First(&user, id).Error
	return &user, err
}

func (s *Service) UpdateUser(id string, input *User) (*User, error) {
	var user User
	if err := s.DB.First(&user, id).Error; err != nil {
		return nil, err
	}
	user.Name = input.Name
	user.Email = input.Email
	if err := s.DB.Save(&user).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

func (s *Service) DeleteUser(id string) error {
	return s.DB.Delete(&User{}, id).Error
}
