package article

import (
	"Server/config/vars"
	"Server/models"
	"github.com/gin-gonic/gin"
	"net/http"
)

type category struct {
	Id   string `json:"id"`
	Name string `json:"name"`
}

// CategoryList  返回分类列表
func CategoryList(ctx *gin.Context) {
	rows, _ := vars.DB0.Table("category").Model(&models.Category{}).Rows()
	for rows.Next() {
		var ca models.Category
		_ = vars.DB0.ScanRows(rows, &ca)
		ctx.JSON(http.StatusForbidden, category{
			Id:   ca.Id,
			Name: ca.Name,
		})
	}
}
