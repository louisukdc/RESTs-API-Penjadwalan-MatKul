# 🖥️ RESTs-API-Penjadwalan-MatKul 🔥

## *Overview*
Ini adalah API RESTful sederhana yang dibuat menggunakan [Go](https://go.dev/doc/) dan router [Gorilla Mux](https://github.com/gorilla/mux.git). API ini memungkinkan pengguna untuk mengelola jadwal kuliah, termasuk membuat, membaca, memperbarui, dan menghapus jadwal.

## End Point
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
4. **PUT /schedule/{id}**
 **Memperbarui jadwal yang ada.**
_Mengharapkan muatan JSON dengan bidang-bidang berikut:_
* id: Pengidentifikasi unik untuk jadwal. course: Nama kursus.
* day: Hari dalam seminggu (misalnya, "Senin", "Selasa", dll.).
* time: Waktu kursus (misalnya, "08:00 - 10:00").
* location: Lokasi kursus (misalnya, "A101", "B202", dll.).
* Mengembalikan jadwal yang diperbarui dalam format JSON.
* Jika jadwal tidak ditemukan, mengembalikan kesalahan 404.
5. **DELETE /schedule/{id}**
* Menghapus jadwal berdasarkan ID.
* Mengembalikan pesan sukses dalam format JSON.
* Jika jadwal tidak ditemukan, mengembalikan kesalahan 404.

# Create a new schedule
Menggunakan method **PUT** pada JSON
```*   The API will return the newly created schedule in JSON format.

### Update an existing schedule

    ```json
{
    "id": "1",
    "course": "Matematika",
    "day": "Senin",
    "time": "09:00 - 11:00",
    "location": "A101"
}
```

* The API will return the updated schedule in JSON format.

# Delete a schedule
* Send a DELETE request to /schedule/2.
* The API will return a success message in JSON format.
!good day
