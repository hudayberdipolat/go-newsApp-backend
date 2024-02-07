package service

import (
	"errors"
	dto "github.com/hudayberdipolat/go-newsApp-backend/internal/domain/users/dto"
	"github.com/hudayberdipolat/go-newsApp-backend/internal/domain/users/repository"
	"github.com/hudayberdipolat/go-newsApp-backend/internal/models"
	"github.com/hudayberdipolat/go-newsApp-backend/pkg/jwtToken/userToken"
	"golang.org/x/crypto/bcrypt"
	"time"
)

type userServiceImp struct {
	userRepo repository.UserRepository
}

func NewUserService(repo repository.UserRepository) UserService {
	return userServiceImp{
		userRepo: repo,
	}
}

func (u userServiceImp) RegisterUser(request dto.RegisterUserRequest) (*dto.AuthUserResponse, error) {
	hashPassword, _ := bcrypt.GenerateFromPassword([]byte(request.Password), bcrypt.DefaultCost)
	user := models.User{
		FullName:    request.FullName,
		PhoneNumber: request.PhoneNumber,
		UserStatus:  "active",
		Password:    string(hashPassword),
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}
	if err := u.userRepo.Create(user); err != nil {
		return nil, err
	}
	getUser, _ := u.userRepo.GetUserByPhoneNumber(request.PhoneNumber)
	// generate access Token
	accessToken, err := userToken.GenerateUserToken(getUser.ID, getUser.PhoneNumber, getUser.UserStatus)
	if err != nil {
		return nil, err
	}
	userResponse := dto.NewAuthUserResponse(getUser, accessToken)
	return &userResponse, nil
}

func (u userServiceImp) LoginUser(request dto.LoginUserRequest) (*dto.AuthUserResponse, error) {
	getUser, err := u.userRepo.GetUserByPhoneNumber(request.PhoneNumber)
	if err != nil {
		return nil, errors.New("Phone number ya-da password nädogry!!!")
	}

	errPassword := bcrypt.CompareHashAndPassword([]byte(getUser.Password), []byte(request.Password))
	if errPassword != nil {
		return nil, errors.New("Phone number ya-da password nädogry!!!")
	}

	// user status check
	if getUser.UserStatus != "active" {
		return nil, errors.New("Siziň profiliňiz admin tarapyndan bloklanan! Acmak ucin admin-e yüz tutup bilersiňiz!")
	}
	accessToken, err := userToken.GenerateUserToken(getUser.ID, getUser.PhoneNumber, getUser.UserStatus)
	if err != nil {
		return nil, err
	}
	userResponse := dto.NewAuthUserResponse(getUser, accessToken)
	return &userResponse, nil
}

func (u userServiceImp) GetUserData(userID int, phoneNumber string) (*dto.UserResponse, error) {
	getUser, err := u.userRepo.GetUserData(userID, phoneNumber)
	if err != nil {
		return nil, err
	}
	userResponse := dto.NewUserResponse(getUser)
	return &userResponse, nil
}

func (u userServiceImp) UpdateUserData(userID int, data dto.ChangeUserData) (*dto.AuthUserResponse, error) {
	updateUser, err := u.userRepo.GetUserByID(userID)
	if err != nil {
		return nil, errors.New("can't updated user Data")
	}
	updateUser.FullName = data.FullName
	updateUser.PhoneNumber = data.PhoneNumber
	updateUser.UpdatedAt = time.Now()

	if errUpdate := u.userRepo.Update(userID, *updateUser); errUpdate != nil {
		return nil, errUpdate
	}

	getUser, err := u.userRepo.GetUserByPhoneNumber(updateUser.PhoneNumber)
	if err != nil {
		return nil, err
	}
	accessToken, err := userToken.GenerateUserToken(getUser.ID, getUser.PhoneNumber, getUser.UserStatus)
	userResponse := dto.NewAuthUserResponse(getUser, accessToken)
	return &userResponse, nil
}

func (u userServiceImp) UpdateUserPassword(userID int, password dto.ChangeUserPassword) error {
	hashPassword, _ := bcrypt.GenerateFromPassword([]byte(password.Password), bcrypt.DefaultCost)
	strPassword := string(hashPassword)
	if err := u.userRepo.ChangeUserPassword(userID, strPassword); err != nil {
		return err
	}
	return nil
}

func (u userServiceImp) DeleteUser(userID int, phoneNumber string) error {
	getUser, err := u.userRepo.GetUserByID(userID)
	if err != nil {
		return errors.New("can't deleted account")
	}
	if errDelete := u.userRepo.Delete(getUser.ID, phoneNumber); errDelete != nil {
		return errors.New("can't deleted account")
	}
	return nil
}
