package service

import (
	"errors"
	dto2 "github.com/hudayberdipolat/go-newsApp-backend/internal/domain/admins/dto"
	"github.com/hudayberdipolat/go-newsApp-backend/internal/domain/admins/repository"
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

func (a adminServiceImp) FindAll() ([]dto2.AdminResponse, error) {
	admins, err := a.adminRepo.GetAll()
	if err != nil {
		return nil, err
	}

	var adminResponses []dto2.AdminResponse

	for _, admin := range admins {
		adminResponse := dto2.NewAdminResponse(admin)
		adminResponses = append(adminResponses, adminResponse)
	}

	return adminResponses, nil
}

func (a adminServiceImp) FindOne(adminID int) (*dto2.AdminResponse, error) {
	admin, err := a.adminRepo.GetOne(adminID)
	if err != nil {
		return nil, err
	}
	adminResponse := dto2.NewAdminResponse(*admin)
	return &adminResponse, nil
}

func (a adminServiceImp) Create(request dto2.CreateAdminRequest) error {
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

func (a adminServiceImp) Update(adminID int, request dto2.UpdateAdminRequest) error {

	updateAdmin, err := a.adminRepo.GetOne(adminID)
	if err != nil {
		return errors.New("admin not found")
	}
	updateAdmin.FullName = request.FullName
	updateAdmin.PhoneNumber = request.PhoneNumber
	updateAdmin.UpdatedAt = time.Now()
	if errUpdate := a.adminRepo.Update(updateAdmin.ID, *updateAdmin); errUpdate != nil {
		return errUpdate
	}
	return nil
}

func (a adminServiceImp) Delete(adminID int) error {
	deleteAdmin, err := a.adminRepo.GetOne(adminID)
	if err != nil {
		return errors.New("admin not found")
	}

	if errDelete := a.adminRepo.Delete(deleteAdmin.ID); errDelete != nil {
		return errDelete
	}
	return nil
}

func (a adminServiceImp) ChangePassword(adminID int, changePassword dto2.ChangeAdminPassword) error {

	getAdmin, err := a.adminRepo.GetOne(adminID)
	if err != nil {
		return errors.New("admin not found")
	}

	if err := bcrypt.CompareHashAndPassword([]byte(getAdmin.Password), []byte(changePassword.OldPassword)); err != nil {
		return errors.New("Köne password nädogry!!!")
	}
	if errPasswordChange := a.adminRepo.ChangePassword(getAdmin.ID, changePassword.Password); errPasswordChange != nil {
		return err
	}
	return nil
}
