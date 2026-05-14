package repository

import (
	"errors"
	"reservasi/internal/domain"

	"gorm.io/gorm"
)

type bookingRepository struct {
	db *gorm.DB
}

// NewBookingRepository membuat instance repository baru
func NewBookingRepository(db *gorm.DB) domain.BookingRepository {
	return &bookingRepository{db: db}
}

func (r *bookingRepository) CreateBooking(booking *domain.Booking) error {
	return r.db.Create(booking).Error
}

func (r *bookingRepository) CreateBookingAddon(addon *domain.BookingAddon) error {
	return r.db.Create(addon).Error
}

func (r *bookingRepository) GetBookingByID(id string) (*domain.Booking, error) {
	var booking domain.Booking
	// Preload Addons untuk menarik data layanan tambahan sekaligus
	err := r.db.Preload("Addons").Where("id = ?", id).First(&booking).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("booking not found")
		}
		return nil, err
	}
	return &booking, nil
}

func (r *bookingRepository) UpdateBooking(booking *domain.Booking) error {
	return r.db.Save(booking).Error
}

func (r *bookingRepository) GetRoomByID(id string) (*domain.Room, error) {
	var room domain.Room
	err := r.db.Where("id = ?", id).First(&room).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("room not found")
		}
		return nil, err
	}
	return &room, nil
}

func (r *bookingRepository) GetAddonByID(id string) (*domain.Addon, error) {
	var addon domain.Addon
	err := r.db.Where("id = ?", id).First(&addon).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("addon not found")
		}
		return nil, err
	}
	return &addon, nil
}

func (r *bookingRepository) GetGuestByID(id string) (*domain.Guest, error) {
	var guest domain.Guest
	err := r.db.Where("id = ?", id).First(&guest).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("guest not found")
		}
		return nil, err
	}
	return &guest, nil
}
