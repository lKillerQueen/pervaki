package rateclient

import (
	"golang.org/x/time/rate"
	"net/http"
)

type RLHTTPClient struct {
	client      *http.Client
	RateLimiter *rate.Limiter
}

func NewClient(cli *http.Client, rl *rate.Limiter) RLHTTPClient {
	return RLHTTPClient{
		client:      cli,
		RateLimiter: rl,
	}
}

func (c *RLHTTPClient) Do(req *http.Request) (*http.Response, error) {
	err := c.RateLimiter.Wait(req.Context())
	if err != nil {
		return nil, err
	}

	resp, err := c.client.Do(req)
	if err != nil {
		return nil, err
	}

	return resp, nil
}
