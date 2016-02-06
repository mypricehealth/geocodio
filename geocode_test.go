package geocodio_test

import (
	"github.com/stevepartridge/geocodio"
	"testing"
)

func TestGeocodeFullAddress(t *testing.T) {
	Geocodio, err := geocodio.NewGeocodio(API_KEY)
	if err != nil {
		t.Error("Failed with API KEY set.", err)
	}
	result, err := Geocodio.Geocode(TEST_ADDRESS_1_FULL)
	if err != nil {
		t.Error(err)
	}

	t.Log(result.Debug)

	if len(result.Results) == 0 {
		t.Error("Results length is 0")
	}

	if result.Results[0].Location.Latitude != 33.739464 {
		t.Error("Location latitude does not match", result.Results[0].Location.Latitude, 33.739464)
	}

	if result.Results[0].Location.Longitude != -116.40803 {
		t.Error("Location longitude does not match", result.Results[0].Location.Longitude, -116.40803)
	}
}

func TestGeocodeFullAddressReturningTimezone(t *testing.T) {
	Geocodio, err := geocodio.NewGeocodio(API_KEY)
	if err != nil {
		t.Error("Failed with API KEY set.", err)
	}
	result, err := Geocodio.GeocodeAndReturnTimezone(TEST_ADDRESS_1_FULL)
	if err != nil {
		t.Error(err)
	}

	t.Log(result.Debug)

	if len(result.Results) == 0 {
		t.Error("Results length is 0")
	}

	if result.Results[0].Location.Latitude != 33.739464 {
		t.Error("Location latitude does not match", result.Results[0].Location.Latitude, 33.739464)
	}

	if result.Results[0].Location.Longitude != -116.40803 {
		t.Error("Location longitude does not match", result.Results[0].Location.Longitude, -116.40803)
	}

	if result.Results[0].Fields.Timezone.Name == "" {
		t.Error("Timezone field not found")
	}

	if !result.Results[0].Fields.Timezone.ObservesDST {
		t.Error("Timezone field does not match", result.Results[0].Fields.Timezone)
	}
}

func TestGeocodeFullAddressReturningCongressionalDistrict(t *testing.T) {
	Geocodio, err := geocodio.NewGeocodio(API_KEY)
	if err != nil {
		t.Error("Failed with API KEY set.", err)
		t.Fail()
	}
	result, err := Geocodio.GeocodeAndReturnCongressionalDistrict(TEST_ADDRESS_1_FULL)
	if err != nil {
		t.Error(err)
	}

	t.Log(result.Debug)

	if len(result.Results) == 0 {
		t.Error("Results length is 0")
		t.Fail()
	}

	if result.Results[0].Location.Latitude != 33.739464 {
		t.Error("Location latitude does not match", result.Results[0].Location.Latitude, 33.739464)
		t.Fail()
	}

	if result.Results[0].Location.Longitude != -116.40803 {
		t.Error("Location longitude does not match", result.Results[0].Location.Longitude, -116.40803)
		t.Fail()
	}

	if result.Results[0].Fields.CongressionalDistrict.Name == "" {
		t.Error("Congressional District field not found", result.Results[0].Fields.CongressionalDistrict)
		t.Fail()
	}

	if result.Results[0].Fields.CongressionalDistrict.DistrictNumber != 36 {
		t.Error("Congressional District field does not match", result.Results[0].Fields.CongressionalDistrict)
		t.Fail()
	}
}

func TestGeocodeFullAddressReturningStateLegislativeDistricts(t *testing.T) {
	Geocodio, err := geocodio.NewGeocodio(API_KEY)
	if err != nil {
		t.Error("Failed with API KEY set.", err)
		t.Fail()
	}

	result, err := Geocodio.GeocodeAndReturnStateLegislativeDistricts(TEST_ADDRESS_1_FULL)
	if err != nil {
		t.Error(err)
	}

	t.Log(result.Debug)

	if len(result.Results) == 0 {
		t.Error("Results length is 0", result)
		t.Fail()
	}

	if result.Results[0].Location.Latitude != 33.739464 {
		t.Error("Location latitude does not match", result.Results[0].Location.Latitude, 33.739464)
		t.Fail()
	}

	if result.Results[0].Location.Longitude != -116.40803 {
		t.Error("Location longitude does not match", result.Results[0].Location.Longitude, -116.40803)
		t.Fail()
	}

	if result.Results[0].Fields.StateLegislativeDistricts.House.DistrictNumber != "42" {
		t.Error("State Legislative Districts house does not match", result.Results[0].Fields.StateLegislativeDistricts.House)
		t.Fail()
	}

	if result.Results[0].Fields.StateLegislativeDistricts.Senate.DistrictNumber != "28" {
		t.Error("State Legislative Districts senate does not match", result.Results[0].Fields.StateLegislativeDistricts.Senate)
		t.Fail()
	}
}

func TestGeocodeFullAddressReturningMultipleFields(t *testing.T) {
	Geocodio, err := geocodio.NewGeocodio(API_KEY)
	if err != nil {
		t.Error("Failed with API KEY set.", err)
	}
	result, err := Geocodio.GeocodeReturnFields(TEST_ADDRESS_1_FULL, "timezone", "cd")
	if err != nil {
		t.Error(err)
	}

	t.Log(result.Debug)

	if len(result.Results) == 0 {
		t.Error("Results length is 0")
	}

	if result.Results[0].Location.Latitude != 33.739464 {
		t.Error("Location latitude does not match", result.Results[0].Location.Latitude, 33.739464)
	}

	if result.Results[0].Location.Longitude != -116.40803 {
		t.Error("Location longitude does not match", result.Results[0].Location.Longitude, -116.40803)
	}

	if result.Results[0].Fields.Timezone.Name == "" {
		t.Error("Timezone field not found")
	}

	if !result.Results[0].Fields.Timezone.ObservesDST {
		t.Error("Timezone field does not match", result.Results[0].Fields.Timezone)
	}

	// check congressional district
	if result.Results[0].Fields.CongressionalDistrict.Name == "" {
		t.Error("Congressional District field not found", result.Results[0].Fields.CongressionalDistrict)
	}

	if result.Results[0].Fields.CongressionalDistrict.DistrictNumber != 36 {
		t.Error("Congressional District field does not match", result.Results[0].Fields.CongressionalDistrict)
	}

}

// TODO: School District (school)
