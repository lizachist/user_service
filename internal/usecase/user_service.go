package usecase

import (
	"golang.org/x/crypto/bcrypt"
	"user_service/internal/domain"
)

type UserService struct {
	userRepo domain.UserRepository
}

func NewUserService(userRepo domain.UserRepository) domain.UserService {
	return &UserService{userRepo: userRepo}
}

func (u *UserService) Create(user *domain.User) error {
	// Хэширование пароля
	hashedPassword, err := hashPassword(user.Password)
	if err != nil {
		return err
	}
	user.PasswordHash = hashedPassword

	// Создание пользователя
	return u.userRepo.Create(user)
}

func (u *UserService) GetByID(id int) (*domain.User, error) {
	return u.userRepo.GetByID(id)
}

func (u *UserService) Update(user *domain.User) error {
	// Проверяем, существует ли пользователь
	existingUser, err := u.userRepo.GetByID(user.ID)
	if err != nil {
		return err
	}

	// Обновляем только те поля, которые предоставлены
	if user.Username != "" {
		existingUser.Username = user.Username
	}
	if user.Email != "" {
		existingUser.Email = user.Email
	}
	if user.FirstName != "" {
		existingUser.FirstName = user.FirstName
	}
	if user.LastName != "" {
		existingUser.LastName = user.LastName
	}
	existingUser.IsActive = user.IsActive

	// Если предоставлен новый пароль, хэшируем его
	if user.Password != "" {
		hashedPassword, err := hashPassword(user.Password)
		if err != nil {
			return err
		}
		existingUser.PasswordHash = hashedPassword
	}

	// Обновляем пользователя
	return u.userRepo.Update(existingUser)
}

func hashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	return string(bytes), err
}
