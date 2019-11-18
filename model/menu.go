package model

import (
	"github.com/jinzhu/gorm"
)

type MenuModel struct {
	BaseModel
	Name string `json:"name" gorm:"column:name;not null" binding:"required"`
	Path string `json:"path" gorm:"column:path;not null" binding:"required"`
	Method string `json:"method" gorm:"column:method;not null" binding:"required"`
}


func (m *MenuModel) TableName() string {
	return "menu"
}

func (m *MenuModel)ExistMenuByID(id int) (bool, error) {
	err := DB.Self.Where("id = ? AND deleted_on = ? ", id, 0).First(&m).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return false, err
	}
	if m.Id > 0 {
		return true, nil
	}
	return false, nil
}

func (m *MenuModel)GetMenuTotal(maps interface{}) (uint64, error) {
	var count uint64 = 0
	if err := DB.Self.Model(&m).Count(&count).Error; err != nil {
		return count, err
	}
	return count, nil
}

func (m *MenuModel)GetMenus(offset int, limit int, maps interface{}) ([]*MenuModel, error) {
	var menu []*MenuModel
	if err := DB.Self.Where("").Offset(offset).Limit(limit).Order("id desc").Find(&menu).Error; err != nil {
		return menu, err
	}
	return menu, nil
}

//func GetMenu(id int) (*Menu, error) {
//	var menu Menu
//	err := db.Where("id = ? AND deleted_on = ? ", id, 0).First(&menu).Error
//	if err != nil && err != gorm.ErrRecordNotFound {
//		return nil, err
//	}
//
//	return &menu, nil
//}
//
//func EditMenu(id int, data interface{}) error {
//	if err := db.Model(&Menu{}).Where("id = ? AND deleted_on = ? ", id, 0).Updates(data).Error; err != nil {
//		return err
//	}
//	return nil
//}
//
//func AddMenu(data map[string]interface{}) error {
//	menu := Menu{
//		Name:   data["name"].(string),
//		Path:   data["path"].(string),
//		Method: data["method"].(string),
//	}
//	if err := db.Create(&menu).Error; err != nil {
//		return err
//	}
//
//	return nil
//}
//
//func DeleteMenu(id int) error {
//	if err := db.Where("id = ?", id).Delete(Menu{}).Error; err != nil {
//		return err
//	}
//
//	return nil
//}
//
//func CleanAllMenu() error {
//	if err := db.Unscoped().Where("deleted_on != ? ", 0).Delete(&Menu{}).Error; err != nil {
//		return err
//	}
//	return nil
//}
//
//func EditMenuGetRoles(id int) []int {
//	var menu Menu
//	var role []Role
//
//	db.Model(&menu).Where("id = ? AND deleted_on = ? ", id, 0)
//	db.Joins(" left join go_role_menu b on go_role.id=b.role_id left join go_menu c on c.id=b.menu_id").Where("c.id = ?", id).Find(&role)
//
//	roleList := []int{}
//	for _, v := range role {
//		roleList = append(roleList, v.ID)
//	}
//	return roleList
//}
