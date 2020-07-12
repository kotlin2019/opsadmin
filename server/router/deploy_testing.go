package router

import (
	"gin-vue-admin/api/v1"
	"gin-vue-admin/middleware"
	"github.com/gin-gonic/gin"
)

func InitDeployTestingRouter(Router *gin.RouterGroup) {
	ApiRouter := Router.Group("deploy/test").Use(middleware.JWTAuth()).Use(middleware.CasbinHandler())
	{
		ApiRouter.POST("testingList", v1.TestingList) // 提测列表
	}
}