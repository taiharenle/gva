package mould

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/mould"
    mouldReq "github.com/flipped-aurora/gin-vue-admin/server/model/mould/request"
)

type MouldMouldService struct {
}

// CreateMouldMould 创建模具表记录
// Author [piexlmax](https://github.com/piexlmax)
func (mouldMouldService *MouldMouldService) CreateMouldMould(mouldMould *mould.MouldMould) (err error) {
	err = global.GVA_DB.Create(mouldMould).Error
	return err
}

// DeleteMouldMould 删除模具表记录
// Author [piexlmax](https://github.com/piexlmax)
func (mouldMouldService *MouldMouldService)DeleteMouldMould(ID string) (err error) {
	err = global.GVA_DB.Delete(&mould.MouldMould{},"id = ?",ID).Error
	return err
}

// DeleteMouldMouldByIds 批量删除模具表记录
// Author [piexlmax](https://github.com/piexlmax)
func (mouldMouldService *MouldMouldService)DeleteMouldMouldByIds(IDs []string) (err error) {
	err = global.GVA_DB.Delete(&[]mould.MouldMould{},"id in ?",IDs).Error
	return err
}

// UpdateMouldMould 更新模具表记录
// Author [piexlmax](https://github.com/piexlmax)
func (mouldMouldService *MouldMouldService)UpdateMouldMould(mouldMould mould.MouldMould) (err error) {
	err = global.GVA_DB.Model(&mould.MouldMould{}).Where("id = ?",mouldMould.ID).Updates(&mouldMould).Error
	return err
}

// GetMouldMould 根据ID获取模具表记录
// Author [piexlmax](https://github.com/piexlmax)
func (mouldMouldService *MouldMouldService)GetMouldMould(ID string) (mouldMould mould.MouldMould, err error) {
	err = global.GVA_DB.Where("id = ?", ID).First(&mouldMould).Error
	return
}

// GetMouldMouldInfoList 分页获取模具表记录
// Author [piexlmax](https://github.com/piexlmax)
func (mouldMouldService *MouldMouldService)GetMouldMouldInfoList(info mouldReq.MouldMouldSearch) (list []mould.MouldMould, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
    // 创建db
	db := global.GVA_DB.Model(&mould.MouldMould{})
    var mouldMoulds []mould.MouldMould
    // 如果有条件搜索 下方会自动创建搜索语句
    if info.StartCreatedAt !=nil && info.EndCreatedAt !=nil {
     db = db.Where("created_at BETWEEN ? AND ?", info.StartCreatedAt, info.EndCreatedAt)
    }
    if info.Name != "" {
        db = db.Where("name LIKE ?","%"+ info.Name+"%")
    }
	err = db.Count(&total).Error
	if err!=nil {
    	return
    }

	if limit != 0 {
       db = db.Limit(limit).Offset(offset)
    }
	
	err = db.Find(&mouldMoulds).Error
	return  mouldMoulds, total, err
}