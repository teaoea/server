package article

import (
	"server/config/vars"
	"server/models"

	"github.com/gin-gonic/gin"
)

// CategoryList
/// return category list
func CategoryList(ctx *gin.Context) {
	rows, _ := vars.DB0.Table("category").Model(&models.Category{}).Rows()
	for rows.Next() {
		var category models.Category
		_ = vars.DB0.ScanRows(rows, &category)
		ctx.JSON(200, gin.H{
			"id":   category.Id,
			"name": category.Name,
		})
	}
}
