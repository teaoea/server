package article

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"server/config/vars"
	"server/models"
	"server/tools"
	"time"
)

// ReplyComment
/// 回复评论
func ReplyComment(c *gin.Context) {
	value := c.GetHeader("Authorization")
	rows, _ := vars.DB0.Table("user").Model(&models.User{}).Where("id = ?", tools.Parse(value)).Rows()

	for rows.Next() {
		var (
			user    models.User
			comment models.CommentTwo
			one     models.Comment
		)

		_ = vars.DB0.ScanRows(rows, &user)
		_ = c.ShouldBindJSON(&comment)
		affected := vars.DB0.Table("reply").Where(&models.Comment{Id: comment.Comment}, "id").Find(&one).RowsAffected

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
			content := tools.WriteMd("./static/reply", comment.Content)
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
