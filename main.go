package main

import (
	"fmt"
	"get-pr-comment/usecase"
)

const (
	token      = "" // access_token
	repository = "" // https://github.com/{owner}/{repository}/...
	owner      = "" // https://github.com/{owner}/{repository}/...
	user       = "" // user名
)

type Comment struct {
	Message string
	User    string //コメントを記載してくれた方
}

func main() {
	// PRを取得
	prs, err := usecase.FetchPR(repository, owner, token)
	if err != nil {
		fmt.Println(err)

		return
	}

	// 特定の行に指摘をもらった場合のコメント（以降、レビューコメントと呼ぶ）
	allPrComments, err := usecase.FetchAllReviwComment(repository, owner, token)
	if err != nil {
		fmt.Println(err)

		return
	}

	commentsMap := make(map[string][]Comment, 0)
	for _, pr := range prs {
		if pr.User.Login != user {
			// 自身が作成したPRのみコメント取得対象とする
			continue
		}

		var comments []Comment

		// 通常のコメント（以降、コメントと呼ぶ）
		normalComments, err := usecase.FetchCommentFromIssues(
			repository,
			owner,
			token,
			pr.Number,
		)
		if err != nil {
			fmt.Println(err)

			return
		}

		for _, com := range normalComments {
			// コメントの追加
			v := Comment{
				Message: com.Comment,
				User:    com.User.Login,
			}

			comments = append(comments, v)
		}

		for _, com := range allPrComments {
			// レビューコメントの追加
			if pr.Url != com.PrUrl {
				continue
			}

			v := Comment{
				Message: com.Comment,
				User:    com.User.Login,
			}

			comments = append(comments, v)
		}

		commentsMap[pr.Title] = comments
	}

	// 結果を見やすいように出力
	for title, comments := range commentsMap {
		fmt.Printf("***** 【PR】%s: comment print start *****\n", title)
		for i, com := range comments {
			fmt.Printf("%d: ", i+1)
			fmt.Printf("(from)%s\t %s\n", com.User, com.Message)
		}

		fmt.Printf("***** 【PR】%s: comment print end *****\n", title)
	}
}
