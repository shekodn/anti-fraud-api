package models

import(
  // "fmt"
  "os"
  s "strings"
  "github.com/gocarina/gocsv"
)

type City struct {
  Name string `csv:"city_ascii"`
  Coordinate
  CountryCode string `csv:"iso2"`
  NotUsed string `csv:"-"`
}

type Coordinate struct {
	Lat float64 `csv:"lat"`
  Lng float64 `csv:"lng"`
}

var citiesDB = make(map[string]Coordinate)

func init() {
  citiesFile, err := os.OpenFile("worldcities.csv", os.O_RDWR|os.O_CREATE, os.ModePerm)
  if err != nil {
    panic(err)
  }
  defer citiesFile.Close()

  cities := []*City{}

  if err := gocsv.UnmarshalFile(citiesFile, &cities); err != nil { // Load cities from file
    panic(err)
  }

  for _, city := range cities {
    citiesDB[s.ToLower(city.Name+city.CountryCode)] = Coordinate{Lat: city.Lat, Lng: city.Lng }
    // city.Lat: city.Lat, city.Lng: city.Lng}
  }
}

func GetCitiesDB() map[string]Coordinate {
  return citiesDB
}
