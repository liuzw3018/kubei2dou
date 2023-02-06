package initialize

import (
	"github.com/gin-gonic/gin"
	"github.com/liuzw3018/kubei2dou/server/docs"
	"github.com/liuzw3018/kubei2dou/server/middleware"
	businessrouter "github.com/liuzw3018/kubei2dou/server/router/business"
	sysrouter "github.com/liuzw3018/kubei2dou/server/router/sys"
	"github.com/liuzw3018/saber/lib"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"

	_ "github.com/liuzw3018/kubei2dou/server/docs"
)

/**
 * @Author: liu zw
 * @Date: 2022/10/20 17:50
 * @File: sys_router.go
 * @Description: //TODO $
 * @Version:
 */

// @title Swagger Example API
// @version 1.0
// @description This is a sample server celler server.
// @termsOfService http://swagger.io/terms/

// @contact.name API Support
// @contact.url http://www.swagger.io/support
// @contact.email support@swagger.io

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @host localhost:8080
// @BasePath /api/v1
// @query.collection.format multi

// @securityDefinitions.basic BasicAuth

// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name Authorization

// @securitydefinitions.oauth2.application OAuth2Application
// @tokenUrl https://example.com/oauth/token
// @scope.write Grants write access
// @scope.admin Grants read and write access to administrative information

// @securitydefinitions.oauth2.implicit OAuth2Implicit
// @authorizationurl https://example.com/oauth/authorize
// @scope.write Grants write access
// @scope.admin Grants read and write access to administrative information

// @securitydefinitions.oauth2.password OAuth2Password
// @tokenUrl https://example.com/oauth/token
// @scope.read Grants read access
// @scope.write Grants write access
// @scope.admin Grants read and write access to administrative information

// @securitydefinitions.oauth2.accessCode OAuth2AccessCode
// @tokenUrl https://example.com/oauth/token
// @authorizationurl https://example.com/oauth/authorize
// @scope.admin Grants read and write access to administrative information

// @x-extension-openapi {"example": "value on a json format"}

func InitRouters(middlewares ...gin.HandlerFunc) *gin.Engine {
	// programatically set swagger info
	docs.SwaggerInfo.Title = lib.GetStringConf("base.swagger.title")
	docs.SwaggerInfo.Description = lib.GetStringConf("base.swagger.desc")
	docs.SwaggerInfo.Version = "1.0"
	docs.SwaggerInfo.Host = lib.GetStringConf("base.swagger.host")
	docs.SwaggerInfo.BasePath = lib.GetStringConf("base.swagger.base_path")
	docs.SwaggerInfo.Schemes = []string{"http", "https"}

	app := gin.Default()
	app.Use(middlewares...)
	app.Use(middleware.Cors())
	app.Use(middleware.RecoveryMiddleware(), middleware.TranslationMiddleware())
	app.GET("/health", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "OK",
		})
	})
	app.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	app.GET("/metrics", gin.WrapH(promhttp.Handler()))
	// 路由注册
	baseApiGroup := app.Group("/api")
	baseApiGroup.Use(middleware.RequestLog())

	// v1版本路由
	v1ApiGroup := baseApiGroup.Group("/v1")
	{
		sysrouter.RegisterSysLoginRouter(v1ApiGroup)
	}
	v1AuthApiGroup := baseApiGroup.Group("/v1")
	v1AuthApiGroup.Use(middleware.JwtAuthMiddleware())
	{
		sysrouter.RegisterSysAdminRouter(v1AuthApiGroup)
		v1K8SAuthApiGroup := v1AuthApiGroup.Group("/k8s")
		{
			businessrouter.RegisterK8SNamespacesRouter(v1K8SAuthApiGroup)
			businessrouter.RegisterK8SNodesRouter(v1K8SAuthApiGroup)
		}
	}
	return app
}
