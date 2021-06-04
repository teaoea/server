package article

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"server/config/vars"
	"server/models"
	"server/tools"
	"time"
)

// CommentArticle
/// user can't delete their own comment
func CommentArticle(c *gin.Context) {
	value := c.GetHeader("Authorization")
	id := tools.Parse(value)
	if id == 0 {
		c.SecureJSON(403, "not sign in,please sign in and visit again")
		return
	}
	rows, _ := vars.DB0.Table("user").Model(&models.User{}).Where("id = ?", tools.Parse(value)).Rows()
	for rows.Next() {
		var (
			user    models.User
			article models.Article
			comment models.Comment
		)

		_ = vars.DB0.ScanRows(rows, &user)
		_ = c.ShouldBindJSON(&comment)
		affected := vars.DB0.Table("article").Where(&models.Article{Id: comment.Title}, "id").Find(&article).RowsAffected
		switch {

		case !user.IsActive:
			c.SecureJSON(403, gin.H{
				"message": "account isn't activated",
			})

		case len(comment.Content) > 300:
			c.SecureJSON(411, gin.H{
				"message": "comment content is too long",
			})

		case affected == 0:
			c.SecureJSON(404, gin.H{
				"message": fmt.Sprintf("article \"%d\" don't exist", comment.Title),
			})

		default:
			content := tools.WriteMd("./static/comment", comment.Content)
			vars.DB0.Table("comment").Create(&models.Comment{
				Id:        tools.NewId(),
				Title:     comment.Title,
				Content:   content,
				User:      user.Username,
				CreatedAt: time.Now().Format("2006-01-02 15:04:05"),
			})

			c.SecureJSON(200, nil)
		}
	}
}
