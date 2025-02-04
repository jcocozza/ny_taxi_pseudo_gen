package internal

import (
	"context"
	"time"

	"github.com/jcocozza/ny_taxi_pseudo_gen/internal/snowflake"
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

func CreateTaxiRecordFromLocs(puloc, doloc int) TaxiRecord {
	return TaxiRecord{
		VendorId:             Taxi_DiscreteInt[vendorId].WeightedRandomSelection(),
		Pickup:               time.Now(),
		Dropoff:              time.Now().Add(time.Duration(float64(time.Minute) * (Taxi_Continuous[tripDuration].GenNormRand()))),
		PassengerCount:       Taxi_DiscreteInt[passengerCount].WeightedRandomSelection(),
		TripDistance:         Taxi_Continuous[tripDistance].GenNormRand(),
		RateCodeId:           Taxi_DiscreteInt[ratecodeid].WeightedRandomSelection(),
		StoreAndFwdFlag:      Taxi_DiscreteStr[storeAndFwdFlag].WeightedRandomSelection(),
		PulocationId:         puloc,
		DolocationId:         doloc,
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

func WriteToSnowflake(taxiRecord TaxiRecord) error {
	db, err := snowflake.SnowflakeConn()
	if err != nil {
		return err
	}
	defer db.Close()
	sql := "INSERT INTO real_time_test (pulocationid, dolocationid, total_amount) VALUES (?, ?, ?)"
	return snowflake.RunSQL(db, sql, taxiRecord.PulocationId, taxiRecord.DolocationId, taxiRecord.TotalAmount)
}

func GetPricingModifier(taxiRecord TaxiRecord) (float64, error) {
	db, err := snowflake.SnowflakeConn()
	if err != nil {
		return -1, err
	}
	defer db.Close()

	//sql := "SELECT MAX(modifier) FROM pricing_modifier_by_zone WHERE location_id IN (?, ?);"

	sql := `WITH cte AS (
SELECT borough as b
FROM dim_zone_lookup
WHERE locationid IN (?,?)
)
SELECT MAX(percent_increase)
FROM borough_hr_pricing
INNER JOIN cte ON UPPER(borough_hr_pricing.borough) = UPPER(cte.b)
WHERE borough_hr_pricing.hour = ?`

	row := db.QueryRowContext(context.TODO(), sql, taxiRecord.PulocationId, taxiRecord.DolocationId, taxiRecord.Pickup.Hour())

	var modifier float64
	err = row.Scan(&modifier)
	if err != nil {
		return -1, err
	}
	return modifier, nil
}
