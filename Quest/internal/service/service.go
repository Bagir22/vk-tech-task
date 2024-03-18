package service

import (
	"Quest/internal/types"
	"Quest/internal/utils"
	"errors"
	"golang.org/x/net/context"
)

type Repository interface {
	AddUser(ctx context.Context, User types.User) error
	AddQuest(ctx context.Context, quest types.Quest) error
	ProcessSignal(ctx context.Context, signal types.Signal) (types.User, error)
	GetUserHistory(ctx context.Context, id int) ([]types.UserHistory, error)
	GetUsers(ctx context.Context) ([]types.UserFromDb, error)
	GetQuestById(ctx context.Context, id int) (types.QuestFromDb, error)
	UpdateQuest(ctx context.Context, quest types.Quest, id int) error
	GetQuests(ctx context.Context) ([]types.QuestFromDb, error)
}

type Service struct {
	repo Repository
}

func InitService(repo Repository) *Service {
	return &Service{
		repo: repo,
	}
}

func (s *Service) AddUser(ctx context.Context, user types.User) error {
	validate := utils.ValidateUser(user)
	if !validate {
		return errors.New("Can't validate user")
	}
	return s.repo.AddUser(ctx, user)
}

func (s *Service) AddQuest(ctx context.Context, quest types.Quest) error {
	validate := utils.ValidateQuest(quest)
	if !validate {
		return errors.New("Can't validate quest")
	}
	return s.repo.AddQuest(ctx, quest)
}

func (s *Service) ProcessSignal(ctx context.Context, signal types.Signal) (types.User, error) {
	validate := utils.ValidateSignal(signal)
	if !validate {
		return types.User{}, errors.New("Can't validate signal")
	}
	return s.repo.ProcessSignal(ctx, signal)
}

func (s *Service) GetUserHistory(ctx context.Context, id int) ([]types.UserHistory, error) {
	return s.repo.GetUserHistory(ctx, id)
}

func (s *Service) GetUsers(ctx context.Context) ([]types.UserFromDb, error) {
	return s.repo.GetUsers(ctx)
}

func (s *Service) GetQuestById(ctx context.Context, id int) (types.QuestFromDb, error) {
	return s.repo.GetQuestById(ctx, id)
}

func (s *Service) UpdateQuest(ctx context.Context, quest types.Quest, id int) error {
	validate := utils.ValidateQuest(quest)
	if !validate {
		return errors.New("Can't validate quest")
	}
	return s.repo.UpdateQuest(ctx, quest, id)
}

func (s *Service) GetQuests(ctx context.Context) ([]types.QuestFromDb, error) {
	return s.repo.GetQuests(ctx)
}
