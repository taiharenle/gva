// 自动生成模板MouldMould
package mould

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"gorm.io/datatypes"
	"time"
)

// 模具表 结构体  MouldMould
type MouldMould struct {
	global.GVA_MODEL
	MouldTime      *time.Time     `json:"mouldTime" form:"mouldTime" gorm:"column:mould_time;comment:年份;size:19;" binding:"required"` //年份
	Name           string         `json:"name" form:"name" gorm:"column:name;comment:名称;size:191;" binding:"required"`                //名称
	Size           string         `json:"size" form:"size" gorm:"column:size;comment:门幅;size:191;"`                                   //门幅
	MachinesNumber *int           `json:"machinesNumber" form:"machinesNumber" gorm:"column:machines_number;comment:冲床只数;size:19;"`   //冲床(只/冲)
	PaperNumber    *int           `json:"paperNumber" form:"paperNumber" gorm:"column:paper_number;comment:每张个数;size:19;"`            //个数(只/吸)
	UnitPrice      *int           `json:"unitPrice" form:"unitPrice" gorm:"column:unit_price;comment:单价;size:19;" binding:"required"` //单价(元/只)
	Quantity       *int           `json:"quantity" form:"quantity" gorm:"column:quantity;comment:数量;size:19;"`                        //数量(袋)
	Length         *int           `json:"length" form:"length" gorm:"column:length;comment:长度;size:19;"`                              //长度(毫米mm)
	Weight         *int           `json:"weight" form:"weight" gorm:"column:weight;comment:重量;size:19;"`                              //重量
	Color          string         `json:"color" form:"color" gorm:"column:color;comment:颜色;size:191;"`                                //颜色
	Material       string         `json:"material" form:"material" gorm:"column:material;comment:原料;size:191;"`                       //原料
	MouldPicture   datatypes.JSON `json:"mouldPicture" form:"mouldPicture" gorm:"column:mould_picture;comment:图片;"`                   //图片
}

// TableName 模具表 MouldMould自定义表名 mould_mould
func (MouldMould) TableName() string {
	return "mould_mould"
}
