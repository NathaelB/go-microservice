package application

import "guild-service/domain"

type GuildServiceImpl struct {
	repo domain.GuildRepository
}

func NewGuildService(r domain.GuildRepository) domain.GuildService {
	return &GuildServiceImpl{repo: r}
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

	return guild, nil
}

func (s *GuildServiceImpl) GetGuildByID(id string) (*domain.Guild, error) {
	guild, err := s.repo.FindByID(id)
	if err != nil {
		return nil, err
	}

	return guild, nil
}
