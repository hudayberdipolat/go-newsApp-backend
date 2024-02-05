package service

import (
	"errors"
	dto2 "github.com/hudayberdipolat/go-newsApp-backend/internal/domain/roles/dto"
	"github.com/hudayberdipolat/go-newsApp-backend/internal/domain/roles/repository"
	"github.com/hudayberdipolat/go-newsApp-backend/internal/models"
)

type roleServiceImp struct {
	roleRepo repository.RoleRepository
}

func NewRoleService(repo repository.RoleRepository) RoleService {
	return roleServiceImp{
		roleRepo: repo,
	}
}

func (r roleServiceImp) FindAll() ([]dto2.RoleResponse, error) {
	roles, err := r.roleRepo.GetAll()
	if err != nil {
		return nil, err
	}
	var roleResponses []dto2.RoleResponse
	for _, role := range roles {
		roleResponse := dto2.NewRoleResponse(role)
		roleResponses = append(roleResponses, roleResponse)
	}

	return roleResponses, nil
}

func (r roleServiceImp) FindOne(roleID int) (*dto2.RoleResponse, error) {
	role, err := r.roleRepo.GetOne(roleID)
	if err != nil {
		return nil, err
	}
	roleResponse := dto2.NewRoleResponse(*role)
	return &roleResponse, nil
}

func (r roleServiceImp) Create(request dto2.CreateRoleRequest) error {
	role := models.Role{
		RoleName: request.RoleName,
	}
	if err := r.roleRepo.Create(role); err != nil {
		return err
	}
	return nil
}

func (r roleServiceImp) Update(roleID int, request dto2.UpdateRoleRequest) error {
	updateRole, err := r.roleRepo.GetOne(roleID)
	if err != nil {
		return errors.New("role not found")
	}
	updateRole.RoleName = request.RoleName
	if errUpdate := r.roleRepo.Update(updateRole.ID, *updateRole); errUpdate != nil {
		return errUpdate
	}
	return nil
}

func (r roleServiceImp) Delete(roleID int) error {
	deleteRole, err := r.roleRepo.GetOne(roleID)
	if err != nil {
		return errors.New("role not found")
	}
	if errDelete := r.roleRepo.Delete(deleteRole.ID); err != nil {
		return errDelete
	}
	return nil
}
