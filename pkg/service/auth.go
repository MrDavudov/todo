package servise

import (
	"crypto/sha1"
	"fmt"

	"github.com/MrDavudov/todo/internal/model"
	"github.com/MrDavudov/todo/pkg/repository"
)

const salt = "fgdsfg3e32r23fd"

type AuthService struct {
	repo repository.Auth
}

func NewAuthService(repo repository.Auth) *AuthService {
	return &AuthService{
		repo: repo,
	}
}

func (s *AuthService) CreateUser(user model.User) (int, error) {
	user.Password = generatePasswordHash(user.Password)
	return s.repo.CreateUser(user)
}

func generatePasswordHash(password string) string {
	hash := sha1.New()
	hash.Write([]byte(password))

	return fmt.Sprintf("%x", hash.Sum([]byte(salt)))
}