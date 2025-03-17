package application

import (
	"role-service/domain"
	"role-service/infrastructure"
)

type RoleServiceImpl struct {
	repo domain.RoleRepository
	kafka *infrastructure.KafkaClient
}

func NewRoleService(r domain.RoleRepository, k *infrastructure.KafkaClient) domain.RoleService {
	return &RoleServiceImpl{repo: r, kafka: k}
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

func (s *RoleServiceImpl) GetAllRoles() ([]*domain.Role, error) {
	return s.repo.FindAll()
}