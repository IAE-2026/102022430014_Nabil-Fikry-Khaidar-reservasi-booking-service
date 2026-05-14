Kontrak API

1. Catalog Service (Pilih-Pilih Kamar)
GET /rooms: Menampilkan daftar kamar yang dibuat (berdasarkan lokasi & tanggal).
POST /rooms: Menyimpan kamar ke katalog 
GET /rooms/{id}: Membuka detail lengkap satu kamar (foto, fasilitas, deskripsi).
GET /addons: Menampilkan menu tambahan seperti sarapan atau asuransi.

2. Guest Service (Data Diri Tamu)
GET /{guestId}: Mengambil data profil secara otomatis agar tidak perlu ketik ulang nama/email.
POST /profile: Menyimpan atau memperbarui data identitas pemesan.
POST /validate-ktp: Mengecek data teman menginap (misal: memastikan nomor KTP teman valid).

3. Booking / Reservasi  Service (Tanpa Transaksi & Bayar)
POST /bookings: Membuat pesanan awal (mengunci kamar agar tidak diambil orang lain).
POST /{id}/addons: Menambahkan pilihan tambahan (makanan/asuransi) ke dalam tagihan pesananmu.
GET /{id}/summary: Menampilkan nota total (Harga Kamar + Layanan Tambahan) sebelum kamu bayar.

Catatan
- Semua Endpoint di atas harus menggunakan Header : X-IAE-KEY: 102022430014 agar valid. (Tidak menggunakan Authentication JWT, Bearer, atau Login).
- di project ini kita hanya fokus membuat service untuk API layanan reservasi booking ini saja , bukan membuat front-end atau tampilan UI.