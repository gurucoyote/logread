package cmd

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net"
	"net/http"

	"github.com/spf13/cobra"
)

type GeoInfo struct {
	Query      string  `json:"query"`
	Status     string  `json:"status"`
	Country    string  `json:"country"`
	RegionName string  `json:"regionName"`
	City       string  `json:"city"`
	Zip        string  `json:"zip"`
	Lat        float64 `json:"lat"`
	Lon        float64 `json:"lon"`
	ISP        string  `json:"isp"`
	Org        string  `json:"org"`
	As         string  `json:"as"`
}

var geoipCmd = &cobra.Command{
	Use:   "geoip [ip]",
	Short: "Get geolocation information for an IP address",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		ip := net.ParseIP(args[0])
		if ip == nil {
			fmt.Println("Invalid IP address:", args[0])
			return
		}
		getGeoInfo(args[0])
	},
}

func getGeoInfo(ip string) {
	resp, err := http.Get("http://ip-api.com/json/" + ip)
	if err != nil {
		log.Fatalln(err)
	}

	body, readErr := ioutil.ReadAll(resp.Body)
	if readErr != nil {
		log.Fatal(readErr)
	}

	var geoInfo GeoInfo
	decodeErr := json.Unmarshal(body, &geoInfo)
	if decodeErr != nil {
		log.Fatalf("Decoder error: %v", decodeErr)
	}

	if geoInfo.Status == "success" {
		fmt.Println("Query IP:", geoInfo.Query)
		fmt.Println("Country:", geoInfo.Country)
		fmt.Println("Region:", geoInfo.RegionName)
		fmt.Println("City:", geoInfo.City)
		fmt.Println("Zip:", geoInfo.Zip)
		fmt.Println("Lat:", geoInfo.Lat)
		fmt.Println("Lon:", geoInfo.Lon)
		fmt.Println("ISP:", geoInfo.ISP)
		fmt.Println("Organization:", geoInfo.Org)
		fmt.Println("As:", geoInfo.As)
	} else {
		fmt.Println("Unable to get information for IP:", ip)
	}
}

func init() {
	RootCmd.AddCommand(geoipCmd)
}
