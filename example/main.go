package main

import (
	"os"
	"encoding/json"

	"github.com/savaki/geoip2"
)

func main() {
	api := geoip2.New(os.Getenv("MAXMIND_USER_ID"), os.Getenv("MAXMIND_LICENSE_KEY"))
	resp, _ := api.City(nil, "8.8.8.8")
	json.NewEncoder(os.Stdout).Encode(resp)
}
