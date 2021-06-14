package article

import (
	"time"

	"server/config/vars"
	"server/models"
	"server/tools"

	"github.com/gin-gonic/gin"
)

// ReplyComment
/// reply to comment
func ReplyComment(c *gin.Context) {
	value := c.GetHeader("Authorization")
	rows, _ := vars.DB0.Table("user").Model(&models.User{}).Where("id = ?", tools.Parse(value)).Rows()

	for rows.Next() {
		var (
			user    models.User
			reply   models.Reply
			comment models.Comment
		)

		_ = vars.DB0.ScanRows(rows, &user)
		_ = c.ShouldBindJSON(&reply)
		affected := vars.DB0.Table("reply").Where(&models.Comment{Id: reply.Comment}, "id").Find(&comment).RowsAffected

		switch {
		case !user.IsActive:
			c.SecureJSON(452, nil)

		case len(reply.Content) > 300:
			c.SecureJSON(453, nil)

		case affected == 0:
			c.SecureJSON(454, nil)

		default:
			content := tools.WriteMd("./static/reply", reply.Content)
			vars.DB0.Table("reply").Create(&models.Reply{
				Id:        tools.NewId(),
				Comment:   reply.Comment,
				Content:   content,
				User:      user.Username,
				CreatedAt: time.Now().Format("2006-01-02 15:04:05"),
			})

			c.SecureJSON(200, nil)
		}
	}
}
