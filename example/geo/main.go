package main

import (
	"fmt"
	"log"
	"net"
	"os"

	"github.com/ipinfo/go/ipinfo"
)

func main() {
	if len(os.Args) == 1 {
		info, err := ipinfo.GetGeo(nil)
		if err != nil {
			log.Println(err)
		}
		printInfo(info)
		return
	}
	for _, s := range os.Args[1:] {
		ip := net.ParseIP(s)
		if ip == nil {
			continue
		}
		info, err := ipinfo.GetGeo(ip)
		if err != nil {
			log.Println(err)
		}
		printInfo(info)
	}
}

func printInfo(info *ipinfo.Geo) {
	fmt.Printf("IP: %v\nCity: %s\nRegion: %s\nCountry: %s\nLocation: %s\nPhone: %s\nPostal: %s\n",
		info.IP, info.City, info.Region, info.Country, info.Location, info.Phone, info.Postal)
}
