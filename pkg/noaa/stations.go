package noaa

import (
	"encoding/xml"
	"fmt"
)

type Stations struct {
	XMLName  xml.Name  `xml:"stations" json:"-"`
	Stations []Station `xml:"station"`
}

type Station struct {
	XMLName      xml.Name `xml:"station" json:"-"`
	ID           string   `xml:"id,attr"`
	Lat          float32  `xml:"lat,attr"`
	Lon          float32  `xml:"lon,attr"`
	Name         string   `xml:"name,attr"`
	Owner        string   `xml:"owner,attr"`
	Type         string   `xml:"type,attr"`
	Met          string   `xml:"met,attr"`
	Currents     string   `xml:"currents,attr"`
	WaterQuality string   `xml:"waterquality,attr"`
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
