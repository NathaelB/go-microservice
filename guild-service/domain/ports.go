package domain

type GuildRepository interface {
	Save(guild *Guild) error
	FindByID(id string) (*Guild, error)
}

type GuildService interface {
	CreateGuild(name, ownerID string) (*Guild, error)
	GetGuildByID(id string) (*Guild, error)
}
