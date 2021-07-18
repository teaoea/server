package modify

import (
	"server/config/vars"
	"server/models"
	"server/tools"

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
			c.SecureJSON(460, gin.H{
				"message": "Email address is already used",
			})
			return
		}

		if !tools.SuffixCheck(email.Email) {
			c.SecureJSON(461, gin.H{
				"message": "The suffix of this email address can't be bound to the account",
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

		c.SecureJSON(200, gin.H{
			"message": "Modify the email address successfully",
		})
	}
}
