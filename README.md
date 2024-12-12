# RESTs-API-Penjadwalan-MatKul

## *Overview*
Ini adalah API RESTful sederhana yang dibuat menggunakan Go dan router Gorilla Mux. API ini memungkinkan pengguna untuk mengelola jadwal kuliah, termasuk membuat, membaca, memperbarui, dan menghapus jadwal.

# End Point
1. **GET /schedules**
Mengembalikan daftar semua jadwal dalam format JSON.
2. **GET /schedule/{id}**
Mengembalikan satu jadwal berdasarkan ID dalam format JSON.
Jika jadwal tidak ditemukan, mengembalikan kesalahan 404.
3. **POST /schedule**
**Membuat jadwal baru.**
_Expects a JSON dengan bidang-bidang berikut:_
* id: Pengidentifikasi unik untuk jadwal.
* course: Nama kursus.
* day: Hari dalam seminggu (misalnya, "Senin", "Selasa", dll.).
* time: Waktu kursus (misalnya, "08:00 - 10:00").
* location: Lokasi kursus (misalnya, "A101", "B202", dll.).
* Mengembalikan jadwal yang baru dibuat dalam format JSON.
4. *PUT /schedule/{id}*
 **Memperbarui jadwal yang ada.**
_Mengharapkan muatan JSON dengan bidang-bidang berikut:_
* id: Pengidentifikasi unik untuk jadwal. course: Nama kursus.
* day: Hari dalam seminggu (misalnya, "Senin", "Selasa", dll.).
* time: Waktu kursus (misalnya, "08:00 - 10:00").
* location: Lokasi kursus (misalnya, "A101", "B202", dll.).
* Mengembalikan jadwal yang diperbarui dalam format JSON.
* Jika jadwal tidak ditemukan, mengembalikan kesalahan 404.
5. *DELETE /schedule/{id}*
* Menghapus jadwal berdasarkan ID.
* Mengembalikan pesan sukses dalam format JSON.
* Jika jadwal tidak ditemukan, mengembalikan kesalahan 404.

# Create a new schedule
Send a POST request to /schedule with the following JSON payload
```*   The API will return the newly created schedule in JSON format.

### Update an existing schedule

*   Send a PUT request to `/schedule/1` with the following JSON payload:
    ```json
{
    "id": "1",
    "course": "Matematika",
    "day": "Senin",
    "time": "09:00 - 11:00",
    "location": "A101"
}
