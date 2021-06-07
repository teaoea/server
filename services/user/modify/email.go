package modify

import (
	"fmt"
	"net/http"
	"strings"

	"server/config/vars"
	"server/models"
	"server/tools"
	"server/tools/mail"

	"github.com/gin-gonic/gin"
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
				"message": fmt.Sprintf("email address \"%s\" is already used", email.Email),
			})
			return
		}

		if !mail.SuffixCheck(email.Email) {
			addr := strings.Split(email.Email, "@") // string segmentation
			suffix := "@" + addr[1]                 // intercept email address suffix
			c.SecureJSON(http.StatusForbidden, gin.H{
				"message": fmt.Sprintf("the suffix \"%s\" of this email address can't be bound to the account", suffix),
			})
			return
		}

		// modify email address
		vars.DB0.Table("user").Model(&models.User{}).Where("id = ?", user.Id).Update("email", email.Email)
		// modify email address status
		vars.DB0.Table("user").Model(&models.User{}).Where("id = ?", user.Id).Update("email_active", false)
		// if the phone number isn't activated,
		// modify account status to not activated
		if !user.NumberActive {
			vars.DB0.Table("user").Model(&models.User{}).Where("id = ?", user.Id).Update("is_active", false)
		}

		c.SecureJSON(http.StatusOK, nil)
	}
}