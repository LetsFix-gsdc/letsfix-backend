package seed

import (
	"fmt"
	"gsdc/letsfix/models"

	"gsdc/letsfix/repository"
	"io/ioutil"
	"os"
	"strconv"
	"strings"

	"github.com/antchfx/xmlquery"
)

func SeedFromXml(path string) {

	file, err := os.Open(path)
	if err != nil {
		panic("error opening file")
	}

	byteValue, _ := ioutil.ReadAll(file)
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
		lat, _ := strconv.ParseFloat(coordList[0], 64)
		long, _ := strconv.ParseFloat(coordList[1], 64)

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
