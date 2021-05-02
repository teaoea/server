package models

type User struct {
	Id           int64  `json:"id" gorm:"id"`
	Name         string `json:"name" gorm:"name"`
	Password     string `json:"password" gorm:"password"`
	Email        string `json:"email" gorm:"email"`
	EmailActive  bool   `json:"email_active" gorm:"email_active"`
	Country      string `json:"country" gorm:"country"`
	Number       string `json:"number" gorm:"number"`
	NumberActive bool   `json:"number_active" gorm:"number_active"`
	Avatar       string `json:"avatar" gorm:"avatar"`
	Gender       string `json:"gender" gorm:"gender"`
	CreatedAt    string `json:"created_at" gorm:"created_at"`
	IsActive     bool   `json:"is_active" gorm:"is_active"`
	IsAdmin      bool   `json:"is_admin" gorm:"is_admin"`
	IsHide       bool   `json:"is_hide" gorm:"is_hide"`
}

func (User) TableName() string {
	return "user"
}

type Github struct {
	Id       int64  `gorm:"id" json:"id"`
	GithubId string `gorm:"github_id" json:"github_id"`
	Name     string `gorm:"name" json:"name"`
	Email    string `gorm:"email" json:"email"`
}

func (Github) TableName() string {
	return "github"
}
