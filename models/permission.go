package models

type Permission struct {
	Id                int    `json:"id" gorm:"id"`
	UserId            int64  `json:"user_id" gorm:"user_id"` // 被授权用户id
	Name              string `json:"name" gorm:"name"`       // 被授权用户
	HideArticle       bool   `json:"hide_article" gorm:"hide_article"`
	HideArticleAuth   string `json:"hide_article_auth" gorm:"hide_article_auth"` // 谁授予用户隐藏文章的权限
	HideUser          bool   `json:"hide_user" gorm:"hide_user"`
	HideUserAuth      string `json:"hide_user_auth" gorm:"hide_user_auth"` // 谁授予用户隐藏用户的权限
	DeleteComment     bool   `json:"delete_comment" gorm:"delete_comment"`
	DeleteCommentAuth string `json:"delete_comment_auth" gorm:"delete_comment_auth"` // 谁授权用户删除评论的权限
}

func (Permission) TableName() string {
	return "permission"
}
