package service

import (
	"seclab/model"
	"time"

	"github.com/golang-jwt/jwt/v4"
)

type IRepository interface {
	FindProductByCategory(category string) ([]model.Product, error)
	Login(username, password string) (model.User, error)
	FindUserByID(id int) (model.User, error)
}

type Service struct {
	repo IRepository
}

func NewService(repo IRepository) *Service {
	return &Service{
		repo: repo,
	}
}

func (s *Service) Login(username, password string) (string, error) {
	user, err := s.repo.Login(username, password)
	if err != nil {
		return "", err
	}
	// Set custom claims
	claims := &model.JwtCustomClaims{
		ID:   user.ID,
		Name: user.Username,
		Role: user.Role,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour * 72)),
		},
	}

	// Create token with claims
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// Generate encoded token and send it as response.
	t, err := token.SignedString([]byte("secret"))
	if err != nil {
		return "", err
	}

	return t, nil
}

func (s *Service) FindUserByID(id int) (model.User, error) {
	return s.repo.FindUserByID(id)
}

func (s *Service) FindProductByCategory(category string) (*[]model.Product, error) {
	products, err := s.repo.FindProductByCategory(category)

	if err != nil {
		return nil, err
	}

	return &products, nil
}
