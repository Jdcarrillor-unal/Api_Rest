package main

import (
	"encoding/json"
	"log"
	"net/http"

	"APi_Rest/models"

	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()
	models.CreateConnection()
	defer models.CloseConnection()
	models.Ping()
	r.HandleFunc("/temp/{id}", GetDataforId).Methods("GET")
	log.Println("EL servido se encuentra en el puerto 4040")
	log.Fatal(http.ListenAndServe(":4040", r))

}
func GetDataforId(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	data_id := vars["id"]
	data := models.GetTempById(data_id)

	json.NewEncoder(w).Encode(data)

}
