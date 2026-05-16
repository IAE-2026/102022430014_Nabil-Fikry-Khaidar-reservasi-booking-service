package broker

import (
	"context"
	"encoding/json"
	"log"
	"os"
	"strings"

	"reservasi/internal/domain"
	"reservasi/internal/infrastructure"
)

type PaymentTimeoutMessage struct {
	BookingID string `json:"booking_id"`
}

// StartBrokerListener mulai listener untuk event payment timeout.
func StartBrokerListener(bookingUsecase domain.BookingUsecase) {
	channel := os.Getenv("BROKER_CHANNEL")
	if channel == "" {
		channel = "booking.payment.timeout"
	}

	if infrastructure.BrokerClient == nil {
		log.Println("Message broker tidak tersedia, listener dibatalkan.")
		return
	}

	ctx := context.Background()
	pubsub := infrastructure.BrokerClient.Subscribe(ctx, channel)
	_, err := pubsub.Receive(ctx)
	if err != nil {
		log.Printf("Gagal subscribe ke channel message broker %s: %v", channel, err)
		return
	}

	log.Printf("Mendengarkan event message broker pada channel: %s", channel)
	ch := pubsub.Channel()

	go func() {
		for msg := range ch {
			var event PaymentTimeoutMessage
			if err := json.Unmarshal([]byte(msg.Payload), &event); err != nil {
				log.Printf("Payload message broker tidak valid: %v", err)
				continue
			}

			if strings.TrimSpace(event.BookingID) == "" {
				log.Println("Event broker diterima tetapi booking_id kosong")
				continue
			}

			if err := bookingUsecase.HandleBookingPaymentTimeout(event.BookingID); err != nil {
				log.Printf("Gagal memproses event payment timeout untuk booking %s: %v", event.BookingID, err)
				continue
			}

			log.Printf("Booking %s direvert ke available karena payment timeout", event.BookingID)
		}
	}()
}
