package models

type Post struct {
	Model
	Title   string `sql:"not null" json:"title" binding:"required"`
	Desc    string `sql:"not null" json:"desc" binding:"required"`
	Content string `sql:"not null;type:text" json:"content" binding:"required"`
	UserId  int    `json:"postId,omitempty"`
}

func ExistPostByID(id int) bool {
	var post Post
	db.Select("id").Where("id = ?", id).First(&post)
	if post.ID > 0 {
		return true
	}

	return false
}

func ExistPostByEmail(email string) bool {
	var post Post
	db.Select("id").Where("email = ?", email).First(&post)
	if post.ID > 0 {
		return true
	}

	return false
}

func GetPostTotal(maps interface{}) (count int) {
	db.Model(&Post{}).Where(maps).Count(&count)

	return
}

func GetPosts() (posts []Post) {
	db.Order("updated_at desc").Find(&posts)

	return
}

func GetPost(id int) (post Post) {
	db.Where("id = ?", id).First(&post)

	return
}

func AddPost(post Post) bool {
	db.Create(&post)

	return true
}

func EditPost(id int, data interface{}) bool {
	db.Model(&Post{}).Where("id = ?", id).Updates(data)

	return true
}

func DeletePost(id int) bool {
	db.Where("id = ?", id).Delete(Post{})

	return true
}
