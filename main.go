package main

import (
	//"net/http"
	"github.com/gin-gonic/gin"
	"github.com/gosimple/slug"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

type Article struct {
	gorm.Model
	Title string
	Slug  string `gorm:"unique_index"`
	Desc  string `sql:"type:text;"`
}

var DB *gorm.DB

// Router func
func main() {
	var err error
	DB, err = gorm.Open("mysql", "root:@tcp(127.0.0.1:3306)/learngin?charset=utf8mb4&parseTime=True&loc=Local")
	if err != nil {
		panic("failed to connect database")
	}
	defer DB.Close()

	// Migrate the schema
	DB.AutoMigrate(&Article{})

	router := gin.Default()
	//endopint
	v1 := router.Group("api/v1/")
	{
		articles := v1.Group("/article")
		articles.GET("/", getHome)
		articles.GET("/:slug", getArticle)
		articles.POST("/", postArticle)
	}
	router.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}

func getHome(c *gin.Context) {
	items := []Article{}
	// Get all records
	DB.Find(&items)
	//// SELECT * FROM users;
	c.JSON(200, gin.H{
		"status": "berhasil ke halaman home",
		"data":   items,
	})
}

func getArticle(c *gin.Context) {
	slug := c.Param("slug")

	var item Article

	if DB.First(&item, "slug = ?", slug).RecordNotFound() {
		c.JSON(404, gin.H{"status": "error", "message": "record not found"})
		c.Abort()
		return
	}
	// ambil data detail dari database/API
	// mengolah hasilnya
	c.JSON(200, gin.H{
		"status": "berhasil",
		"data":   item,
	})
}

func postArticle(c *gin.Context) {
	item := Article{
		Title: c.PostForm("title"),
		Desc:  c.PostForm("desc"),
		Slug:  slug.Make(c.PostForm("title")),
	}
	/* Tugas/PR ? kalau slug sama, maka generate random slug
	cara: ngecek database apakah sudah ada slug yg sama
	kalo sudah ada slug yg sama akan buat random slug: judul-pertama-randomslugblabla
	*/
	DB.Create(&item)
	// ambil data detail dari database/API
	// mengolah hasilnya
	c.JSON(200, gin.H{
		"status": "berhasil ngepost",
		"data":   item,
	})
}
