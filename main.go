package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

// Schedule represents the course schedule
type Schedule struct {
	ID       string `json:"id"`
	Course   string `json:"course"`
	Day      string `json:"day"`
	Time     string `json:"time"`
	Location string `json:"location"`
}

// Data awal jadwal mata kuliah
var schedules = []Schedule{
	{ID: "1", Course: "Matematika", Day: "Senin", Time: "08:00 - 10:00", Location: "A101"},
	{ID: "2", Course: "Fisika", Day: "Selasa", Time: "10:00 - 12:00", Location: "B202"},
	{ID: "3", Course: "Kimia", Day: "Rabu", Time: "13:00 - 15:00", Location: "C303"},
}

// Handler untuk mengambil semua jadwal
func getSchedules(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(schedules)
}

// Handler untuk mengambil jadwal berdasarkan ID
func getSchedule(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	for _, schedule := range schedules {
		if schedule.ID == params["id"] {
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(schedule)
			return
		}
	}
	http.Error(w, "Schedule not found", http.StatusNotFound)
}

// Handler untuk menambahkan jadwal baru
func createSchedule(w http.ResponseWriter, r *http.Request) {
	var newSchedule Schedule
	if err := json.NewDecoder(r.Body).Decode(&newSchedule); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	schedules = append(schedules, newSchedule)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(newSchedule)
}

// Handler untuk memperbarui jadwal berdasarkan ID
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
	http.Error(w, "Schedule not found", http.StatusNotFound)
}

// Handler untuk menghapus jadwal berdasarkan ID
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
	http.Error(w, "Schedule not found", http.StatusNotFound)
}

// Fungsi utama untuk menjalankan server
func main() {
	r := mux.NewRouter()

	// Definisikan route untuk API
	r.HandleFunc("/schedules", getSchedules).Methods("GET")
	r.HandleFunc("/schedule/{id}", getSchedule).Methods("GET")
	r.HandleFunc("/schedule", createSchedule).Methods("POST")
	r.HandleFunc("/schedule/{id}", updateSchedule).Methods("PUT")
	r.HandleFunc("/schedule/{id}", deleteSchedule).Methods("DELETE")

	// Mulai server di port 8000
	fmt.Println("Server is running on port 8000...")
	log.Fatal(http.ListenAndServe(":8000", r))
}
