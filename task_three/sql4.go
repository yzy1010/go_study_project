package main

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"time"
)

// User 用户模型
type User struct {
	ID        uint           `gorm:"primaryKey"`
	Username  string         `gorm:"size:50;not null;unique"`
	Email     string         `gorm:"size:100;not null;unique"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`

	// 关联字段
	Posts     []Post         `gorm:"foreignKey:UserID"`
	PostCount int            `gorm:"default:0"` // 文章数量统计
}

// Post 文章模型
type Post struct {
	ID          uint           `gorm:"primaryKey"`
	Title       string         `gorm:"size:200;not null"`
	Content     string         `gorm:"type:text"`
	UserID      uint           `gorm:"not null"`
	CommentStatus string       `gorm:"size:20;default:'有评论'"`
	CreatedAt   time.Time
	UpdatedAt   time.Time
	DeletedAt   gorm.DeletedAt `gorm:"index"`

	// 关联字段
	User      User           `gorm:"foreignKey:UserID"`
	Comments  []Comment      `gorm:"foreignKey:PostID"`
}

// Comment 评论模型
type Comment struct {
	ID        uint           `gorm:"primaryKey"`
	Content   string         `gorm:"type:text;not null"`
	PostID    uint           `gorm:"not null"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`

	// 关联字段
	Post      Post           `gorm:"foreignKey:PostID"`
}

// 在Post模型中添加钩子函数
func (p *Post) AfterCreate(tx *gorm.DB) error {
	// 文章创建后自动更新用户的文章数量
	return tx.Model(&User{}).Where("id = ?", p.UserID).UpdateColumn("post_count",
		gorm.Expr("post_count + ?", 1)).Error
}

// 在Comment模型中添加钩子函数
func (c *Comment) AfterDelete(tx *gorm.DB) error {
	// 获取当前文章的评论数量
	var count int64
	tx.Model(&Comment{}).Where("post_id = ?", c.PostID).Count(&count)

	if count == 0 {
		// 如果评论数量为0，更新文章的评论状态
		return tx.Model(&Post{}).Where("id = ?", c.PostID).Update("comment_status", "无评论").Error
	}

	return nil
}

// 测试钩子函数
func testHooks(db *gorm.DB) {
	// 创建用户
	user := User{Username: "testuser", Email: "test@example.com"}
	db.Create(&user)

	// 创建文章（会触发AfterCreate钩子）
	post := Post{Title: "测试文章", Content: "这是测试内容", UserID: user.ID}
	db.Create(&post)

	// 检查用户文章数量是否更新
	db.First(&user, user.ID)
	fmt.Printf("用户文章数量：%d\n", user.PostCount)

	// 创建评论
	comment := Comment{Content: "测试评论", PostID: post.ID}
	db.Create(&comment)

	// 删除评论（会触发AfterDelete钩子）
	db.Delete(&comment)

	// 检查文章评论状态是否更新
	db.First(&post, post.ID)
	fmt.Printf("文章评论状态：%s\n", post.CommentStatus)
}

// 查询某个用户发布的所有文章及其对应的评论信息
func GetUserWithPostsAndComments(db *gorm.DB, userID uint) (User, error) {
	var user User

	err := db.Preload("Posts").                   // 预加载文章
		   Preload("Posts.Comments").            // 预加载文章的评论
		   First(&user, userID).Error

	return user, err
}

// 查询评论数量最多的文章信息
func GetPostWithMostComments(db *gorm.DB) (Post, error) {
	var post Post

	err := db.Select("posts.*, COUNT(comments.id) as comment_count").
		   Joins("LEFT JOIN comments ON posts.id = comments.post_id").
		   Group("posts.id").
		   Order("comment_count DESC").
		   First(&post).Error

	return post, err
}

// 在main函数中添加测试代码
func testQueries(db *gorm.DB) {
	// 测试查询用户及其文章和评论
	user, err := GetUserWithPostsAndComments(db, 1)
	if err == nil {
		fmt.Printf("用户：%s\n", user.Username)
		for _, post := range user.Posts {
			fmt.Printf("  文章：%s (评论数：%d)\n", post.Title, len(post.Comments))
		}
	}

	// 测试查询评论最多的文章
	post, err := GetPostWithMostComments(db)
	if err == nil {
		fmt.Printf("评论最多的文章：%s\n", post.Title)
	}
}

func main() {
	// 数据库连接
	dsn := "root:password@tcp(127.0.0.1:3306)/blog_system?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("连接数据库失败：" + err.Error())
	}

	// 自动迁移创建表
	err = db.AutoMigrate(&User{}, &Post{}, &Comment{})
	if err != nil {
		panic("创建表失败：" + err.Error())
	}

	fmt.Println("数据库表创建成功！")
}

