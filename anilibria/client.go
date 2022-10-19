package anilibria

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"pervaki/anilibria/model"
	"pervaki/lib/rateclient"
	serviceModel "pervaki/model"
	"time"

	"golang.org/x/time/rate"

	"go.uber.org/zap"
)

const (
	rateEvery = 1 * time.Second
	rateBurst = 3
)

const (
	host              = `http://api.anilibria.tv/v2`
	urlGetTitleFormat = `%s/getTitle?%s`
)

type Client struct {
	logger *zap.SugaredLogger
	cli    rateclient.RLHTTPClient
}

func NewClient(logger *zap.SugaredLogger, cli *http.Client) Client {
	return Client{
		logger: logger,
		cli:    rateclient.NewClient(cli, rate.NewLimiter(rate.Every(rateEvery), rateBurst)),
	}
}

func (c Client) GetTitle(ctx context.Context, filter model.GetTitleFilter) (serviceModel.Title, error) {
	var urlValues = make(url.Values)
	if len(filter.Code) != 0 {
		urlValues.Set("code", filter.Code)
	}

	var data model.Title
	err := c.do(ctx, http.MethodGet, fmt.Sprintf(urlGetTitleFormat, host, urlValues.Encode()), nil, &data)
	if err != nil {
		return serviceModel.Title{}, err
	}

	return MapClientToService(data), nil
}

func (c Client) do(ctx context.Context, method string, url string, body io.Reader, output interface{}) error {
	req, err := http.NewRequestWithContext(ctx, method, url, body)
	if err != nil {
		return err
	}

	res, err := c.cli.Do(req)
	if err != nil {
		return err
	}

	defer func() {
		_ = res.Body.Close()
	}()

	bodyBytes, err := io.ReadAll(res.Body)
	if err != nil {
		return err
	}

	if res.StatusCode != http.StatusOK {
		return fmt.Errorf("status code: %d; body: %v", res.StatusCode, string(bodyBytes))
	}

	if output != nil {
		err = json.Unmarshal(bodyBytes, &output)
		if err != nil {
			return err
		}
	}

	return nil
}
