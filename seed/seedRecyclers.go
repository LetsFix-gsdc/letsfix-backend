package seed

import (
	"encoding/json"
	"fmt"
	"gsdc/letsfix/models"

	"gsdc/letsfix/repository"
	"io/ioutil"
	"os"
	"strconv"
	"strings"

	"github.com/antchfx/xmlquery"
)

type Recycler_Data struct {
	Type string `json:"type"`
	Crs  struct {
		Type       string `json:"type"`
		Properties struct {
			Name string `json:"name"`
		} `json:"properties"`
	} `json:"crs"`
	Features []struct {
		Type       string `json:"type"`
		Properties struct {
			Name        string `json:"Name"`
			Description string `json:"Description"`
		} `json:"properties"`
		Geometry struct {
			Type        string    `json:"type"`
			Coordinates []float64 `json:"coordinates"`
		} `json:"geometry"`
	} `json:"features"`
}

func ReadFile(path string) *os.File {

	file, err := os.Open(path)
	if err != nil {
		panic("error opening file")
	}

	byteValue, err := ioutil.ReadAll(file)
	var recyclerData Recycler_Data
	json.Unmarshal(byteValue, &recyclerData)

	for i := 0; i < len(recyclerData.Features); i++ {
		var feature = recyclerData.Features[i]
		fmt.Printf(
			"name: %s lat: %f long: %f\n",
			feature.Properties.Name, feature.Geometry.Coordinates[0], feature.Geometry.Coordinates[1],
		)
	}

	defer file.Close()
	return file
}

func SeedFromXml(path string) {

	file, err := os.Open(path)
	if err != nil {
		panic("error opening file")
	}

	byteValue, err := ioutil.ReadAll(file)
	s := string(byteValue)

	doc, err := xmlquery.Parse(strings.NewReader(s))
	if err != nil {
		panic(err)
	}

	recyclerRepo := repository.NewRecyclerRepository()

	for _, n := range xmlquery.Find(doc, "//Placemark") {
		name := n.SelectElement("//SimpleData[4]").InnerText()
		email := n.SelectElement("//SimpleData[3]").InnerText()
		address := n.SelectElement("//SimpleData[5]").InnerText()
		coords := n.SelectElement("//coordinates").InnerText()
		coordList := strings.Split(coords, ",")
		lat, _ := strconv.ParseFloat(coordList[0], 32)
		long, _ := strconv.ParseFloat(coordList[0], 32)
		// fmt.Println(name, email, address, coordList[0], coordList[1])

		recycler := models.Recycler{
			Name:       name,
			Email:      email,
			Address:    address,
			Latitude:   lat,
			Longtitude: long,
		}

		fmt.Println(recycler)
		recyclerRepo.SaveRecycler(recycler)
	}

}

func SeedRecyclers() {

}
