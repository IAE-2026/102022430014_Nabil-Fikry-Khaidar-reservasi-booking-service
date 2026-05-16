package usecase

import "errors"

func (u *bookingUsecase) HandleBookingPaymentTimeout(bookingID string) error {
	booking, err := u.bookingRepo.GetBookingByID(bookingID)
	if err != nil {
		return err
	}

	if booking.Status != "LOCKED" {
		return nil // Nothing to revert if booking already confirmed or cancelled
	}

	if err := u.bookingRepo.UpdateBookingStatus(bookingID, "CANCELLED"); err != nil {
		return errors.New("gagal memperbarui status pesanan")
	}

	if err := u.bookingRepo.UpdateRoomStatus(booking.RoomID.String(), "AVAILABLE"); err != nil {
		return errors.New("gagal mengembalikan status kamar ke available")
	}

	return nil
}
