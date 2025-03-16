package domain

type RoleRepository interface {
	Save(guild *Role) error
	FindByID(id string) (*Role, error)
}

type RoleService interface {
	CreateRole(name string, guildID string) (*Role, error)
	GetRoleByID(id string) (*Role, error)
}