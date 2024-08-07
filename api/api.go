package api

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/jcocozza/ny_taxi_pseudo_gen/internal"
)

const (
	port = ":8088"
)

func generateTaxiRecord(w http.ResponseWriter, r *http.Request) {
	taxi := internal.CreateNewTaxiRecord()

	w.Header().Set("Content-Type", "application/json")
	err := json.NewEncoder(w).Encode(taxi)
	if err != nil {
		panic(err)
	}
}

func servePage(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	w.Write([]byte(htmlContent))
}

func handleFormSubmission(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		fromStr := r.FormValue("from")
		toStr := r.FormValue("to")

		from, _ := strconv.Atoi(fromStr)
		to, _ := strconv.Atoi(toStr)
		taxi := internal.CreateTaxiRecordFromLocs(from, to)
		fmt.Println("taxi price original: ", taxi.TotalAmount)

		// very simple pricing logic modifier
		priceModifier, err := internal.GetPricingModifier(taxi)
		if err != nil {
			panic(err)
		}
		taxi.TotalAmount += priceModifier
		fmt.Println("taxi with updated pricing:", taxi.TotalAmount)

		w.Header().Set("Content-Type", "text/html")
		w.Write([]byte(fmt.Sprintf(`
            <html>
            <body>
				<p>Based on your zone and the current time: %s, your total price is: %.2f</p>
                <p>You will be redirected in 8 seconds...</p>
                <meta http-equiv="refresh" content="8; url=/request_taxi" />
            </body>
            </html>
        `, taxi.Pickup.String(), taxi.TotalAmount)))

		err = internal.WriteToSnowflake(taxi)
		if err != nil {
			panic(err)
		}

		/*
			w.Header().Set("Content-Type", "application/json")
			err = json.NewEncoder(w).Encode(taxi)
			if err != nil {
				panic(err)
			}
		*/
	} else {
		http.Redirect(w, r, "/", http.StatusSeeOther)
	}
}

func Serve() {
	http.HandleFunc("/request_taxi", servePage)
	http.HandleFunc("/submit", handleFormSubmission)
	http.HandleFunc("/generate_taxi_record", generateTaxiRecord)
	fmt.Printf("Server is running on http://localhost%s\n", port)

	err := http.ListenAndServe(port, nil)
	if err != nil {
		panic(err)
	}
}
