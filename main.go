package main

import "github.com/jcocozza/ny_taxi_pseudo_gen/api"

func main() {
	/*
	var cfg snowflake.SnowConfig
	cfg.Read()
	conn,err := snowflake.SnowflakeConn()
	if err != nil {
		panic(err)
	}
	defer conn.Close()
	*/
	api.Serve()
}
