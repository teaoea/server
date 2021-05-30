package models

type Permission struct {
	Id              int    `json:"id"`
	UserId          int64  `json:"user_id"`            // authorized user's id
	Name            string `json:"name"`               // authorized username
	HideArticle     bool   `json:"hide_article"`       // hide article
	HideArticleAuth string `json:"hide_article_auth" ` // who authorized user permission to hide articles?
	HideUser        bool   `json:"hide_user"`          // hide user
	HideUserAuth    string `json:"hide_user_auth"`     // who authorized user permission to hide users?
	DelComment      bool   `json:"del_comment"`        // delete comment
	DelCommentAuth  string `json:"del_comment_auth"`   // who authorized user permission to delete comment?
	AddCategory     bool   `json:"add_category"`       // create category
	AddCategoryAuth string `json:"add_category_auth"`  // who authorizes user permission to create category?
	DelCategory     bool   `json:"del_category"`       // delete category
	DelCategoryAuth string `json:"del_category_auth"`  // who authorized user permission to delete category?
}

func (Permission) TableName() string {
	return "permission"
}
