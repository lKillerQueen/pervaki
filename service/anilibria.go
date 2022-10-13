package service

import (
	"context"
	"go.uber.org/zap"
	"pervaki/anilibria"
	"pervaki/anilibria/model"
	serviceModel "pervaki/model"
)

type AnilibriaRepo interface {
	Upsert(ctx context.Context, title serviceModel.Title) error
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
