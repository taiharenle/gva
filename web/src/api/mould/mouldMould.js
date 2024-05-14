import service from '@/utils/request'

// @Tags MouldMould
// @Summary 创建模具表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.MouldMould true "创建模具表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"创建成功"}"
// @Router /mouldMould/createMouldMould [post]
export const createMouldMould = (data) => {
  return service({
    url: '/mouldMould/createMouldMould',
    method: 'post',
    data
  })
}

// @Tags MouldMould
// @Summary 删除模具表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.MouldMould true "删除模具表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /mouldMould/deleteMouldMould [delete]
export const deleteMouldMould = (params) => {
  return service({
    url: '/mouldMould/deleteMouldMould',
    method: 'delete',
    params
  })
}

// @Tags MouldMould
// @Summary 批量删除模具表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除模具表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /mouldMould/deleteMouldMould [delete]
export const deleteMouldMouldByIds = (params) => {
  return service({
    url: '/mouldMould/deleteMouldMouldByIds',
    method: 'delete',
    params
  })
}

// @Tags MouldMould
// @Summary 更新模具表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.MouldMould true "更新模具表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /mouldMould/updateMouldMould [put]
export const updateMouldMould = (data) => {
  return service({
    url: '/mouldMould/updateMouldMould',
    method: 'put',
    data
  })
}

// @Tags MouldMould
// @Summary 用id查询模具表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query model.MouldMould true "用id查询模具表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /mouldMould/findMouldMould [get]
export const findMouldMould = (params) => {
  return service({
    url: '/mouldMould/findMouldMould',
    method: 'get',
    params
  })
}

// @Tags MouldMould
// @Summary 分页获取模具表列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取模具表列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /mouldMould/getMouldMouldList [get]
export const getMouldMouldList = (params) => {
  return service({
    url: '/mouldMould/getMouldMouldList',
    method: 'get',
    params
  })
}
