package usecase

import (
	"encoding/json"
	"fmt"
	"get-pr-comment/request"
)

type ReviewCommentResponse struct {
	Url               string `json:"url"`
	ID                int    `json:"id"`
	PrUrl             string `json:"pull_request_url"`
	Comment           string `json:"body"`
	AuthorAssociation string `json:"author_association"`
	CreatedAt         string `json:"created_at"`
	UpdateAt          string `json:"updated_at"`
	User              User   `json:"user"`
}

func FetchAllReviwComment(
	repository string,
	owner string,
	token string,
) ([]ReviewCommentResponse, error) {
	url := fmt.Sprintf(
		"https://api.github.com/repos/%s/%s/pulls/comments?sort=updated&direction=desc&per_page=100",
		owner,
		repository,
	)

	res, err := request.RequestGitHubAPI(url, token)
	if err != nil {
		return nil, fmt.Errorf("failed request.RequestGitHubAPI: %w", err)
	}
	defer res.Body.Close()

	var commnets []ReviewCommentResponse
	err = json.NewDecoder(res.Body).Decode(&commnets)
	if err != nil {
		return nil, fmt.Errorf("comment json decode err %w", err)
	}

	return commnets, nil
}
