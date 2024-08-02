package api

import (
    "fmt"
	"encoding/json"
	"net/http"

	"github.com/jcocozza/ny_taxi_pseudo_gen/internal"
)

const (
    port = ":8088"
)

func requestTaxiHandler(w http.ResponseWriter, r *http.Request) {
    taxi := internal.CreateNewTaxiRecord()

    w.Header().Set("Content-Type", "application/json")
    err := json.NewEncoder(w).Encode(taxi)
    if err != nil {
        panic(err)
    }
}


func Serve() {
    http.HandleFunc("/request_taxi", requestTaxiHandler)
    fmt.Printf("Server is running on http://localhost%s\n", port)

    err := http.ListenAndServe(port, nil)
    if err != nil {
        panic(err)
    }
}
