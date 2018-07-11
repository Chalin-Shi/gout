package models

type User struct {
	Model
	Email    string `sql:"not null" json:"email"`
	Company  string `sql:"not null" json:"company"`
	Position string `sql:"not null" json:"position"`
	Username string `sql:"not null" json:"username"`
	Mobile   string `sql:"not null" json:"mobile"`
	Status   int    `sql:"default: 1" json:"status"`
}

func ExistUserByID(id int) bool {
	var user User
	db.Select("id").Where("id = ?", id).First(&user)
	if user.ID > 0 {
		return true
	}

	return false
}

func ExistUserByEmail(email string) bool {
	var user User
	db.Select("id").Where("email = ?", email).First(&user)
	if user.ID > 0 {
		return true
	}

	return false
}

func CheckUser(email string, password string) int {
	var user User
	db.Select("id").Where("email = ? AND password = ? ", email, password).First(&user)
	if user.ID > 0 {
		return user.ID
	}

	return 0
}

func GetUserTotal(maps interface{}) (count int) {
	db.Model(&User{}).Where(maps).Count(&count)

	return
}

func GetUsers(limit int, offset int, maps map[string]interface{}) (users []User) {
	db.Where(maps).Order("updated_at desc").Limit(limit).Offset(offset).Find(&users)

	return
}

func GetUser(id int) (user User) {
	db.Where("id = ?", id).First(&user)

	return
}

func AddUser(user User) bool {
	db.Create(&user)

	return true
}

func EditUser(id int, data interface{}) bool {
	db.Model(&User{}).Where("id = ?", id).Updates(data)

	return true
}

func DeleteUser(id int) bool {
	db.Where("id = ?", id).Delete(User{})

	return true
}
