package service

import (
	"fmt"
	"github.com/amulya.bl/demoCRUDApp/entities"
	_ "github.com/amulya.bl/demoCRUDApp/entities"
	"github.com/amulya.bl/demoCRUDApp/models"
	"github.com/amulya.bl/demoCRUDApp/repositories"
	"github.com/amulya.bl/demoCRUDApp/request"
	_ "github.com/amulya.bl/demoCRUDApp/response"
	_ "strings"
	"time"
	_ "time"
)

type UserService struct {
	Repo *repositories.UserRepository
}

func NewUserService(repo *repositories.UserRepository) *UserService {
	return &UserService{Repo: repo}
}
func validateUser(name string) error {
	if len(name) < 3 {
		return fmt.Errorf("Name should be at least 3 characters")
	}
	return nil
}
func validateUserId(userId int) error {
	if userId <= 0 {
		return fmt.Errorf("Invalid user, Enter valid userID")
	}
	return nil
}
func modeltoEntity(userModel models.User) (userEntity entities.User) {
	userEntity = entities.User{
		ID:    userModel.ID,
		Name:  userModel.Name,
		Email: userModel.Email,
	}
	return userEntity
}
func entityToModel(userEntity entities.User) (userModel models.User) {
	userModel = models.User{
		ID:        userEntity.ID,
		Name:      userEntity.Name,
		Email:     userEntity.Email,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		CreatedBy: userEntity.Name,
		UpdatedBy: userEntity.Name,
	}
	return userModel
}
func (s *UserService) CreateUser(user *request.CreateUser) (err error) {
	err = validateUser(user.Name)
	if err != nil {
		return err
	}
	userEntity := entities.User{
		Name:  user.Name,
		Email: user.Email,
	}
	err = s.Repo.CreateUser(userEntity)

	if err != nil {
		return fmt.Errorf("User could not be created")
	}
	return nil

}

func (s *UserService) GetUser(userId int) (userEntity entities.User, err error) {
	err = validateUserId(userId)
	if err != nil {
		return entities.User{}, err
	}
	userModel, err := s.Repo.GetUser(userId)
	userEntity = modeltoEntity(userModel)
	if err != nil {
		return entities.User{}, err
	}
	return userEntity, nil

}

func (s *UserService) UpdateUser(userId int, updateUser request.UpdateUser) (err error) {
	err = validateUserId(userId)
	if err != nil {
		return err
	}
	userEntity, err := s.GetUser(userId)
	fmt.Println(err)
	if len(*updateUser.Name) > 0 {
		userEntity.Name = *updateUser.Name
	}
	if len(*updateUser.Email) > 0 {
		userEntity.Email = *updateUser.Email
	}

	userModel := entityToModel(userEntity)

	err = s.Repo.UpdateUser(userModel)
	return err

}
func (s *UserService) DeleteUser(userId int) (err error) {
	err = validateUserId(userId)
	if err != nil {
		return err
	}
	userEntity, err := s.GetUser(userId)
	userModel := entityToModel(userEntity)
	err = s.Repo.DeleteUser(userModel)

	return err

}

func (s *UserService) GetAllUsers() (responseUsers []map[string]interface{}, err error) {
	users, err := s.Repo.GetAllUser()
	if err != nil {
		return nil, err
	}
	//var responseUsers []map[string]interface{}

	for _, user := range users {
		responseUser := map[string]interface{}{
			"id":    user.ID,
			"name":  user.Name,
			"email": user.Email,
		}
		responseUsers = append(responseUsers, responseUser)
	}

	return responseUsers, nil

}
