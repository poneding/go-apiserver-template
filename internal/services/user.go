package services

import (
	"errors"
	"go-apiserver-template/internal/db"
	"go-apiserver-template/internal/db/entities"
	"go-apiserver-template/internal/models"
)

type UserInterface interface {
	ListUsers() ([]models.User, error)
	GetUser(id uint64) (*models.User, error)
	CreateUser(user *models.User) (*models.User, error)
	UpdateUser(user *models.User) (*models.User, error)
	DeleteUser(id uint64) error
}

func NewUserService() UserInterface {
	return &userService{}
}

type userService struct{}

func (s *userService) ListUsers() ([]models.User, error) {
	var dbusers []entities.User
	if err := db.Repository().Find(&dbusers).Error; err != nil {
		return nil, err
	}

	var users []models.User

	for i := range dbusers {
		u := models.User{
			ID:   dbusers[i].ID,
			Name: dbusers[i].Name,
		}

		var orders []entities.Order
		if err := db.Repository().Where("user_id = ?", dbusers[i].ID).Find(&orders).Error; err != nil {
			return nil, err
		}

		for _, o := range orders {
			u.Orders = append(u.Orders, models.Order{
				ID:    o.ID,
				Price: o.Price,
			})
		}
		users = append(users, u)
	}

	return users, nil
}

func (s *userService) GetUser(id uint64) (*models.User, error) {
	var dbuser entities.User
	if err := db.Repository().Where("id = ?", id).First(&dbuser).Error; err != nil {
		return nil, err
	}

	user := models.User{
		ID:   dbuser.ID,
		Name: dbuser.Name,
	}

	var orders []entities.Order
	if err := db.Repository().Where("user_id = ?", dbuser.ID).Find(&orders).Error; err != nil {
		return nil, err
	}
	for _, o := range orders {
		user.Orders = append(user.Orders, models.Order{
			ID:    o.ID,
			Price: o.Price,
		})
	}

	return &user, nil
}

func (s *userService) CreateUser(user *models.User) (*models.User, error) {
	if user == nil {
		return nil, errors.New("user is nil")
	}
	dbuser := entities.User{
		Name: user.Name,
	}

	// TODO: use transaction
	if err := db.Repository().Save(&dbuser).Error; err != nil {
		return user, err
	}

	if len(user.Orders) > 0 {
		for i := range user.Orders {
			dborder := &entities.Order{
				UserID: dbuser.ID,
				Price:  user.Orders[i].Price,
			}
			if err := db.Repository().Save(dborder).Error; err != nil {
				return user, err
			}
			user.Orders[i].ID = dborder.ID
		}
	}
	user.ID = dbuser.ID

	return user, nil
}

func (s *userService) UpdateUser(user *models.User) (*models.User, error) {
	if user == nil {
		return user, errors.New("user is nil")
	}

	var dbuser entities.User
	if err := db.Repository().Where("id = ?", user.ID).First(&dbuser).Error; err != nil {
		return user, err
	}

	dbuser.Name = user.Name
	if err := db.Repository().Save(&dbuser).Error; err != nil {
		return user, err
	}

	return user, nil
}

func (s *userService) DeleteUser(id uint64) error {
	var user entities.User
	if err := db.Repository().Where("id = ?", id).Delete(&user).Error; err != nil {
		return err
	}

	return nil
}
