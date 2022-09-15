package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/gosimple/slug"
	"github.com/itsadityo/gin-full-api/config"
	"github.com/itsadityo/gin-full-api/models"
)

func GetHome(c *gin.Context) {
	items := []models.Article{}
	// Get all records
	config.DB.Find(&items)
	//// SELECT * FROM users;
	c.JSON(200, gin.H{
		"status": "berhasil ke halaman home",
		"data":   items,
	})
}

func GetArticle(c *gin.Context) {
	slug := c.Param("slug")

	var item models.Article

	if config.DB.First(&item, "slug = ?", slug).RecordNotFound() {
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

func PostArticle(c *gin.Context) {
	item := models.Article{
		Title: c.PostForm("title"),
		Desc:  c.PostForm("desc"),
		Slug:  slug.Make(c.PostForm("title")),
	}
	/* Tugas/PR ? kalau slug sama, maka generate random slug
	cara: ngecek database apakah sudah ada slug yg sama
	kalo sudah ada slug yg sama akan buat random slug: judul-pertama-randomslugblabla
	*/
	config.DB.Create(&item)
	// ambil data detail dari database/API
	// mengolah hasilnya
	c.JSON(200, gin.H{
		"status": "berhasil ngepost",
		"data":   item,
	})
}
