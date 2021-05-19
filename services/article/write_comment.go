package article

import (
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"server/config/vars"
	"server/models"
	"server/services/user/auth"
	"server/tools"
	"time"
)

// WriteArticleComment
/// 评论无法删除
func WriteArticleComment(c *gin.Context) {
	token := c.GetHeader("Authorization")
	parse := auth.Parse(token)
	id := parse.(jwt.MapClaims)["id"]
	rows, _ := vars.DB0.Table("user").Model(&models.User{}).Where("id = ?", id).Rows()

	for rows.Next() {
		var user models.User
		var article models.Article
		var comment models.Comment

		_ = vars.DB0.ScanRows(rows, &user)
		_ = c.ShouldBindJSON(&comment)
		affected := vars.DB0.Table("article").Where(&models.Article{Id: comment.Title}, "id").Find(&article).RowsAffected
		switch {

		case !user.IsActive:
			c.SecureJSON(403, gin.H{
				"message": "账户未激活",
			})

		case len(comment.Content) > 300:
			c.SecureJSON(411, gin.H{
				"message": "评论内容过长",
			})

		case affected == 0:
			c.SecureJSON(404, gin.H{
				"message": fmt.Sprintf("文章不存在：%d", comment.Title),
			})

		default:
			content := tools.WriteMd("./static/comment", comment.Content)
			vars.DB0.Table("comment").Create(&models.Comment{
				Id:      tools.NewId(),
				Title:   comment.Title,
				Content: content,
				User:    user.Name,
				Time:    time.Now().Format("2006-01-02 15:04:05"),
			})

			c.SecureJSON(200, nil)
		}
	}
}

// CommentToComment
/// 评论的评论
func CommentToComment(c *gin.Context) {
	token := c.GetHeader("Authorization")
	parse := auth.Parse(token)
	id := parse.(jwt.MapClaims)["id"]
	rows, _ := vars.DB0.Table("user").Model(&models.User{}).Where("id = ?", id).Rows()

	for rows.Next() {
		var user models.User
		var comment models.CommentTwo
		var one models.Comment

		_ = vars.DB0.ScanRows(rows, &user)
		_ = c.ShouldBindJSON(&comment)
		affected := vars.DB0.Table("comment").Where(&models.Comment{Id: comment.Comment}, "id").Find(&one).RowsAffected

		switch {
		case !user.IsActive:
			c.SecureJSON(403, gin.H{
				"message": "账户未激活",
			})

		case len(comment.Content) > 300:
			c.SecureJSON(411, gin.H{
				"message": "评论内容过长",
			})

		case affected == 0:
			c.SecureJSON(404, gin.H{
				"message": fmt.Sprintf("评论不存在：%d", comment.Comment),
			})

		default:
			content := tools.WriteMd("./static/commenttwo", comment.Content)
			vars.DB0.Table("comment_two").Create(&models.CommentTwo{
				Id:      tools.NewId(),
				Comment: comment.Comment,
				Content: content,
				User:    user.Name,
				Time:    time.Now().Format("2006-01-02 15:04:05"),
			})

			c.SecureJSON(200, nil)
		}
	}
}
