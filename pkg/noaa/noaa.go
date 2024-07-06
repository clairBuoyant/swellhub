package noaa

import (
	"encoding/csv"
	"io"
	"log/slog"
	"net/http"
	"regexp"
	"strconv"
	"strings"
	"time"
)

func extractHeaders(r *csv.Reader) ([]string, []string, error) {
	var headers []string
	var units []string

	for i := 0; i < 2; i++ {
		record, err := r.Read()
		if err == io.EOF {
			return nil, nil, NewError("not enough data", err)
		}
		if err != nil {
			return nil, nil, NewError("error reading CSV", err)
		}
		if i == 0 {
			headers = record
		} else {
			units = record
		}
	}

	return headers, units, nil
}

func request(url string) ([]byte, int, error) {
	response, err := http.Get(url)
	if err != nil {
		return nil, 0, err
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			slog.Error(err.Error())
		}
	}(response.Body)

	body, err := io.ReadAll(response.Body)
	if err != nil {
		return nil, 0, err
	}

	return body, response.StatusCode, nil
}

func parseValue(value string) (float32, error) {
	if value == "MM" {
		return 0, nil
	}

	parsedValue, err := strconv.ParseFloat(value, 32)
	if err != nil {
		return 0, err
	}

	return float32(parsedValue), nil
}

func parseRecordToStruct(record []string, mo *MeteorologicalObservation) error {
	// TODO: refactor parsing approach

	row := record[0]
	trimmed := strings.TrimSpace(row)
	singleSpacePattern := regexp.MustCompile(`\s+`)
	rowValues := strings.Split(singleSpacePattern.ReplaceAllString(trimmed, " "), " ")

	year, _ := strconv.ParseInt(rowValues[0], 10, 16)
	month, _ := strconv.ParseInt(rowValues[1], 10, 16)
	day, _ := strconv.ParseInt(rowValues[2], 10, 16)
	hour, _ := strconv.ParseInt(rowValues[3], 10, 16)
	minute, _ := strconv.ParseInt(rowValues[4], 10, 16)
	mo.Datetime = time.Date(int(year), time.Month(month), int(day), int(hour), int(minute), 0, 0, time.UTC)

	if value, err := parseValue(rowValues[5]); err == nil {
		mo.WindDirection = int16(value)
	} else {
		return err
	}

	if value, err := parseValue(rowValues[6]); err == nil {
		mo.WindSpeed = value
	} else {
		return err
	}

	if value, err := parseValue(rowValues[7]); err == nil {
		mo.WindGust = value
	} else {
		return err
	}

	if value, err := parseValue(rowValues[8]); err == nil {
		mo.WaveHeight = value
	} else {
		return err
	}

	if value, err := parseValue(rowValues[9]); err == nil {
		mo.DominantWavePeriod = value
	} else {
		return err
	}

	if value, err := parseValue(rowValues[10]); err == nil {
		mo.AverageWavePeriod = value
	} else {
		return err
	}

	if value, err := parseValue(rowValues[11]); err == nil {
		mo.WaveDirection = int16(value)
	} else {
		return err
	}

	if value, err := parseValue(rowValues[12]); err == nil {
		mo.SeaLevelPressure = value
	} else {
		return err
	}

	if value, err := parseValue(rowValues[13]); err == nil {
		mo.AirTemperature = value
	} else {
		return err
	}

	if value, err := parseValue(rowValues[14]); err == nil {
		mo.WaterTemperature = value
	} else {
		return err
	}

	if value, err := parseValue(rowValues[15]); err == nil {
		mo.DewpointTemperature = value
	} else {
		return err
	}

	if value, err := parseValue(rowValues[16]); err == nil {
		mo.Visibility = value
	} else {
		return err
	}

	if value, err := parseValue(rowValues[17]); err == nil {
		mo.PressureTendency = value
	} else {
		return err
	}

	if value, err := parseValue(rowValues[18]); err == nil {
		mo.Tide = value
	} else {
		return err
	}

	return nil
}

func realtimeMeteorological(url string) ([]MeteorologicalObservation, error) {
	body, statusCode, err := request(url)
	if err != nil {
		return nil, NewRequestError(statusCode, err.Error(), err)
	}

	r := csv.NewReader(strings.NewReader(string(body)))
	r.FieldsPerRecord = 0
	r.TrimLeadingSpace = true

	extractHeaders(r)

	var mos []MeteorologicalObservation
	for {
		var mo MeteorologicalObservation
		record, err := r.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			return nil, NewError("error reading CSV", err)
		}

		if err := parseRecordToStruct(record, &mo); err != nil {
			return nil, NewError("error parsing meteorological record to struct", err)
		}

		mos = append(mos, mo)
	}
	return mos, err
}
