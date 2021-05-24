package oauth

import (
	"github.com/gin-gonic/gin"
	"server/config/vars"
	"server/models"
	"server/tools"
)

type github struct {
	GithubId string `json:"github_id"`
	Name     string `json:"name"`
	Email    string `json:"email"`
}

func Github(c *gin.Context) {
	value := c.GetHeader("Authorization")
	rows, _ := vars.DB0.Table("user").Model(&models.User{}).Where("id = ?", tools.Parse(value)).Rows()

	for rows.Next() {
		var (
			user models.User
			g    github
		)

		_ = vars.DB0.ScanRows(rows, &user)
		_ = c.ShouldBindJSON(&g)

		github := models.Github{
			Id:       user.Id,
			GithubId: g.GithubId,
			Name:     g.Name,
			Email:    g.Email,
		}
		vars.DB0.Table("github").Create(&github)
		c.SecureJSON(200, gin.H{
			"message": "github账户绑定成功",
		})
	}
}
