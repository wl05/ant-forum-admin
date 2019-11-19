package model

import "ant-forum/pkg/constvar"

type RoleModel struct {
	BaseModel
	Name string `json:"name" gorm:"column:name;not null" binding:"required"`
}

func (r *RoleModel) TableName() string {
	return "role"
}

type RoleInfo struct {
	Id   uint64 `json:"id"`
	Name string `json:"name"`
}

// 创建新角色
func (t *RoleModel) Create() error {
	return DB.Self.Create(&t).Error
}

// 获取角色列表
func ListRole(offset, limit int) ([]*RoleModel, uint64, error) {
	t := RoleModel{}
	if limit == 0 {
		limit = constvar.DefaultLimit
	}
	list := make([]*RoleModel, 0)
	var count uint64
	if err := DB.Self.Model(&t).Count(&count).Error; err != nil {
		return list, count, err
	}
	if err := DB.Self.Where("").Offset(offset).Limit(limit).Order("id desc").Find(&t).Error; err != nil {
		return list, count, err
	}
	return list, count, nil
}

// 根据标签id获取角色
func (t *RoleModel) GetRoleById(id uint64) (*RoleModel, error) {
	d := DB.Self.First(&t, id)
	return t, d.Error
}

// 根据标签id删除角色
func (t *RoleModel) DeleteRole(id uint64) error {
	t.BaseModel.Id = id
	return DB.Self.Delete(&t).Error
}
