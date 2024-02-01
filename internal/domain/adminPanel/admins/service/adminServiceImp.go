package service

import (
	"errors"
	"github.com/hudayberdipolat/go-newsApp-backend/internal/domain/adminPanel/admins/dto"
	"github.com/hudayberdipolat/go-newsApp-backend/internal/domain/adminPanel/admins/repository"
	"github.com/hudayberdipolat/go-newsApp-backend/internal/models"
	"golang.org/x/crypto/bcrypt"
	"time"
)

type adminServiceImp struct {
	adminRepo repository.AdminRepository
}

func NewAdminService(repo repository.AdminRepository) AdminService {
	return adminServiceImp{
		adminRepo: repo,
	}
}

func (a adminServiceImp) FindAll() ([]dto.AdminResponse, error) {
	admins, err := a.adminRepo.GetAll()
	if err != nil {
		return nil, err
	}

	var adminResponses []dto.AdminResponse

	for _, admin := range admins {
		adminResponse := dto.NewAdminResponse(admin)
		adminResponses = append(adminResponses, adminResponse)
	}

	return adminResponses, nil
}

func (a adminServiceImp) FindOne(adminID int) (*dto.AdminResponse, error) {
	admin, err := a.adminRepo.GetOne(adminID)
	if err != nil {
		return nil, err
	}
	adminResponse := dto.NewAdminResponse(*admin)
	return &adminResponse, nil
}

func (a adminServiceImp) Create(request dto.CreateAdminRequest) error {
	hashPassword, _ := bcrypt.GenerateFromPassword([]byte(request.Password), bcrypt.DefaultCost)
	createAdmin := models.Admin{
		FullName:    request.FullName,
		PhoneNumber: request.PhoneNumber,
		AdminRole:   "admin",
		Password:    string(hashPassword),
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}

	if err := a.adminRepo.Create(createAdmin); err != nil {
		return err
	}

	return nil
}

func (a adminServiceImp) Update(adminID int, request dto.UpdateAdminRequest) error {

	updateAdmin, err := a.adminRepo.GetOne(adminID)
	if err != nil {
		return errors.New("admin not found")
	}
	updateAdmin.FullName = request.FullName
	updateAdmin.PhoneNumber = request.PhoneNumber
	updateAdmin.UpdatedAt = time.Now()
	if errUpdate := a.adminRepo.Update(adminID, *updateAdmin); errUpdate != nil {
		return errUpdate
	}
	return nil
}

func (a adminServiceImp) Delete(adminID int) error {

	_, err := a.adminRepo.GetOne(adminID)
	if err != nil {
		return errors.New("admin not found")
	}

	if errDelete := a.adminRepo.Delete(adminID); errDelete != nil {
		return errDelete
	}
	return nil
}

func (a adminServiceImp) ChangePassword(adminID int, changePassword dto.ChangeAdminPassword) error {

	getAdmin, err := a.adminRepo.GetOne(adminID)
	if err != nil {
		return errors.New("admin not found")
	}

	if err := bcrypt.CompareHashAndPassword([]byte(getAdmin.Password), []byte(changePassword.OldPassword)); err != nil {
		return errors.New("Köne password nädogry!!!")
	}
	if errPasswordChange := a.adminRepo.ChangePassword(adminID, changePassword.Password); errPasswordChange != nil {
		return err
	}
	return nil
}
