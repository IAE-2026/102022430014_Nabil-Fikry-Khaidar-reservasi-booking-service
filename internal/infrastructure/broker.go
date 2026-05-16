package infrastructure

import (
	"context"
	"log"
	"os"
	"strings"

	"github.com/redis/go-redis/v9"
)

// BrokerClient adalah instance global dari message broker client.
// Saat ini menggunakan Redis Pub/Sub sebagai placeholder untuk message broker.
var BrokerClient *redis.Client

// ConnectBroker menginisialisasi koneksi message broker.
// Jika BROKER_URL atau BROKER_ADDR tidak diset, broker akan dinonaktifkan.
func ConnectBroker() {
	brokerURL := os.Getenv("BROKER_URL")
	if brokerURL == "" {
		brokerURL = os.Getenv("BROKER_ADDR")
	}
	if brokerURL == "" {
		log.Println("Warning: BROKER_URL atau BROKER_ADDR tidak dikonfigurasi, message broker dinonaktifkan.")
		return
	}

	var opts *redis.Options
	var err error
	if strings.HasPrefix(brokerURL, "redis://") || strings.HasPrefix(brokerURL, "rediss://") {
		opts, err = redis.ParseURL(brokerURL)
		if err != nil {
			log.Printf("Warning: gagal parse BROKER_URL sebagai Redis URL: %v", err)
		}
	}

	if opts == nil {
		opts = &redis.Options{Addr: brokerURL, Password: "", DB: 0}
	}

	BrokerClient = redis.NewClient(opts)

	ctx := context.Background()
	if err := BrokerClient.Ping(ctx).Err(); err != nil {
		log.Printf("Warning: gagal terhubung ke message broker di %s: %v", brokerURL, err)
		BrokerClient = nil
		return
	}

	log.Println("Koneksi ke message broker berhasil!")
}
