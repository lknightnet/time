package main

import (
	"encoding/json"
	"log"
	"net/http"
	"time"
)

func main() {
	log.Println("app is starting . . .")
	mux := http.NewServeMux()

	mux.HandleFunc("/time", func(w http.ResponseWriter, r *http.Request) {
		date := time.Now().Format("02.01.2006")
		timeT := time.Now().Format("15:04:05")
		m := map[string]interface{}{
			"date": date,
			"time": timeT,
		}

		data, err := json.Marshal(m)
		if err != nil {
			log.Println(err)
		}
		_, err = w.Write(data)
		if err != nil {
			return
		}
	})

	err := http.ListenAndServe(":8000", mux)
	if err != nil {
		log.Println(err)
	}
}
