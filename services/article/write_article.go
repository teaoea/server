package article

import (
	"context"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"net/http"
	"server/config/vars"
	"server/models"
	"server/services/user/auth"
	"server/tools"
	"time"
)

func WriteArticle(c *gin.Context) {

	token := c.GetHeader("Authorization")
	parse := auth.Parse(token)
	id := parse.(jwt.MapClaims)["id"]
	rows, _ := vars.DB0.Table("user").Model(&models.User{}).Where("id = ?", id).Rows()

	for rows.Next() {
		var user models.User
		var article models.Article
		_ = vars.DB0.ScanRows(rows, &user)
		_ = c.ShouldBindJSON(&article)
		if !user.IsActive {
			c.SecureJSON(http.StatusUnauthorized, gin.H{
				"message": "账户未激活",
			})
			return
		}
		if len(article.Title) > 90 || article.Title == "" {
			c.SecureJSON(http.StatusForbidden, gin.H{
				"message": "标题过长",
			})
			return
		}

		var ca models.Category
		caCheck := vars.DB0.Table("category").Where(&models.Category{Name: article.Category}, "name").Find(&ca).RowsAffected
		if caCheck == 0 {
			c.SecureJSON(http.StatusNotFound, gin.H{
				"message": fmt.Sprintf("类别'%s'不存在", article.Category),
			})
			return
		}
		body := tools.WriteMd(article.Body)
		switch {
		case !article.Status:
			_, _ = vars.MongoDraft.InsertOne(context.Background(), bson.D{
				bson.E{Key: "_id", Value: tools.NewId()},
				bson.E{Key: "title", Value: article.Title},
				bson.E{Key: "body", Value: body},
				bson.E{Key: "img", Value: article.Img},
				bson.E{Key: "category", Value: article.Category},
				bson.E{Key: "show", Value: article.Show},
				bson.E{Key: "view", Value: 0},
				bson.E{Key: "sha256", Value: tools.Checksum(article.Body)},
				bson.E{Key: "author", Value: user.Name},
				bson.E{Key: "license", Value: article.License},
				bson.E{Key: "is_hide", Value: false},
				bson.E{Key: "created_at", Value: time.Now().Format("2006-01-02 15:04:05")},
			})
			c.SecureJSON(http.StatusOK, gin.H{
				"message": "文章已保存到草稿箱",
			})
			return
		default:
			_, _ = vars.MongoPublish.InsertOne(context.Background(), bson.D{
				bson.E{Key: "_id", Value: tools.NewId()},
				bson.E{Key: "title", Value: article.Title},
				bson.E{Key: "body", Value: body},
				bson.E{Key: "img", Value: article.Img},
				bson.E{Key: "category", Value: article.Category},
				bson.E{Key: "show", Value: article.Show},
				bson.E{Key: "view", Value: 0},
				bson.E{Key: "sha256", Value: tools.Checksum(article.Body)},
				bson.E{Key: "author", Value: user.Name},
				bson.E{Key: "license", Value: article.License},
				bson.E{Key: "is_hide", Value: false},
				bson.E{Key: "created_at", Value: time.Now().Format("2006-01-02 15:04:05")},
			})
			c.SecureJSON(http.StatusOK, gin.H{
				"message": "文章已发布",
			})
		}
	}
}
