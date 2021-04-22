package controller

import (
	"time"
	"web/config"
	"web/models"

	"github.com/gin-gonic/gin"
)

func GetHome(c *gin.Context) {
	items := []models.User{}
	config.DB.Find(&items)

	c.JSON(200, gin.H{
		"status": "berhasil ke halaman home",
		"data":   items,
	})
}

func GetUserById(c *gin.Context) {
	id := c.Param("id")

	var item models.User

	if config.DB.First(&item, "id = ?", id).RecordNotFound() {
		c.JSON(404, gin.H{"status": "error", "message": "record not found"})
		c.Abort()
		return
	}

	c.JSON(200, gin.H{
		"status": "berhasil",
		"data":   item,
	})
}

func GetAtasan(c *gin.Context) {
	items := []models.User{}
	config.DB.Find(&items)

	c.JSON(200, gin.H{
		"status": "berhasil ke halaman home",
		"data":   items,
	})
}

func GetAtasanById(c *gin.Context) {
	id := c.Param("id")
	items := []models.User{}

	config.DB.Where("id <> ?", id).Find(&items)

	c.JSON(200, gin.H{
		"status": "berhasil",
		"data":   items,
	})
}

func PostUser(c *gin.Context) {

	var olditem models.User
	var layoutFormat, value string
	var date time.Time

	layoutFormat = "2006-01-02"
	value = c.PostForm("TanggalLahir")
	date, _ = time.Parse(layoutFormat, value)

	if !config.DB.First(&olditem, "nip = ?", c.PostForm("Nip")).RecordNotFound() {
		c.JSON(200, gin.H{"status": "error", "message": "Nip sudah ada"})
		// c.Abort()
		return
	}

	if !config.DB.First(&olditem, "nama_lengkap = ?", c.PostForm("NamaLengkap")).RecordNotFound() {
		c.JSON(200, gin.H{"status": "error", "message": "Nama sudah ada"})
		// c.Abort()
		return
	}

	item := models.User{
		Nip:          c.PostForm("Nip"),
		NamaLengkap:  c.PostForm("NamaLengkap"),
		TempatLahir:  c.PostForm("TempatLahir"),
		TanggalLahir: date,
		Atasan:       c.PostForm("Atasan"),
		Status:       c.PostForm("Status"),
	}

	config.DB.Create(&item)

	c.JSON(200, gin.H{
		"status": "berhasil ngepost",
		"data":   item,
	})
}

func UpdateUser(c *gin.Context) {
	id := c.Param("id")
	var layoutFormat, value string
	var date time.Time

	layoutFormat = "2006-01-02"
	value = c.PostForm("TanggalLahir")
	date, _ = time.Parse(layoutFormat, value)

	var item models.User

	if config.DB.First(&item, "id = ?", id).RecordNotFound() {
		c.JSON(404, gin.H{"status": "error", "message": "record not found"})
		c.Abort()
		return
	}

	config.DB.Model(&item).Where("id = ?", id).Updates(models.User{
		Nip:          c.PostForm("Nip"),
		NamaLengkap:  c.PostForm("NamaLengkap"),
		TempatLahir:  c.PostForm("TempatLahir"),
		TanggalLahir: date,
		Atasan:       c.PostForm("Atasan"),
		Status:       c.PostForm("Status"),
	})

	c.JSON(200, gin.H{
		"status": "berhasil update",
		"data":   item,
	})
}

func GetDetailUser(c *gin.Context) {
	id := c.Param("id")
	items := models.User{}
	var hasil []string

	config.DB.Where("id = ?", id).Find(&items)
	// hasil = append(hasil, items.NamaLengkap)

	result := getBawahan(items.NamaLengkap)
	for _, val := range result {
		hasil = append(hasil, val)
	}

	c.JSON(200, gin.H{
		"status":   "berhasil update",
		"datadiri": items,
		"data":     hasil,
	})
}

func getBawahan(atasan string) []string {
	items := []models.User{}
	var s []string
	config.DB.Where("atasan = ?", atasan).Find(&items)

	for _, val := range items {
		s = append(s, val.NamaLengkap)
		loop := getBawahan(val.NamaLengkap)
		if len(loop) > 0 {
			for _, v := range loop {
				s = append(s, v)
			}
		}
	}
	return s
}
