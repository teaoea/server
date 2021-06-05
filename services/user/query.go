package user

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"server/config/vars"
)

// Query
/// user table public query api
/// return
/// false: can't pass
/// true: pass
func Query(c *gin.Context) {
	var query struct {
		Key   string `json:"key"`   // field name
		Value string `json:"value"` // the value to be queried
	}
	_ = c.ShouldBindJSON(&query)

	switch {
	case query.Key == "" || query.Value == "":
		c.SecureJSON(200, gin.H{
			"message": false,
		})

	default:
		affected := vars.DB0.Table("user").Where(fmt.Sprintf("%s = %s", query.Key, query.Value)).RowsAffected
		if affected == 0 {
			c.SecureJSON(200, gin.H{
				"message": true,
			})
		} else {
			c.SecureJSON(200, gin.H{
				"message": false,
			})
		}
	}
}
