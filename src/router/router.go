package router

import (
	"path/filepath"

	"github.com/spf13/viper"

	"github.com/gin-gonic/gin"
	"github.com/marshhu/ma-api/docs"
	"github.com/marshhu/ma-api/src/interface/controller"
	"github.com/marshhu/ma-frame/utils"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// Register 向 Gin 注册业务路由
func Register() func(rg *gin.Engine) error {
	return func(rg *gin.Engine) error {
		rg.Use(gin.Logger())
		// Recovery middleware recovers from any panics and writes a 500 if there was one.
		rg.Use(gin.Recovery())

		rg.GET("/", func(c *gin.Context) {
			c.JSON(200, "OK")
		})

		staticRouter(rg)
		swaggerRouter(rg)
		bizRouter(rg)
		gin.SetMode(viper.GetString("app.runMode"))
		return nil
	}
}

// 静态资源
func staticRouter(eng *gin.Engine) {
	staticGroup := eng.Group("/assets")
	rootDir := utils.RootDir()
	staticGroup.Static("", filepath.Join(rootDir, "static"))
}

//swagger
func swaggerRouter(eng *gin.Engine) {
	docs.SwaggerInfo.Title = "MA API"
	docs.SwaggerInfo.Description = "MA API."
	docs.SwaggerInfo.Version = "1.0"
	docs.SwaggerInfo.BasePath = "/api/v1"
	docs.SwaggerInfo.Schemes = []string{"http", "https"}
	eng.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
}

//业务路由
func bizRouter(eng *gin.Engine) {
	apiV1 := eng.Group("/api/v1")
	apiV1.POST("/auth/login", controller.Login())

	jwtRg := apiV1.Group("", AuthMiddleware())
	{
		jwtRg.POST("/user/register", controller.Register())

		jwtRg.GET("/bill", controller.GetBillByUser())
		jwtRg.POST("/bill", controller.AddBill())
	}
}
