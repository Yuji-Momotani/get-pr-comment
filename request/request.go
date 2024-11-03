package request

import (
	"fmt"
	"net/http"
	"time"
)

func RequestGitHubAPI(url string, token string) (
	*http.Response,
	error,
) {
	client := http.Client{
		Timeout: 30 * time.Second,
	}

	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return nil, fmt.Errorf("failed http.NewRequest : %w", err)
	}

	req.Header.Set("Accept", "application/vnd.github.v3+json")
	req.Header.Set("Authorization", fmt.Sprintf("token %s", token))

	res, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("failed client.Do : %w", err)
	}

	if res.StatusCode != http.StatusOK {
		return nil, fmt.Errorf(
			"unexpected status got %v",
			res.StatusCode,
		)
	}

	return res, nil
}
