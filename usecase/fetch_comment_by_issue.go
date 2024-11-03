package usecase

import (
	"encoding/json"
	"fmt"
	"get-pr-comment/request"
)

type IssueCommentResponse struct {
	Url               string `json:"url"`
	ID                int    `json:"id"`
	Comment           string `json:"body"`
	AuthorAssociation string `json:"author_association"`
	CreatedAt         string `json:"created_at"`
	UpdateAt          string `json:"updated_at"`
	User              User   `json:"user"`
}

func FetchCommentFromIssues(
	repository string,
	owner string,
	token string,
	prNumber int,
) ([]IssueCommentResponse, error) {
	url := fmt.Sprintf(
		// クローズ済みPRのみ、更新日の降順に取得。
		"https://api.github.com/repos/%s/%s/issues/%d/comments?state=close&sort=updated&direction=desc&per_page=50",
		owner,
		repository,
		prNumber,
	)

	res, err := request.RequestGitHubAPI(url, token)
	if err != nil {
		return nil, fmt.Errorf("failed request.RequestGitHubAPI: %w", err)
	}
	defer res.Body.Close()

	var commnets []IssueCommentResponse
	err = json.NewDecoder(res.Body).Decode(&commnets)
	if err != nil {
		return nil, fmt.Errorf("comment json decode err %w", err)
	}

	return commnets, nil
}
