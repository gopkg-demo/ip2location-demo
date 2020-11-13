package main

import (
	"fmt"
	"net"
	"strings"

	"github.com/pkg/errors"

	"github.com/ip2location/ip2location-go"
)

func main() {
	country, err := GetCountryByIP("79.53.228.61")
	if err != nil {
		panic(err)
	}
	fmt.Println(country)
}

func GetCountryByIP(s string) (string, error) {

	ip := net.ParseIP(s)

	if ip == nil || s == "::1" {
		return "N/A", errors.New("invalid ip")
	}

	var dbFilePath string

	if strings.Count(s, ":") >= 2 {
		dbFilePath = "./IP2LOCATION-LITE-DB1.IPV6.BIN"
	} else {
		dbFilePath = "./IP2LOCATION-LITE-DB1.BIN"
	}

	db, err := ip2location.OpenDB(dbFilePath)

	if err != nil {
		return "N/A", errors.WithMessage(err, "invalid open IP2LOCATION db file.")
	}

	defer db.Close()

	results, err := db.Get_country_short(s)

	if err != nil {
		return "N/A", errors.WithMessage(err, "call to Get_country_short failed.")
	}

	return results.Country_short, nil
}
