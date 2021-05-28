package router

import (
	"github.com/gin-gonic/gin"
	"server/services/angular"
	"server/services/article"
	"server/services/permission"
	"server/services/user"
	"server/services/user/email"
	"server/services/user/modify"
	"server/services/user/oauth"
)

func Router() *gin.Engine {
	router := gin.New()
	router.Use(Server(), Cor())
	v1 := router.Group("/v1")
	v1.GET("/category", article.CategoryList) // 分类列表
	{
		accountGroup := v1.Group("/account")
		{
			accountGroup.POST("/signup", user.SignUp)         //注册
			accountGroup.POST("/signin", user.SignIn)         //登录
			accountGroup.GET("/me", user.Me, Authorization()) //个人中心

			emailGroup := accountGroup.Group("/email", Authorization())
			{
				emailGroup.GET("sendcode", email.SendCode) //邮箱验证码发送
				emailGroup.POST("/active", email.Active)   //邮箱激活
			}

			modifyGroup := accountGroup.Group("/modify", Authorization())
			{
				modifyGroup.POST("/password", modify.Password) //修改密码
				modifyGroup.POST("/email", modify.Email)       //修改邮箱
				modifyGroup.POST("/profile", modify.Profile)   //修改个人信息
			}

			oauthGroup := accountGroup.Group("/oauth", Authorization())
			{
				oauthGroup.POST("/github", oauth.Github) //github账户绑定
			}
		}

		Article := v1.Group("/article", Authorization())
		{
			Article.POST("/write", article.WriteArticle)     // 编写文章
			Article.POST("/comment", article.CommentArticle) // 编写评论
			Article.POST("/comments", article.ReplyComment)  // 回复评论
		}

		upload := v1.Group("/uploaded", Authorization())
		{
			upload.POST("/img", article.UploadedFile) // 上传图片
		}

		Permission := v1.Group("/permission", ProxyAuth(), Authorization())
		{
			Permission.POST("/article", permission.HideArticle) // 隐藏文章
		}

		Angular := v1.Group("/angular")
		{
			Angular.POST("/error", angular.Error) // angular错误日志收集
		}
	}
	return router
}
