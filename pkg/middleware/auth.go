package middleware

import (
	"net/http"
	"os"
	"reservasi/internal/domain"

	"github.com/gin-gonic/gin"
)

// AuthMiddleware memvalidasi header X-IAE-KEY sebelum meneruskan request
func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Mengambil nilai header X-IAE-KEY dari request
		key := c.GetHeader("X-IAE-KEY")

		// Mengambil kunci valid dari environment, fallback ke nilai wajib 102022430014
		expectedKey := os.Getenv("IAE_KEY")
		if expectedKey == "" {
			expectedKey = "102022430014"
		}

		// Validasi key
		if key != expectedKey {
			// Membuat response error menggunakan wrapper global
			response := domain.ErrorResponse{
				Status:  "error",
				Message: "Unauthorized: Invalid or missing X-IAE-KEY header",
			}

			// Mengembalikan response HTTP 401 Unauthorized dan menghentikan eksekusi handler selanjutnya
			c.JSON(http.StatusUnauthorized, response)
			c.Abort()
			return
		}

		// Jika valid, teruskan ke handler berikutnya
		c.Next()
	}
}
