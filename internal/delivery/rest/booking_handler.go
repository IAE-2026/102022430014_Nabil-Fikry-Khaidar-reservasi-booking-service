package rest

import (
	"net/http"
	"reservasi/internal/domain"

	"github.com/gin-gonic/gin"
)

type BookingHandler struct {
	bookingUsecase domain.BookingUsecase
}

// NewBookingHandler mendaftarkan endpoint untuk Booking Service
func NewBookingHandler(r *gin.Engine, us domain.BookingUsecase) {
	handler := &BookingHandler{
		bookingUsecase: us,
	}

	// Semua route terkait booking
	// Asumsinya global middleware sudah di-set di main.go untuk API key validation
	r.POST("/bookings", handler.CreateBooking)
	r.POST("/bookings/:id/addons", handler.AddAddon)
	r.GET("/bookings/:id/summary", handler.GetSummary)
}

func (h *BookingHandler) CreateBooking(c *gin.Context) {
	var req domain.CreateBookingRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, domain.ErrorResponse{
			Status:  "error",
			Message: "Payload tidak valid: " + err.Error(),
		})
		return
	}

	booking, err := h.bookingUsecase.CreateBooking(&req)
	if err != nil {
		// Menggunakan StatusInternalServerError atau BadRequest tergantung err
		c.JSON(http.StatusBadRequest, domain.ErrorResponse{
			Status:  "error",
			Message: err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, domain.SuccessResponse{
		Status:  "success",
		Message: "Pesanan awal berhasil dibuat (Kamar dikunci)",
		Data:    booking,
	})
}

func (h *BookingHandler) AddAddon(c *gin.Context) {
	bookingID := c.Param("id")

	var req domain.CreateBookingAddonRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, domain.ErrorResponse{
			Status:  "error",
			Message: "Payload tidak valid: " + err.Error(),
		})
		return
	}

	addon, err := h.bookingUsecase.AddAddon(bookingID, &req)
	if err != nil {
		c.JSON(http.StatusBadRequest, domain.ErrorResponse{
			Status:  "error",
			Message: err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, domain.SuccessResponse{
		Status:  "success",
		Message: "Layanan tambahan berhasil dimasukkan ke tagihan",
		Data:    addon,
	})
}

func (h *BookingHandler) GetSummary(c *gin.Context) {
	bookingID := c.Param("id")

	summary, err := h.bookingUsecase.GetSummary(bookingID)
	if err != nil {
		c.JSON(http.StatusNotFound, domain.ErrorResponse{
			Status:  "error",
			Message: err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, domain.SuccessResponse{
		Status:  "success",
		Message: "Nota total berhasil diambil",
		Data:    summary,
	})
}
