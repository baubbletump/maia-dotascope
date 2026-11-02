package dota

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/baubbletump/maia-dotascope/internal/scope/dota/models"
)

type Client struct {
	client *http.Client
}

func NewClient() Client {
	return Client{
		client: &http.Client{
			Timeout: 5 * time.Second,
		},
	}
}

func (c *Client) GetPlayerById(ctx context.Context, accountID int) (*models.PlayerRankResponse, error) {
	url := fmt.Sprintf("https://api.opendota.com/api/players/%d", accountID)

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
	if err != nil {
		return nil, err
	}

	req.Header.Add("Accept", "application/json")

	resp, err := c.client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("request to OpenDota failed: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("OpenDota returned status: %s", resp.Status)
	}

	var player models.PlayerRankResponse

	if err := json.NewDecoder(resp.Body).Decode(&player); err != nil {
		return nil, fmt.Errorf("decode OpenDota response: %w", err)
	}

	return &player, nil
}
