package repositories

import (
	"fmt"
	"github.com/amulya.bl/demoCRUDApp/entities"
	"github.com/amulya.bl/demoCRUDApp/models"
	"gorm.io/gorm"
	"time"
)

type UserRepository struct {
	DB *gorm.DB
}

func NewUserRepository(db *gorm.DB) *UserRepository {
	return &UserRepository{DB: db}
}
func entityToModel(userEntity entities.User) (userModel models.User) {
	userModel = models.User{
		Name:      userEntity.Name,
		Email:     userEntity.Email,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		CreatedBy: userEntity.Name,
		UpdatedBy: userEntity.Name,
	}
	return userModel
}
func modeltoEntity(userModel models.User) (userEntity entities.User) {
	userEntity = entities.User{
		ID:    userModel.ID,
		Name:  userModel.Name,
		Email: userModel.Email,
	}
	return userEntity
}
func (r *UserRepository) GetUser(userId int) (userModel models.User, err error) {
	result := r.DB.First(&userModel, userId)
	if result.Error == gorm.ErrRecordNotFound {
		return models.User{}, result.Error
	}
	return userModel, nil

}

func (r *UserRepository) CreateUser(user entities.User) error {

	userModel := entityToModel(user)
	result := r.DB.Create(&userModel)
	if result.Error != nil {
		return fmt.Errorf("User could not be created")
	}
	return nil

}

func (r *UserRepository) UpdateUser(userModel models.User) error {
	result := r.DB.Save(&userModel)
	if result.Error != nil {
		return fmt.Errorf("user cannotbe updated")

	}
	return nil
}

func (r *UserRepository) DeleteUser(userModel models.User) error {
	result := r.DB.Delete(&userModel)
	if result.Error != nil {
		return fmt.Errorf("user could not be deleted")
	}
	return nil

}

func (r *UserRepository) GetAllUser() (users []models.User, err error) {

	result := r.DB.Find(&users)
	if result.Error != nil {
		return nil, err
	}
	return users, nil
}
