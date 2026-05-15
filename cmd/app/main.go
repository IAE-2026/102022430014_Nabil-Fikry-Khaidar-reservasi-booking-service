// @title Layanan Reservasi API
// @version 1.0
// @description API untuk layanan reservasi booking hotel (IAE Tubes)
// @host localhost:8080
// @BasePath /
package main

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"

	"reservasi/internal/delivery/rest"
	"reservasi/internal/domain"
	"reservasi/internal/infrastructure"
	"reservasi/internal/repository"
	"reservasi/internal/usecase"
	"reservasi/pkg/middleware"

	_ "reservasi/docs"

	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func main() {
	// 1. Muat environment variable dari .env
	if err := godotenv.Load("configs/.env"); err != nil {
		log.Println("Warning: File configs/.env tidak ditemukan, menggunakan variabel environment sistem.")
	}

	// 2. Inisialisasi koneksi Database & Redis
	infrastructure.ConnectPostgres()
	infrastructure.ConnectRedis()

	// 3. Auto Migrate skema database (Agar tabel terbentuk sesuai relasinya di PostgreSQL)
	// Kita gunakan GORM AutoMigrate yang akan menyamakan dengan struct domain
	err := infrastructure.DB.AutoMigrate(
		&domain.Guest{},
		&domain.Room{},
		&domain.Addon{},
		&domain.Booking{},
		&domain.BookingAddon{},
	)
	if err != nil {
		log.Fatalf("Gagal melakukan migrasi database: %v", err)
	}
	log.Println("Migrasi tabel database berhasil!")

	// 4. Inisialisasi Dependency Injection (Clean Architecture)
	bookingRepo := repository.NewBookingRepository(infrastructure.DB)
	bookingUsecase := usecase.NewBookingUsecase(bookingRepo)

	// 5. Inisialisasi Router Gin
	r := gin.Default()

	// 6. Swagger UI (Daftarkan SEBELUM middleware agar bisa diakses publik)
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	
	// Daftarkan middleware autentikasi secara global (berlaku untuk semua rute di bawahnya)
	r.Use(middleware.AuthMiddleware())

	// Endpoint dasar
	r.GET("/status", func(c *gin.Context) {
		res := domain.SuccessResponse{
			Status:  "success",
			Message: "Layanan Reservasi API Online",
		}
		c.JSON(http.StatusOK, res)
	})

	// 7. Daftarkan Handler Booking Service (REST Handlers)
	rest.NewBookingHandler(r, bookingUsecase)

	// 8. Jalankan Server
	log.Println("Server berjalan di port 8080...")
	r.Run(":8080")
}