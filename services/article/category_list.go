package article

import (
	"net/http"

	"server/config/vars"
	"server/models"

	"github.com/gin-gonic/gin"
)

type category struct {
	Id   string `json:"id"`
	Name string `json:"name"`
}

// CategoryList
/// return category list
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
