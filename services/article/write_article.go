package article

import (
	"fmt"
	"time"

	"server/config/vars"
	"server/models"
	"server/tools"

	"github.com/gin-gonic/gin"
)

func WriteArticle(c *gin.Context) {

	value := c.GetHeader("Authorization")
	rows, _ := vars.DB0.Table("user").Model(&models.User{}).Where("id = ?", tools.Parse(value)).Rows()

	for rows.Next() {
		var (
			user     models.User
			category models.Category
			article  struct {
				models.Article
				Status bool `json:"status"` // article status, ture: success false: draft
			}
		)

		_ = vars.DB0.ScanRows(rows, &user)
		_ = c.ShouldBindJSON(&article)
		affected := vars.DB0.Table("category").Where(&models.Category{Name: article.Category}, "name").Find(&category).RowsAffected

		switch {
		case !user.IsActive:
			c.SecureJSON(452, nil)
		case len(article.Title) >= 90:
			c.SecureJSON(453, nil)
		case article.Title == "":
			c.SecureJSON(454, nil)
		case affected == 0:
			c.SecureJSON(455, nil)
		default:
			if !article.Status {
				content := tools.WriteMd(fmt.Sprintf("./static/article/draft/%d", user.Id), article.Content)
				a := models.Article{
					Id:        tools.NewId(),
					Title:     article.Title,
					Content:   content,
					Img:       article.Img,
					Category:  article.Category,
					Show:      article.Show,
					View:      0,
					SHA256:    tools.Checksum(article.Content),
					Author:    user.Username,
					License:   article.License,
					CreatedAt: time.Now().Format("2006-01-02 15:04:05"),
				}
				vars.DB0.Table("draft").Create(&a)
				c.SecureJSON(230, nil)
			} else {
				content := tools.WriteMd(fmt.Sprintf("./static/article/%d", user.Id), article.Content)
				a := models.Article{
					Id:        tools.NewId(),
					Title:     article.Title,
					Content:   content,
					Img:       article.Img,
					Category:  article.Category,
					Show:      article.Show,
					View:      0,
					SHA256:    tools.Checksum(article.Content),
					Author:    user.Username,
					License:   article.License,
					CreatedAt: time.Now().Format("2006-01-02 15:04:05"),
				}
				vars.DB0.Table("article").Create(&a)
				c.SecureJSON(231, nil)
			}
		}
	}
}
