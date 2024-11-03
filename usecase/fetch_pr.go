package usecase

import (
	"encoding/json"
	"fmt"
	"get-pr-comment/request"
)

type PullResponse struct {
	Url               string `json:"url"`
	ID                int    `json:"id"`
	CommitsUrl        string `json:"commits_url"`
	ReviewCommentsUrl string `json:"review_comments_url"`
	Number            int    `json:"number"` // PR番号
	State             string `json:"state"`
	Title             string `json:"title"`
	CreatedAt         string `json:"created_at"`
	UpdateAt          string `json:"updated_at"`
	ClosedAt          string `json:"closed_at"`
	MergedAt          string `json:"merged_at"`
	User              User   `json:"user"`
}

type User struct {
	Login     string `json:"login"` // これがユーザー名っぽい
	ID        int    `json:"id"`
	NodeID    string `json:"node_id"`
	Url       string `json:"url"`
	Type      string `json:"type"`
	SiteAdmin bool   `json:"site_admin"`
}

func FetchPR(repository string, owner string, token string) ([]PullResponse, error) {
	url := fmt.Sprintf(
		"https://api.github.com/repos/%s/%s/pulls?state=close&sort=updated&direction=desc&per_page=50",
		owner,
		repository,
	)

	res, err := request.RequestGitHubAPI(url, token)
	if err != nil {
		return nil, fmt.Errorf("failed request.RequestGitHubAPI: %w", err)
	}
	defer res.Body.Close()

	var pulls []PullResponse
	err = json.NewDecoder(res.Body).Decode(&pulls)
	if err != nil {
		return nil, fmt.Errorf("pulls json decode err %w", err)
	}

	return pulls, nil
}
