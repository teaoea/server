package modify

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"server/config/vars"
	"server/models"
	"server/tools"
	"server/tools/mail"
	"strings"
)

type modifyEmail struct {
	Email string `json:"email"`
}

func Email(c *gin.Context) {

	value := c.GetHeader("Authorization")
	rows, _ := vars.DB0.Table("user").Model(&models.User{}).Where("id = ?", tools.Parse(value)).Rows()

	for rows.Next() {
		var (
			user  models.User
			email modifyEmail
		)

		_ = vars.DB0.ScanRows(rows, &user)
		_ = c.BindJSON(&email)

		affected := vars.DB0.Table("user").Where(&models.User{Email: email.Email}, "email").Find(&user).RowsAffected
		if affected != 0 {
			c.SecureJSON(http.StatusUnauthorized, gin.H{
				"message": fmt.Sprintf("邮箱'%s'已注册", email.Email),
			})
			return
		}

		if !mail.SuffixCheck(email.Email) {
			addr := strings.Split(email.Email, "@") // 字符串分割
			suffix := "@" + addr[1]                 // 截取邮箱后缀
			c.SecureJSON(http.StatusForbidden, gin.H{
				"message": fmt.Sprintf("此邮箱后缀'%s'无法绑定账户", suffix),
			})
			return
		}

		// 修改邮箱和邮箱状态
		vars.DB0.Table("user").Model(&models.User{}).Where("id = ?", user.Id).Update("email", email.Email)
		vars.DB0.Table("user").Model(&models.User{}).Where("id = ?", user.Id).Update("email_active", false)
		// 如果手机号未激活,修改账户状态为未激活
		if !user.NumberActive {
			vars.DB0.Table("user").Model(&models.User{}).Where("id = ?", user.Id).Update("is_active", false)
		}

		c.SecureJSON(http.StatusOK, nil)
	}
}
