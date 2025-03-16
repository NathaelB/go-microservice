package application

import (
	"guild-service/domain"
	"guild-service/infrastructure"
)

type GuildServiceImpl struct {
	repo  domain.GuildRepository
	kafka *infrastructure.KafkaClient
}

func NewGuildService(r domain.GuildRepository, k *infrastructure.KafkaClient) domain.GuildService {
	return &GuildServiceImpl{repo: r, kafka: k}
}

func (s *GuildServiceImpl) CreateGuild(name, ownerID string) (*domain.Guild, error) {
	guild, err := domain.NewGuild(name, ownerID)
	if err != nil {
		return nil, err
	}

	err = s.repo.Save(guild)
	if err != nil {
		return nil, err
	}

	err = infrastructure.SendMessage(s.kafka, "create-guild", guild)

	if err != nil {
		return nil, err
	}

	return guild, nil
}

func (s *GuildServiceImpl) GetGuildByID(id string) (*domain.Guild, error) {
	guild, err := s.repo.FindByID(id)
	if err != nil {
		return nil, err
	}

	return guild, nil
}

func (s *GuildServiceImpl) GetAllGuilds() ([]*domain.Guild, error) {
	return s.repo.FindAll()
}
