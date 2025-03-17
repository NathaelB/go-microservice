package domain

type MemberRepository interface {
	Save(guild *Member) error
	FindByID(id string) (*Member, error)
	FindAll() ([]*Member, error)
}

type MemberService interface {
	CreateMember(name string, guildID string) (*Member, error)
	GetMemberByID(id string) (*Member, error)
	GetAllMembers() ([]*Member, error)

	// HandleGuildCreated traite l'événement de création d'une guilde
	// et crée le membre correspondant à l'owner_id
	HandleGuildCreated(event GuildCreatedEvent) error
}
