# AI Log

Rekap jurnal proses pengembangan dan interaksi dengan AI.

[15/5/2026 - 00.18] 
Buatkan wrapper response pada folder internal/domain/ untuk RestAPI agar response sistem konsisten.

[15/5/2026 - 00.36] 
buat konfigurasi untuk koneksi PostgreSQL ( buat di internal/infrastructure/postgres.go). kemudian gunakan grom untuk mempermudah pemetaan model data ke tabel PostgreSQL. (untuk grom sudah saya instal di project ini)

[15/5/2026 - 00.40]
buat konfigurasi untuk koneksi Redis ( buat di internal/infrastructure/redis.go). gunakan go-redis untuk berinteraksi dengan layanan Redis lokal (sudah saya instal)