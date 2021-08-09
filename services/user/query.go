package user

import (
	"fmt"

	"server/config/vars"
	"server/models"

	"github.com/gin-gonic/gin"
)

func check(key string) bool {
	for _, value := range vars.Query {
		if key == value {
			return true
		}
	}
	return false
}

// Query user table public query api
// true: this record exists
// false: this record doesn't exists
func Query(c *gin.Context) {
	var (
		user  models.User
		query struct {
			Key   string `json:"key"`   // field name
			Value string `json:"value"` // the value to be queried
		}
	)

	_ = c.ShouldBindJSON(&query)

	if query.Key == "" || query.Value == "" {
		c.SecureJSON(461, gin.H{
			"message": "Key and value can't be empty",
		})
		return
	}

	if !check(query.Key) {
		c.SecureJSON(462, gin.H{
			"message": "Can't query this key",
		})
		return
	}

	affected := vars.DB0.Table("user").Where(fmt.Sprintf("%s = ?", query.Key), query.Value).Find(&user).RowsAffected

	if affected == 1 {
		c.SecureJSON(200, gin.H{
			"message": true,
		})
		return
	}

	c.SecureJSON(200, gin.H{
		"message": false,
	})

}
