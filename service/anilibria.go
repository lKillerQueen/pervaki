package service

import (
	"context"
	"pervaki/anilibria"
	"pervaki/anilibria/model"
	serviceModel "pervaki/model"

	"go.uber.org/zap"
)

type AnilibriaRepo interface {
	Upsert(ctx context.Context, title serviceModel.Title) error
	Select() (string, error)
	UpsertThroughBuilder(ctx context.Context, title serviceModel.Title) error
}

type AnilibriaService struct {
	logger        *zap.SugaredLogger
	cli           anilibria.Client
	anilibriaRepo AnilibriaRepo
}

func NewAnilibriaService(logger *zap.SugaredLogger, cli anilibria.Client, anilibriaRepo AnilibriaRepo) AnilibriaService {
	return AnilibriaService{
		logger:        logger,
		cli:           cli,
		anilibriaRepo: anilibriaRepo,
	}
}

func (s AnilibriaService) GetTitleName(ctx context.Context, code string) (string, error) {
	title, err := s.cli.GetTitle(ctx, model.GetTitleFilter{Code: code})
	if err != nil {
		return "", err
	}
	err = s.anilibriaRepo.UpsertThroughBuilder(ctx, title)
	if err != nil {
		return "", err
	}

	return title.NameRu, nil
}
func (s AnilibriaService) GetAll(ctx context.Context) (string, error) {

	title, err := s.anilibriaRepo.Select()
	if err != nil {
		return "", err
	}

	return title, nil
}
