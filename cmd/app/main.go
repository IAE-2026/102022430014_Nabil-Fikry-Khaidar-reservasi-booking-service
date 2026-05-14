package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

// Struktur data reservasi
type Reservasi struct {
	Nama   string `json:"nama" binding:"required"`
	Tujuan string `json:"tujuan" binding:"required"`
	Jumlah int    `json:"jumlah"`
}

func main() {
	r := gin.Default()

	// Endpoint GET (yang sudah Tuan buat)
	r.GET("/status", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"status": "online"})
	})

	// Endpoint POST untuk menerima reservasi baru
	r.POST("/pesan", func(c *gin.Context) {
		var input Reservasi

		// Validasi apakah data JSON yang dikirim sesuai dengan struct
		if err := c.ShouldBindJSON(&input); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		// Response sukses
		c.JSON(http.StatusOK, gin.H{
			"pesan": "Reservasi berhasil dibuat!",
			"data":  input,
		})
	})

	r.Run()
}