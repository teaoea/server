package email

import (
	vars2 "Server/config/vars"
	"Server/models"
	"Server/tools"
	"Server/tools/token"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
)

type emailUpdate struct {
	Email string `json:"email"`
}

func Update(c *gin.Context) {
	t := c.GetHeader("Authorization")
	parse := token.Parse(t)
	id := parse.(jwt.MapClaims)["id"]
	rows, _ := vars2.DB0.Table("user").Model(&models.User{}).Where("id = ? ", id).Rows()

	for rows.Next() {
		var (
			user models.User
			e    emailUpdate
		)

		_ = vars2.DB0.ScanRows(rows, &user)
		_ = c.BindJSON(&e)

		affected := vars2.DB0.Table("user").Where(&models.User{Email: e.Email}, "email").Find(&user).RowsAffected
		if affected != 0 {
			c.SecureJSON(http.StatusUnauthorized, gin.H{
				"message": fmt.Sprintf("邮箱'%s'已注册", e.Email),
			})
			return
		}

		if !tools.SuffixCheck(e.Email) {
			addr := strings.Split(e.Email, "@") // 字符串分割
			suffix := "@" + addr[1]             // 截取邮箱后缀
			c.SecureJSON(http.StatusForbidden, gin.H{
				"message": fmt.Sprintf("此邮箱后缀'%s'无法绑定账户", suffix),
			})
			return
		}

		// 修改邮箱和邮箱状态
		vars2.DB0.Table("user").Model(&models.User{}).Where("id = ?", user.Id).Update("email", e.Email)
		vars2.DB0.Table("user").Model(&models.User{}).Where("id = ?", user.Id).Update("email_active", false)
		// 如果手机号未激活,修改账户状态为未激活
		if !user.NumberActive {
			vars2.DB0.Table("user").Model(&models.User{}).Where("id = ?", user.Id).Update("is_active", false)
		}

		c.SecureJSON(http.StatusOK, nil)
	}
}
