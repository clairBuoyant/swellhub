package noaa

import (
	"encoding/xml"
	"fmt"
)

type Stations struct {
	XMLName  xml.Name  `xml:"stations" json:"-"`
	Stations []Station `xml:"station" json:"stations"`
}

type Station struct {
	XMLName      xml.Name `xml:"station" json:"-"`
	ID           string   `xml:"id,attr" json:"id"`
	Lat          float32  `xml:"lat,attr" json:"latitude"`
	Lon          float32  `xml:"lon,attr" json:"longitude"`
	Name         string   `xml:"name,attr" json:"name"`
	Owner        string   `xml:"owner,attr" json:"owner"`
	Type         string   `xml:"type,attr" json:"type"`
	Met          string   `xml:"met,attr" json:"met"`
	Currents     string   `xml:"currents,attr" json:"currents"`
	WaterQuality string   `xml:"waterquality,attr" json:"waterQuality"`
}

// ActiveStations provides information about all currently active
// stations marked as established by the NDBC Data Assembly Center.
//
// TAO stations are excluded. Each station has an indicator showing
// whether the elevation, meteorological data, single point or profile
// currents data, water quality data, or DART data are available.
func ActiveStations() (*Stations, error) {
	url := fmt.Sprintf("%s.%s", ActiveStationsURL, "xml")

	response, code, err := request(url)
	if err != nil {
		return nil, NewRequestError(code, "NDBC request error", err)
	}

	var activeStations Stations
	if err := xml.Unmarshal(response, &activeStations); err != nil {
		return nil, NewError(fmt.Sprintf("unmarshall error: %s", err.Error()), err)
	}

	return &activeStations, nil
}
