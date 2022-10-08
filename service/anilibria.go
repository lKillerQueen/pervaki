package service

import (
	"context"
	"go.uber.org/zap"
	"pervaki/anilibria"
	"pervaki/anilibria/models"
)

type AnilibriaService struct {
	logger *zap.SugaredLogger
	cli    anilibria.Client
}

func NewAnilibriaService(logger *zap.SugaredLogger, cli anilibria.Client) AnilibriaService {
	return AnilibriaService{
		logger: logger,
		cli:    cli,
	}
}

func (s AnilibriaService) GetTitleName(ctx context.Context, code string) (string, error) {
	title, err := s.cli.GetTitle(ctx, models.GetTitleFilter{Code: code})
	if err != nil {
		return "", err
	}

	return title.Names.Ru, nil
}
