package internal

import (
	"time"
)

type TaxiRecord struct {
	VendorId             int       `json:"vendor_id"`
	Pickup               time.Time `json:"pickup_datetime"`
	Dropoff              time.Time `json:"dropoff_datetime"`
	PassengerCount       int       `json:"passenger_count"`
	TripDistance         float64   `json:"trip_distance"`
	RateCodeId           int       `json:"ratecode_id"`
	StoreAndFwdFlag      string    `json:"store_and_fwd_flag"`
	PulocationId         int       `json:"pulocationid"`
	DolocationId         int       `json:"Dolocationid"`
	PaymentType          int       `json:"payment_type"`
	FareAmount           float64   `json:"fare_amount"`
	Extra                float64   `json:"extra"`
	MtaTax               float64   `json:"mta_tax"`
	TipAmount            float64   `json:"tip_amount"`
	TollsAmount          float64   `json:"tolls_amount"`
	ImprovementSurcharge float64   `json:"improvement_surcharge"`
	TotalAmount          float64   `json:"total_amount"`
	CongestionSurcharge  float64   `json:"congestion_surcharge"`
	AirportFee           float64   `json:"airport_fee"`
	TaxiType             string    `json:"taxi_type"`
	TripType             int       `json:"trip_type"`
}

func CreateNewTaxiRecord() TaxiRecord {
	return TaxiRecord{
		VendorId:             Taxi_DiscreteInt[vendorId].WeightedRandomSelection(),
		Pickup:               time.Now(),
		Dropoff:              time.Now().Add(time.Duration(float64(time.Minute) * (Taxi_Continuous[tripDuration].GenNormRand()))),
		PassengerCount:       Taxi_DiscreteInt[passengerCount].WeightedRandomSelection(),
		TripDistance:         Taxi_Continuous[tripDistance].GenNormRand(),
		RateCodeId:           Taxi_DiscreteInt[ratecodeid].WeightedRandomSelection(),
		StoreAndFwdFlag:      Taxi_DiscreteStr[storeAndFwdFlag].WeightedRandomSelection(),
		PulocationId:         Taxi_DiscreteInt[pulocationid].WeightedRandomSelection(),
		DolocationId:         Taxi_DiscreteInt[dolocationid].WeightedRandomSelection(),
		PaymentType:          Taxi_DiscreteInt[paymentType].WeightedRandomSelection(),
		FareAmount:           Taxi_Continuous[fareAmount].GenNormRand(),
		Extra:                Taxi_Continuous[extra].GenNormRand(),
		MtaTax:               Taxi_Continuous[mtaTax].GenNormRand(),
		TipAmount:            Taxi_Continuous[tipAmount].GenNormRand(),
		TollsAmount:          Taxi_Continuous[tollsAmount].GenNormRand(),
		ImprovementSurcharge: Taxi_Continuous[improvementSurcharge].GenNormRand(),
		TotalAmount:          Taxi_Continuous[totalAmount].GenNormRand(),
		CongestionSurcharge:  Taxi_Continuous[congestionSurcharge].GenNormRand(),
		AirportFee:           Taxi_Continuous[airportFee].GenNormRand(),
		TaxiType:             Taxi_DiscreteStr[taxiType].WeightedRandomSelection(),
		TripType:             Taxi_DiscreteInt[tripType].WeightedRandomSelection(),
	}
}
