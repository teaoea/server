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
				Save string `json:"save"`
			}
		)

		_ = vars.DB0.ScanRows(rows, &user)
		_ = c.ShouldBindJSON(&article)
		affected := vars.DB0.Table("category").Where(&models.Category{Name: article.Category}, "name").Find(&category).RowsAffected

		switch {
		case !user.IsActive:
			c.SecureJSON(460, gin.H{
				"message": "The account isn't activated",
			})
		case len(article.Title) >= 90:
			c.SecureJSON(461, gin.H{
				"message": "The title is too long",
			})
		case article.Title == "":
			c.SecureJSON(462, gin.H{
				"message": "The title can't be blank",
			})
		case affected == 0:
			c.SecureJSON(463, gin.H{
				"message": "The category isn't exist",
			})
		default:
			switch article.Save {
			case "draft":
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
				c.SecureJSON(200, gin.H{
					"message": "The article has been saved",
				})
			case "public":
				content := tools.WriteMd(fmt.Sprintf("./static/article/public/%d", user.Id), article.Content)
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
					"message": "The article has been saved",
				})
			default:
				c.SecureJSON(404, nil)
			}
		}
	}
}
