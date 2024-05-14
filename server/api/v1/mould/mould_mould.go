package mould

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
    "github.com/flipped-aurora/gin-vue-admin/server/model/mould"
    mouldReq "github.com/flipped-aurora/gin-vue-admin/server/model/mould/request"
    "github.com/flipped-aurora/gin-vue-admin/server/model/common/response"
    "github.com/flipped-aurora/gin-vue-admin/server/service"
    "github.com/gin-gonic/gin"
    "go.uber.org/zap"
)

type MouldMouldApi struct {
}

var mouldMouldService = service.ServiceGroupApp.MouldServiceGroup.MouldMouldService


// CreateMouldMould 创建模具表
// @Tags MouldMould
// @Summary 创建模具表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body mould.MouldMould true "创建模具表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"创建成功"}"
// @Router /mouldMould/createMouldMould [post]
func (mouldMouldApi *MouldMouldApi) CreateMouldMould(c *gin.Context) {
	var mouldMould mould.MouldMould
	err := c.ShouldBindJSON(&mouldMould)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	if err := mouldMouldService.CreateMouldMould(&mouldMould); err != nil {
        global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// DeleteMouldMould 删除模具表
// @Tags MouldMould
// @Summary 删除模具表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body mould.MouldMould true "删除模具表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /mouldMould/deleteMouldMould [delete]
func (mouldMouldApi *MouldMouldApi) DeleteMouldMould(c *gin.Context) {
	ID := c.Query("ID")
	if err := mouldMouldService.DeleteMouldMould(ID); err != nil {
        global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// DeleteMouldMouldByIds 批量删除模具表
// @Tags MouldMould
// @Summary 批量删除模具表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Success 200 {string} string "{"success":true,"data":{},"msg":"批量删除成功"}"
// @Router /mouldMould/deleteMouldMouldByIds [delete]
func (mouldMouldApi *MouldMouldApi) DeleteMouldMouldByIds(c *gin.Context) {
	IDs := c.QueryArray("IDs[]")
	if err := mouldMouldService.DeleteMouldMouldByIds(IDs); err != nil {
        global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败", c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

// UpdateMouldMould 更新模具表
// @Tags MouldMould
// @Summary 更新模具表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body mould.MouldMould true "更新模具表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /mouldMould/updateMouldMould [put]
func (mouldMouldApi *MouldMouldApi) UpdateMouldMould(c *gin.Context) {
	var mouldMould mould.MouldMould
	err := c.ShouldBindJSON(&mouldMould)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	if err := mouldMouldService.UpdateMouldMould(mouldMould); err != nil {
        global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// FindMouldMould 用id查询模具表
// @Tags MouldMould
// @Summary 用id查询模具表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query mould.MouldMould true "用id查询模具表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /mouldMould/findMouldMould [get]
func (mouldMouldApi *MouldMouldApi) FindMouldMould(c *gin.Context) {
	ID := c.Query("ID")
	if remouldMould, err := mouldMouldService.GetMouldMould(ID); err != nil {
        global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(gin.H{"remouldMould": remouldMould}, c)
	}
}

// GetMouldMouldList 分页获取模具表列表
// @Tags MouldMould
// @Summary 分页获取模具表列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query mouldReq.MouldMouldSearch true "分页获取模具表列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /mouldMould/getMouldMouldList [get]
func (mouldMouldApi *MouldMouldApi) GetMouldMouldList(c *gin.Context) {
	var pageInfo mouldReq.MouldMouldSearch
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if list, total, err := mouldMouldService.GetMouldMouldInfoList(pageInfo); err != nil {
	    global.GVA_LOG.Error("获取失败!", zap.Error(err))
        response.FailWithMessage("获取失败", c)
    } else {
        response.OkWithDetailed(response.PageResult{
            List:     list,
            Total:    total,
            Page:     pageInfo.Page,
            PageSize: pageInfo.PageSize,
        }, "获取成功", c)
    }
}

// GetMouldMouldPublic 不需要鉴权的模具表接口
// @Tags MouldMould
// @Summary 不需要鉴权的模具表接口
// @accept application/json
// @Produce application/json
// @Param data query mouldReq.MouldMouldSearch true "分页获取模具表列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /mouldMould/getMouldMouldList [get]
func (mouldMouldApi *MouldMouldApi) GetMouldMouldPublic(c *gin.Context) {
    // 此接口不需要鉴权
    // 示例为返回了一个固定的消息接口，一般本接口用于C端服务，需要自己实现业务逻辑
    response.OkWithDetailed(gin.H{
       "info": "不需要鉴权的模具表接口信息",
    }, "获取成功", c)
}
