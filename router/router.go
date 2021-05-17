package router

import (
	"github.com/gin-gonic/gin"
	"server/services/article"
	"server/services/permission"
	"server/services/user"
	"server/services/user/auth"
	"server/services/user/email"
	"server/services/user/oauth"
	"server/services/vue"
)

func Router() *gin.Engine {
	router := gin.New()
	router.Use(Server(), Cor())
	v1 := router.Group("/v1")
	v1.GET("/category", article.CategoryList) // 分类列表
	{
		account := v1.Group("/account")
		{
			account.POST("/register", user.Register) //注册
			account.POST("/login", user.Login)       //登录
		}

		accountAuth := v1.Group("/account", LoginAuth())
		{
			accountAuth.GET("/me", user.Me)                           //个人中心
			accountAuth.GET("/email/sendcode", email.SendCode)        //邮箱验证码发送
			accountAuth.POST("/email/active", email.Active)           //邮箱激活
			accountAuth.POST("/email/update", email.Update)           //修改邮箱
			accountAuth.POST("/change_password", user.ChangePassword) //修改密码
			accountAuth.POST("/change_profile", user.UpdateUser)      //修改个人信息
			accountAuth.GET("/token/refresh", user.RefreshToken)      //刷新token
			accountAuth.POST("/oauth/github", oauth.Github)           //github账户绑定
		}

		Article := v1.Group("/article", LoginAuth())
		{
			Article.POST("/write", article.WriteArticle) // 编写文章
		}

		upload := v1.Group("/uploaded", LoginAuth())
		{
			upload.POST("/img", article.UploadedFile) // 上传图片
		}

		Permission := v1.Group("/permission", auth.ProxyAuth(), LoginAuth())
		{
			Permission.POST("/article", permission.HideArticle) // 隐藏文章
		}

		Vue := v1.Group("/vue")
		{
			Vue.POST("/logger", vue.Logger)
		}
	}
	return router
}
