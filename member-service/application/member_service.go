package application

import (
	"member-service/domain"
	"member-service/infrastructure"
)

type MemberServiceImpl struct {
	repo domain.MemberRepository
	kafka *infrastructure.KafkaClient
}

func NewMemberService(r domain.MemberRepository, k *infrastructure.KafkaClient) domain.MemberService {
	return &MemberServiceImpl{repo: r, kafka: k}
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
	err=infrastructure.SendMessage(s.kafka, "create-member", member)
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

func (s *MemberServiceImpl) GetAllMembers() ([]*domain.Member, error) {
	return s.repo.FindAll()
}
