package application

import (
	"fmt"
	"member-service/domain"
	"member-service/infrastructure"
)

type MemberServiceImpl struct {
	repo  domain.MemberRepository
	kafka *infrastructure.KafkaClient
}

func NewMemberService(r domain.MemberRepository, k *infrastructure.KafkaClient) domain.MemberService {
	return &MemberServiceImpl{repo: r, kafka: k}
}

func (s *MemberServiceImpl) CreateMember(name string, guildID string) (*domain.Member, error) {
	member, err := domain.NewMember(name, guildID)
	if err != nil {
		return nil, err
	}

	err = s.repo.Save(member)
	if err != nil {
		return nil, err
	}
	err = infrastructure.SendMessage(s.kafka, "create-member", member)
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

// HandleGuildCreated traite l'événement de création d'une guilde
// et crée le membre correspondant à l'owner_id
func (s *MemberServiceImpl) HandleGuildCreated(event domain.GuildCreatedEvent) error {
	// Vérifie si le membre existe déjà
	existingMember, err := s.repo.FindByID(event.OwnerID)
	if err == nil && existingMember != nil {
		// Le membre existe déjà, on ne fait rien
		return nil
	}

	// Crée un nouveau membre avec l'ID de l'owner
	member := &domain.Member{
		ID:        event.OwnerID,
		Name:      fmt.Sprintf("Owner of %s", event.Name), // Nom par défaut basé sur la guilde
		GuildID:   event.ID,
		CreatedAt: event.CreatedAt,
		UpdatedAt: event.UpdatedAt,
	}

	// Sauvegarde le membre dans la base de données
	if err := s.repo.Save(member); err != nil {
		return fmt.Errorf("failed to save member: %w", err)
	}

	return nil
}
