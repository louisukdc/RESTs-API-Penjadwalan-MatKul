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


## Function GET *Guna menampilkan semua KEY dan VALUE yang ada*

```
func getSchedule(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	for _, schedule := range schedules {
		if schedule.ID == params["id"] {
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(schedule)
			return
		}
	}
	http.Error(w, "Jadwal tidak ditemukan", http.StatusNotFound)
}
```

## Function Create *Guna menambahkan KEY dan VALUE yang baru*

```
func createSchedule(w http.ResponseWriter, r *http.Request) {
	var newSchedule Schedule
	if err := json.NewDecoder(r.Body).Decode(&newSchedule); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	schedules = append(schedules, newSchedule)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(newSchedule)
```

## Function Update *Guna mengubah Value yang sudah ada*

```
func updateSchedule(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	var updatedSchedule Schedule
	if err := json.NewDecoder(r.Body).Decode(&updatedSchedule); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	for index, schedule := range schedules {
		if schedule.ID == params["id"] {
			schedules[index] = updatedSchedule
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(updatedSchedule)
			return
		}
	}
	http.Error(w, "Jadwal tidak ditemukan", http.StatusNotFound)
}
```
## Function Delete *Guna  menghapus VALUE ddan KEY yang ssuddah ada*

```
func deleteSchedule(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	for index, schedule := range schedules {
		if schedule.ID == params["id"] {
			schedules = append(schedules[:index], schedules[index+1:]...)
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode("Schedule deleted successfully")
			return
		}
	}
	http.Error(w, "Jadwal tidak ditemukan", http.StatusNotFound)
}
```


# Create a new schedule
Menggunakan method **PUT** pada JSON
```*   The API will return the newly created schedule in JSON format.

### Update an existing schedule

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
