package application

import "role-service/domain"

type RoleServiceImpl struct {
	repo domain.RoleRepository
}

func NewRoleService(r domain.RoleRepository) domain.RoleService {
	return &RoleServiceImpl{repo: r}
}

func (s *RoleServiceImpl) CreateRole(name string, guildID string) (*domain.Role, error) {
	role, err := domain.NewRole(name, guildID)
	if err != nil {
		return nil, err
	}

	err = s.repo.Save(role)
	if err != nil {
		return nil, err
	}

	return role, nil
}

func (s *RoleServiceImpl) GetRoleByID(id string) (*domain.Role, error) {
	role, err := s.repo.FindByID(id)
	if err != nil {
		return nil, err
	}

	return role, nil
}
