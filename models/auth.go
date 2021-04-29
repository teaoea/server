package models

type Auth struct {
	Id                int64 `json:"id" gorm:"id"` // 被授权用户id
	HideArticle       bool  `json:"hide_article" gorm:"hide_article"`
	HideArticleAuth   int64 `json:"hide_article_auth" gorm:"hide_article_auth"` // 谁授予用户隐藏文章的权限
	HideUser          bool  `json:"hide_user" gorm:"hide_user"`
	HideUserAuth      int64 `json:"hide_user_auth" gorm:"hide_user_auth"` // 谁授予用户隐藏用户的权限
	DeleteComment     bool  `json:"delete_comment" gorm:"delete_comment"`
	DeleteCommentAuth int64 `json:"delete_comment_auth" gorm:"delete_comment_auth"` // 谁授权用户删除评论的权限
}

func (Auth) TableName() string {
	return "auth"
}
