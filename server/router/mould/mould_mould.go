package mould

import (
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1"
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type MouldMouldRouter struct {
}

// InitMouldMouldRouter 初始化 模具表 路由信息
func (s *MouldMouldRouter) InitMouldMouldRouter(Router *gin.RouterGroup,PublicRouter *gin.RouterGroup) {
	mouldMouldRouter := Router.Group("mouldMould").Use(middleware.OperationRecord())
	mouldMouldRouterWithoutRecord := Router.Group("mouldMould")
	mouldMouldRouterWithoutAuth := PublicRouter.Group("mouldMould")

	var mouldMouldApi = v1.ApiGroupApp.MouldApiGroup.MouldMouldApi
	{
		mouldMouldRouter.POST("createMouldMould", mouldMouldApi.CreateMouldMould)   // 新建模具表
		mouldMouldRouter.DELETE("deleteMouldMould", mouldMouldApi.DeleteMouldMould) // 删除模具表
		mouldMouldRouter.DELETE("deleteMouldMouldByIds", mouldMouldApi.DeleteMouldMouldByIds) // 批量删除模具表
		mouldMouldRouter.PUT("updateMouldMould", mouldMouldApi.UpdateMouldMould)    // 更新模具表
	}
	{
		mouldMouldRouterWithoutRecord.GET("findMouldMould", mouldMouldApi.FindMouldMould)        // 根据ID获取模具表
		mouldMouldRouterWithoutRecord.GET("getMouldMouldList", mouldMouldApi.GetMouldMouldList)  // 获取模具表列表
	}
	{
	    mouldMouldRouterWithoutAuth.GET("getMouldMouldPublic", mouldMouldApi.GetMouldMouldPublic)  // 获取模具表列表
	}
}
