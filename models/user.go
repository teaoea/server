package models

type User struct {
	Id           int64  `json:"id"`
	Name         string `json:"name" `
	Password     string `json:"password"`
	Email        string `json:"email" `
	EmailActive  bool   `json:"email_active" `
	Number       string `json:"number"`
	NumberActive bool   `json:"number_active" `
	Avatar       string `json:"avatar" `
	Gender       string `json:"gender" `
	CreatedAt    string `json:"created_at" `
	IsActive     bool   `json:"is_active" `
	IsAdmin      bool   `json:"is_admin" `
	IsHide       bool   `json:"is_hide"`
}

func (User) TableName() string {
	return "user"
}

type Github struct {
	Id       int64  `json:"id"`
	GithubId string `json:"github_id"`
	Name     string `json:"name"`
	Email    string `json:"email"`
}

func (Github) TableName() string {
	return "github"
}
