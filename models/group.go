package models

type Group struct {
	Model
	Users []User `json:"users,omitempty"`

	Name string `sql:"not null" json:"name" binding:"required"`
	Desc string `sql:"type:text;not null" json:"desc" binding:"required"`
}

func ExistGroupByID(id int) bool {
	var group Group
	db.Select("id").Where("id = ?", id).First(&group)
	if group.ID > 0 {
		return true
	}
	return false
}

func ExistGroupByName(name string) bool {
	var group Group
	db.Select("id").Where("name = ?", name).First(&group)
	if group.ID > 0 {
		return true
	}
	return false
}

func GetGroupIdByName(name string) int {
	var group Group
	db.Select("id").Where("name = ?", name).First(&group)
	if group.ID > 0 {
		return group.ID
	}
	return 0
}

func GetGroupTotal(maps interface{}) (count int) {
	db.Model(&Group{}).Where(maps).Count(&count)

	return
}

func GetGroups(limit int, offset int, maps map[string]interface{}) (groups []Group) {
	db.Where(maps).Order("updated_at desc").Limit(limit).Offset(offset).Find(&groups)

	return
}

func GetGroup(id int) (group Group) {
	db.Where("id = ?", id).First(&group)

	return
}

func EditGroup(id int, data interface{}) bool {
	db.Model(&Group{}).Where("id = ?", id).Updates(data)

	return true
}

func EditGroupByAttr(id int, name string, data interface{}) bool {
	db.Model(&Group{}).Where("id = ?", id).Update(name, data)

	return true
}

func AddGroup(group Group) bool {
	db.Create(&group)

	return true
}

func DeleteGroup(id int) bool {
	db.Where("id = ?", id).Delete(Group{})

	return true
}

func GetGroupUsers(id int) (group Group) {
	db.Where("id = ?", id).First(&group)
	db.Model(&group).Related(&group.Users)

	return
}

func GetGroupUser(id int, username string) (group Group) {
	db.Where("id = ?", id).Preload("Users", "group_id = ? and username = ?", id, username).First(&group)

	return
}
