package article

import (
	"github.com/gin-gonic/gin"
	"server/config/vars"
	"server/models"
	"server/tools"
)

type uploadedFile struct {
	Img string `json:"img"` // file uploaded path
}

func UploadedFile(c *gin.Context) {

	value := c.GetHeader("Authorization")
	rows, _ := vars.DB0.Table("user").Model(&models.User{}).Where("id = ?", tools.Parse(value)).Rows()

	for rows.Next() {
		var (
			user models.User
			u    uploadedFile
		)

		_ = vars.DB0.ScanRows(rows, &user)
		_ = c.ShouldBindJSON(&u)

		img := models.UploadedImg{
			Id:  user.Id,
			Img: u.Img,
		}

		vars.DB0.Table("uploaded_img").Create(&img)

		c.JSON(200, gin.H{
			"message": "file uploaded successfully",
		})
	}
}
