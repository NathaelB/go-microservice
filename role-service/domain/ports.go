package domain

type RoleRepository interface {
	Save(role *Role) error
	FindByID(id string) (*Role, error)
	FindAll() ([]*Role, error)
}

type RoleService interface {
	CreateRole(name string, guilId string) (*Role, error)
	GetRoleByID(id string) (*Role, error)
	GetAllRoles() ([]*Role, error)
}
