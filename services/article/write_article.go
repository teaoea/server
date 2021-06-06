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
			user    models.User
			article struct {
				models.Article
				Status bool `json:"status"` // article status, ture: success false: draft
			}
		)

		_ = vars.DB0.ScanRows(rows, &user)
		_ = c.ShouldBindJSON(&article)
		if !user.IsActive {
			c.SecureJSON(403, gin.H{
				"message": "account isn't activated",
			})
			return
		}
		if len(article.Title) > 90 || article.Title == "" {
			c.SecureJSON(411, gin.H{
				"message": "subject is too long",
			})
			return
		}

		var ca models.Category
		caCheck := vars.DB0.Table("category").Where(&models.Category{Name: article.Category}, "name").Find(&ca).RowsAffected
		if caCheck == 0 {
			c.SecureJSON(404, gin.H{
				"message": fmt.Sprintf("category \"%s\" isn't exist", article.Category),
			})
			return
		}
		content := tools.WriteMd("./static/article", article.Content)
		switch {
		case !article.Status:
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
			c.SecureJSON(200, gin.H{
				"message": "the article has been saved to the draft box",
			})
		default:
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
			c.SecureJSON(200, gin.H{
				"message": "the article has been published",
			})
		}
	}
}
