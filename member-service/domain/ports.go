package domain

type MemberRepository interface {
	Save(guild *Member) error
	FindByID(id string) (*Member, error)
	FindAll() ([]*Member, error)
}

type MemberService interface {
	CreateMember(name string) (*Member, error)
	GetMemberByID(id string) (*Member, error)
	GetAllMembers() ([]*Member, error)
}
