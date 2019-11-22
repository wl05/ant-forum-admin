package model

import "ant-forum/pkg/constvar"

type RoleModel struct {
	BaseModel
	Name string `form:"name" json:"name" gorm:"column:name;not null" binding:"required"`
}

func (role *RoleModel) TableName() string {
	return "role"
}

type RoleInfo struct {
	Id   uint64 `json:"id"`
	Name string `json:"name"`
}

// 创建新角色
func (role *RoleModel) Create() error {
	return DB.Self.Create(&role).Error
}

// 获取角色列表
func (role *RoleModel) ListRole(offset, limit int) ([]*RoleModel, uint64, error) {
	t := RoleModel{}
	if limit == 0 {
		limit = constvar.DefaultLimit
	}
	list := make([]*RoleModel, 0)
	var count uint64
	if err := DB.Self.Model(&t).Count(&count).Error; err != nil {
		return list, count, err
	}
	if err := DB.Self.Where("").Offset(offset).Limit(limit).Order("id desc").Find(&list).Error; err != nil {
		return list, count, err
	}
	return list, count, nil
}

// 根据标签id获取角色
func (role *RoleModel) GetRoleById(id uint64) (*RoleModel, error) {
	d := DB.Self.First(&role, id)
	return role, d.Error
}

// 根据标签id删除角色
func (role *RoleModel) DeleteRole(id uint64) error {
	role.BaseModel.Id = id
	return DB.Self.Delete(&role).Error
}

// 更新角色
func (role *RoleModel) Update() error {
	return DB.Self.Omit("created_at").Save(role).Error
}
