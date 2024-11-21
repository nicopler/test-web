package main

import (
	"net/http"
)

func main() {
    // Melayani file HTML
    http.HandleFunc("/healty", func(w http.ResponseWriter, r *http.Request) {
        http.ServeFile(w, r, "index.html")
    })

    // Melayani file CSS dari folder static
    http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("."))))


    // Menjalankan server di port 8080
    http.ListenAndServe(":8080", nil)
};