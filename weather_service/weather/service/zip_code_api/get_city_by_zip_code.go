package zipcodeapi

import (
	"encoding/json"
	"fmt"
	"log"
	"os"

	"github.com/hgtpcastro/go-expert-lab-google-cloud-run/weather_service/pkg/http/client"
	"github.com/hgtpcastro/go-expert-lab-google-cloud-run/weather_service/weather/zip_code/features/getting_city_by_zip_code/v1/dtos"
)

func GetCityByZipCode(zipCode string) (dtos.GetCityByZipCodeResponseDTO, error) {
	client := client.NewHttpClient()

	api := os.Getenv("ZIP_CODE_API_URL")
	url := fmt.Sprintf(api, zipCode)

	log.Println(url)

	resp, err := client.R().Get(url)
	if err != nil {
		return dtos.GetCityByZipCodeResponseDTO{}, err
	}

	var data dtos.GetCityByZipCodeResponseDTO
	err = json.Unmarshal(resp.Body(), &data)

	log.Printf("%q", data)

	if err != nil {
		return dtos.GetCityByZipCodeResponseDTO{}, err
	}

	return data, nil
}
