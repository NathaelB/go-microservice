package application

import "member-service/domain"

type MemberServiceImpl struct {
	repo domain.MemberRepository
}

func NewMemberService(r domain.MemberRepository) domain.MemberService {
	return &MemberServiceImpl{repo: r}
}

func (s *MemberServiceImpl) CreateMember(name string) (*domain.Member, error) {
	member, err := domain.NewMember(name)
	if err != nil {
		return nil, err
	}

	err = s.repo.Save(member)
	if err != nil {
		return nil, err
	}

	return member, nil
}

func (s *MemberServiceImpl) GetMemberByID(id string) (*domain.Member, error) {
	member, err := s.repo.FindByID(id)
	if err != nil {
		return nil, err
	}

	return member, nil
}
