package service

import (
	"context"
	"encoding/json"
	"fmt"
	"go.uber.org/zap"
	"io"
	"net/http"
	"pervaki/model"
)

const (
	anilibriaHost     = `https://api.anilibria.tv/v2/`
	anilibriaGetTitle = `getTitle`
)

type AnilibriaService struct {
	logger *zap.SugaredLogger
	cli    *http.Client
}

func NewAnilibriaService(logger *zap.SugaredLogger, cli *http.Client) AnilibriaService {
	return AnilibriaService{
		logger: logger,
		cli:    cli,
	}
}

func (s AnilibriaService) GetTitle(ctx context.Context, filter model.GetTitleFilter) (string, error) {
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, anilibriaHost+anilibriaGetTitle, nil)
	if err != nil {
		return "", err
	}

	var query = req.URL.Query()
	if len(filter.Code) != 0 {
		query.Set("code", filter.Code)
	}
	req.URL.RawQuery = query.Encode()

	res, err := s.cli.Do(req)
	if err != nil {
		return "", err
	}

	defer func() {
		_ = res.Body.Close()
	}()

	bodyBytes, err := io.ReadAll(res.Body)
	if err != nil {
		return "", err
	}

	if res.StatusCode != http.StatusOK {
		return "", fmt.Errorf("status code: %d; body: %s", res.StatusCode, string(bodyBytes))
	}

	var data model.Title
	err = json.Unmarshal(bodyBytes, &data)
	if err != nil {
		return "", err
	}

	return data.Names.Ru, nil
}
