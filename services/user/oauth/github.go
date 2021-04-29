package oauth

import (
	"Server/models"
	"Server/tools/token"
	"Server/vars"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

type github struct {
	GithubId string `json:"github_id"`
	Name     string `json:"name"`
	Email    string `json:"email"`
}

func Github(c *gin.Context) {
	t := c.GetHeader("Authorization")
	parse := token.Parse(t)
	id := parse.(jwt.MapClaims)["jti"]
	g := github{}
	rows, _ := vars.PDB0.Table("user").Model(&models.User{}).Where("id = ?", id).Rows()

	for rows.Next() {
		var user models.User
		_ = vars.PDB0.ScanRows(rows, &user)
		_ = c.ShouldBindJSON(&g)

		github := models.Github{
			Id:       user.Id,
			GithubId: g.GithubId,
			Name:     g.Name,
			Email:    g.Email,
		}
		vars.PDB0.Table("github").Create(&github)
		c.SecureJSON(200, gin.H{
			"message": "github账户绑定成功",
		})
	}
}
