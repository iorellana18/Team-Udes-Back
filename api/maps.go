package api

import (
	"golang.org/x/net/context"
	"googlemaps.github.io/maps"
)

var apiKey string = "AIzaSyD-cX_4jOdNtz6P0eBiRN30_xIf9YhwjF8"

func MissionReverseGeocodingPoint(latitude float64, longitude float64) (string, error) {
	var address string

	client, err := maps.NewClient(maps.WithAPIKey(apiKey))

	if err != nil {
		return address, err
	} else {
		point := &maps.LatLng{
			Lat: latitude,
			Lng: longitude,
		}

		request := &maps.GeocodingRequest{
			LatLng: point,
		}

		res, err := client.ReverseGeocode(context.Background(), request)
		if err == nil {
			address = res[0].FormattedAddress
		} else if err.Error() == "maps: ZERO_RESULTS - " {
			address = "Zona Desconocida"
			err = nil
		}

		return address, err

	}
}

func EmergencyReverseGeocodingPoint(latitude float64, longitude float64) (string, string, string, error) {
	var commune, city, region string

	client, err := maps.NewClient(maps.WithAPIKey(apiKey))

	if err != nil {
		return commune, city, region, err
	} else {
		point := &maps.LatLng{
			Lat: latitude,
			Lng: longitude,
		}

		request := &maps.GeocodingRequest{
			LatLng: point,
		}

		res, err := client.ReverseGeocode(context.Background(), request)
		if err == nil {
			for i := range res[0].AddressComponents {
				for j := range res[0].AddressComponents[i].Types {
					if res[0].AddressComponents[i].Types[j] == "administrative_area_level_1" {
						region = res[0].AddressComponents[i].LongName
					} else if res[0].AddressComponents[i].Types[j] == "administrative_area_level_2" {
						city = res[0].AddressComponents[i].LongName
					} else if res[0].AddressComponents[i].Types[j] == "administrative_area_level_3" {
						commune = res[0].AddressComponents[i].LongName
					}
				}
			}
		} else if err.Error() == "maps: ZERO_RESULTS - " {
			region = "Zona Desconocida"
			err = nil
		}

		return commune, city, region, err

	}
}
