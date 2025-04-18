package models

type Common struct {
	CreateByID int64    `gorm:"comment:创建人"`
	CreateBy   *SysUser `gorm:"foreignKey:CreateByID;references:ID;comment:创建人"`
	UpdateByID int64    `gorm:"comment:更新人"`
	UpdateBy   *SysUser `gorm:"foreignKey:UpdateByID;references:ID;comment:更新人"`
}
