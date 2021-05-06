package article

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"server/config/vars"
	"server/models"
	"server/services/user/auth"
)

type uploadedFile struct {
	Img string `json:"img"` // 前端上传文件路径
}

func UploadedFile(c *gin.Context) {
	u := uploadedFile{}
	token := c.GetHeader("Authorization")
	parse := auth.Parse(token)
	id := parse.(jwt.MapClaims)["id"]
	rows, _ := vars.DB0.Table("user").Model(&models.User{}).Where("id = ?", id).Rows()

	for rows.Next() {
		var user models.User
		_ = vars.DB0.ScanRows(rows, &user)
		_ = c.ShouldBindJSON(&u)

		img := models.UploadedImg{
			Id:  user.Id,
			Img: u.Img,
		}

		vars.DB0.Table("uploaded_img").Create(&img)

		c.JSON(200, gin.H{
			"message": "图片上传成功",
		})
	}
}
