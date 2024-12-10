package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"log"
	"os"
	"path"
	"strings"

	"github.com/oschwald/geoip2-golang"
	"github.com/oschwald/maxminddb-golang"
)

func containsIgnoreCase(s string, substr string) bool {
	return strings.Contains(strings.ToLower(s), strings.ToLower(substr))
}

func dumpCity(networks *maxminddb.Networks, writer *csv.Writer, lang string) (err error) {
	headers := []string{
		"Network",
		"Country",
		"Region",
		"City",
		"Latitude",
		"Longitude",
	}

	err = writer.Write(headers)
	if err != nil {
		return err
	}

	// City is nested
	record := geoip2.City{}
	empty_string := ""
	for networks.Next() {
		subnet, err := networks.Network(&record)
		if err != nil {
			log.Fatalln(err)
		}
		values := []string{
			subnet.String(),
			record.Country.Names[lang],
			empty_string,
			record.City.Names[lang],
			fmt.Sprintf("%v", record.Location.Latitude),
			fmt.Sprintf("%v", record.Location.Longitude),
		}
		
		err = writer.Write(values)
		
		if err != nil {
			return err
		}
	}
	return nil
}

func dumpCountry(networks *maxminddb.Networks, writer *csv.Writer, lang string) (err error) {
	headers := []string{
		"Network",
		"Country",
		"Region",
		"City",
		"Latitude",
		"Longitude",
	}
	err = writer.Write(headers)
	if err != nil {
		return err
	}
	empty_string := ""
	record := geoip2.Country{}
	for networks.Next() {
		subnet, err := networks.Network(&record)
		if err != nil {
			log.Fatalln(err)
		}
		values := []string{
			subnet.String(),
			record.Country.Names[lang],
			empty_string,
			empty_string,
			empty_string,
			empty_string,
		}
		
		err = writer.Write(values)
		
		if err != nil {
			return err
		}
	}
	return nil
}

func main() {

	//new
	var lang string
	flag.StringVar(&lang, "lang", "ru", "language")
	
	flag.Parse()
	if flag.NArg() == 0 {
		log.Fatal("Please provide mmdb path")
	}

	fullPath := flag.Args()[0]
	fname := path.Base(fullPath)

	// open mmdb
	db, err := maxminddb.Open(fullPath)
	if err != nil {
		log.Fatal(err)
	}
	defer func(db *maxminddb.Reader) {
		err := db.Close()
		if err != nil {
			// ignore
		}
	}(db)

	// open CSV writer, and write header
	writer := csv.NewWriter(os.Stdout)
	defer writer.Flush()

	// skip aliased networks
	networks := db.Networks(maxminddb.SkipAliasedNetworks)
	if networks.Err() != nil {
		log.Fatalln(networks.Err())
	}
	var err2 error
	if containsIgnoreCase(fname, "city") {
		err2 = dumpCity(networks, writer, lang)
	} else if containsIgnoreCase(fname, "country") {
		err2 = dumpCountry(networks, writer, lang)
	} else {
		log.Fatal("Dump type not recognized, please rename mmdb file to contain any of the following strings: city, country")
	}
	if err2 != nil {
		log.Fatal(err2.Error())
	}
	if networks.Err() != nil {
		log.Panic(networks.Err())
	}
}
