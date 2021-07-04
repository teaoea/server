package router

import (
	"server/services/angular"
	"server/services/article"
	"server/services/permission"
	"server/services/user"
	"server/services/user/email"
	"server/services/user/modify"
	"server/services/user/oauth"

	"github.com/gin-gonic/gin"
)

func Router() *gin.Engine {
	router := gin.New()
	router.Use(Server(), Cor())
	v1 := router.Group("/v1")
	v1.GET("/category", article.CategoryList)
	{
		accountGroup := v1.Group("/account")
		{
			accountGroup.POST("/signup", user.SignUp)
			accountGroup.POST("/signin", user.SignIn)
			accountGroup.POST("/query", user.Query)
			accountGroup.GET("/me", user.Me, Authorization())
			accountGroup.POST("/logoff", user.Logoff, Authorization())

			emailGroup := accountGroup.Group("/email", Authorization())
			{
				emailGroup.GET("sendcode", email.SendCode)
				emailGroup.POST("/active", email.Active)
			}

			modifyGroup := accountGroup.Group("/modify", Authorization())
			{
				modifyGroup.POST("/password", modify.Password)
				modifyGroup.POST("/email", modify.Email)
				modifyGroup.POST("/profile", modify.Profile)
			}

			oauthGroup := accountGroup.Group("/oauth", Authorization())
			{
				oauthGroup.POST("/github", oauth.Github)
			}
		}

		articleGroup := v1.Group("/article", Authorization())
		{
			articleGroup.POST("/write", article.WriteArticle)
			articleGroup.POST("/comment", article.CommentArticle)
			articleGroup.POST("/comments", article.ReplyComment)
		}

		uploadGroup := v1.Group("/uploaded", Authorization())
		{
			uploadGroup.POST("/img", article.UploadedFile)
		}

		permissionGroup := v1.Group("/permission", ProxyAuth(), Authorization())
		{
			permissionGroup.POST("/article", permission.HideArticle)
		}

		angularGroup := v1.Group("/angular")
		{
			angularGroup.POST("/error", angular.Error)
			angularGroup.GET("/signin_guard", angular.SigninGuard)
		}
	}
	return router
}
