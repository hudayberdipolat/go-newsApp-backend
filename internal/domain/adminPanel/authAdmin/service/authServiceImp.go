package service

import (
	"errors"
	"github.com/hudayberdipolat/go-newsApp-backend/internal/domain/adminPanel/authAdmin/dto"
	"github.com/hudayberdipolat/go-newsApp-backend/internal/domain/adminPanel/authAdmin/repository"
	"github.com/hudayberdipolat/go-newsApp-backend/pkg/jwtToken/adminToken"
	"golang.org/x/crypto/bcrypt"
)

type authAdminServiceImp struct {
	authAdminRepo repository.AuthAdminRepository
}

func NewAuthAdminService(repo repository.AuthAdminRepository) AuthAdminService {
	return authAdminServiceImp{
		authAdminRepo: repo,
	}
}

func (a authAdminServiceImp) LoginAdmin(request dto.AdminLoginRequest) (*dto.AuthAdminResponse, error) {
	getAdmin, err := a.authAdminRepo.GetAdmin(request.PhoneNumber)
	if err != nil {
		return nil, err
	}

	errPassword := bcrypt.CompareHashAndPassword([]byte(getAdmin.Password), []byte(request.Password))
	if errPassword != nil {
		return nil, errors.New("Phone number ya-da password nadogry")
	}
	// generate access Token
	accessToken, errToken := adminToken.GenerateAdminToken(getAdmin.ID, getAdmin.PhoneNumber, getAdmin.AdminRole)
	if errToken != nil {
		return nil, errToken
	}
	adminLoginResponse := dto.NewAuthAdminResponse(*getAdmin, accessToken)
	return &adminLoginResponse, nil
}
