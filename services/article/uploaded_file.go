package article

import (
	vars2 "Server/config/vars"
	"Server/models"
	"Server/tools/token"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

type uploadedFile struct {
	Img string `json:"img"` // 前端上传文件路径
}

func UploadedFile(c *gin.Context) {
	u := uploadedFile{}
	t := c.GetHeader("Authorization")
	parse := token.Parse(t)
	id := parse.(jwt.MapClaims)["id"]
	rows, _ := vars2.DB0.Table("user").Model(&models.User{}).Where("id = ?", id).Rows()

	for rows.Next() {
		var user models.User
		_ = vars2.DB0.ScanRows(rows, &user)
		_ = c.ShouldBindJSON(&u)

		img := models.UploadedImg{
			Id:  user.Id,
			Img: u.Img,
		}

		vars2.DB0.Table("uploaded_img").Create(&img)

		c.JSON(200, gin.H{
			"message": "图片上传成功",
		})
	}
}
