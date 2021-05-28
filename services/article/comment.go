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
/// 评论无法删除
func CommentArticle(c *gin.Context) {
	value := c.GetHeader("Authorization")
	id := tools.Parse(value)
	if id == 0 {
		c.SecureJSON(403, "登录过期,请重新登录")
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
