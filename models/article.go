package models

type Category struct {
	Id   string `json:"id" gorm:"id"`
	Name string `json:"name" gorm:"name"`
}

func (Category) TableName() string {
	return "category"
}

type UploadedImg struct {
	Id  int64  `json:"id" gorm:"id"`
	Img string `json:"img" gorm:"img"`
}

func (UploadedImg) TableName() string {
	return "uploaded_img"
}

type Article struct {
	Id        int64
	Title     string `json:"title"`    // subject
	Content   string `json:"content"`  // content
	Img       string `json:"img"`      // img
	Category  string `json:"category"` // category
	Show      bool   `json:"show"`     // is the article public,public: true,private: false
	View      int64  `json:"view"`     // number of read
	SHA256    string // article sha256 checksum
	Author    string `json:"author"`  // author
	License   string `json:"license"` // license
	IsHide    bool   // whether the article is hidden, hidden: true, isn't hidden: false,default: false
	CreatedAt string `json:"created_at"` // created time
}

func (Article) TableName() string {
	return "article"
}

type Comment struct {
	Id        int64
	Title     int64  `json:"title"`      // article
	User      string `json:"user"`       // user
	Content   string `json:"content"`    // content
	CreatedAt string `json:"created_at"` // created time
}

func (Comment) TableName() string {
	return "comment"
}

type Reply struct {
	Id        int64
	Comment   int64  `json:"comment"`    // comment
	User      string `json:"user"`       // user
	Content   string `json:"content"`    // content
	CreatedAt string `json:"created_at"` // created time
}

func (Reply) TableName() string {
	return "reply"
}
