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
Expects a JSON JSON dengan bidang-bidang berikut:
* id: Pengidentifikasi unik untuk jadwal.
* course: Nama kursus.
* day: Hari dalam seminggu (misalnya, "Senin", "Selasa", dll.).
* time: Waktu kursus (misalnya, "08:00 - 10:00").
* location: Lokasi kursus (misalnya, "A101", "B202", dll.).
* Mengembalikan jadwal yang baru dibuat dalam format JSON.
4. *PUT /schedule/{id}*
* **Memperbarui jadwal yang ada.**
Mengharapkan muatan JSON dengan bidang-bidang berikut:
* id: Pengidentifikasi unik untuk jadwal. course: Nama kursus.
* day: Hari dalam seminggu (misalnya, "Senin", "Selasa", dll.).
* time: Waktu kursus (misalnya, "08:00 - 10:00").
* location: Lokasi kursus (misalnya, "A101", "B202", dll.).
* Mengembalikan jadwal yang diperbarui dalam format JSON.
Jika jadwal tidak ditemukan, mengembalikan kesalahan 404.
DELETE /schedule/{id}
Menghapus jadwal berdasarkan ID.
Mengembalikan pesan sukses dalam format JSON.
Jika jadwal tidak ditemukan, mengembalikan kesalahan 404.


## Create New Schedule

This section describes the API endpoint for creating a new schedule entry.

**Endpoint:** (Replace with your actual endpoint URL)

**HTTP Method:** POST

**Expected Payload (JSON):**

| Field        | Description                                 | Data Type           | Required |
|--------------|----------------------------------------------|--------------------|----------|
| id            | Unique identifier for the schedule.       | String              | Yes      |
| course        | Name of the course.                          | String              | Yes      |
| day           | Day of the week (e.g., "Senin", "Selasa").    | String              | Yes      |
| time          | Time of the course (e.g., "08:00 - 10:00").   | String              | Yes      |
| location      | Location of the course (e.g., "A101", "B202"). | String              | Yes      |

**Example Payload:**

```json
{
  "id": "schedule-1",
  "course": "Introduction to Programming",
  "day": "Senin",
  "time": "08:00 - 10:00",
  "location": "A101"
}
