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
	Title     string `json:"title"`    // 标题
	Content   string `json:"content"`  // 内容
	Img       string `json:"img"`      // 封面图
	Category  string `json:"category"` // 分类
	Show      bool   `json:"show"`     // 文章属性 true: 公开 false: 私有
	View      int64  `json:"view"`     // 阅读量
	SHA256    string // 文章sha256校验和
	Author    string `json:"author"`  // 作者
	License   string `json:"license"` // 许可协议
	IsHide    bool   // 文章是否被拉黑 true: 拉黑 false: No
	CreatedAt string `json:"created_at"` // 创建时间
}

func (Article) TableName() string {
	return "article"
}

type Comment struct {
	Id      int64
	Title   int64  `json:"title"`   // 文章
	User    string `json:"user"`    // 用户
	Content string `json:"content"` // 内容
	Time    string `json:"time"`    // 时间
}

func (Comment) TableName() string {
	return "comment"
}
