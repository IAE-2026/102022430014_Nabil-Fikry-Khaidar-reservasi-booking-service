Layanan-Reservasi/
├── cmd/                      # Folder file eksekusi aplikasi
│   └── app/                  # Inisialisasi spesifik aplikasi
│       └── main.go           # Titik awal server dijalankan (Entry point)
├── internal/                 # Kode inti yang bersifat privat (tidak bisa di-import luar)
│   ├── domain/               # Definisi model data & interface (Solusi circular dependency)
│   ├── delivery/             # Layer penghubung antara user dan aplikasi (Pintu masuk)
│   │   ├── rest/             # Handler untuk REST API menggunakan Gin
│   │   └── graphql/          # Handler untuk API GraphQL (Pindahan dari folder graph/)
│   ├── usecase/              # Tempat logika bisnis dan aturan aplikasi utama
│   ├── repository/           # Implementasi akses data ke database (Postgres & Redis)
│   └── infrastructure/       # Pengaturan teknis koneksi database, cache, dan driver
├── pkg/                      # Helper/Library pendukung yang bisa digunakan project lain
│   └── middleware/           # Logika penyaring request (Auth, Security, Logging)
├── configs/                  # Tempat menyimpan file konfigurasi (.env, config.yaml)
├── migrations/               # File catatan sejarah perubahan struktur database SQL
├── docs/                     # Dokumentasi teknis API (biasanya file Swagger/OpenAPI)
├── Dockerfile                # Instruksi pembungkusan aplikasi ke dalam container
├── ai-log.md                 # Rekap jurnal proses pengembangan dan interaksi dengan AI
├── go.mod                    # File utama daftar library/dependency project
└── go.sum                    # Catatan verifikasi keamanan dan keaslian library